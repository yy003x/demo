// Copyright 1999-2020 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stat

import (
	"reflect"
	"sync/atomic"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

// BucketLeapArray is the sliding window implementation based on LeapArray (as the sliding window infrastructure)
// and MetricBucket (as the data type). The MetricBucket is used to record statistic
// metrics per minimum time unit (i.e. the bucket time span).
type BucketLeapArray struct {
	data     LeapArray
	dataType string
}

func (bla *BucketLeapArray) NewEmptyBucket() interface{} {
	return NewMetricBucket()
}

func (bla *BucketLeapArray) ResetBucketTo(bw *BucketWrap, startTime uint64) *BucketWrap {
	atomic.StoreUint64(&bw.BucketStart, startTime)
	bw.Value.Store(NewMetricBucket())
	return bw
}

// NewBucketLeapArray creates a BucketLeapArray with given attributes.
//
// The sampleCount represents the number of buckets, while intervalInMs represents
// the total time span of sliding window. Note that the sampleCount and intervalInMs must be positive
// and satisfies the condition that intervalInMs%sampleCount == 0.
// The validation must be done before call NewBucketLeapArray.
func NewBucketLeapArray(sampleCount uint32, intervalInMs uint32) *BucketLeapArray {
	// TODO: also check params here.
	bucketLengthInMs := intervalInMs / sampleCount // 定时时间 / 窗口大小 = 每个窗口持续时间
	ret := &BucketLeapArray{
		data: LeapArray{
			bucketLengthInMs: bucketLengthInMs,
			sampleCount:      sampleCount,
			intervalInMs:     intervalInMs,
			array:            nil,
		},
		dataType: "MetricBucket",
	}
	arr := NewAtomicBucketWrapArray(int(sampleCount), bucketLengthInMs, ret)
	ret.data.array = arr
	return ret
}

func (bla *BucketLeapArray) SampleCount() uint32 {
	return bla.data.sampleCount
}

func (bla *BucketLeapArray) IntervalInMs() uint32 {
	return bla.data.intervalInMs
}

func (bla *BucketLeapArray) BucketLengthInMs() uint32 {
	return bla.data.bucketLengthInMs
}

func (bla *BucketLeapArray) DataType() string {
	return bla.dataType
}

func (bla *BucketLeapArray) GetIntervalInSecond() float64 {
	return float64(bla.IntervalInMs()) / 1000.0
}

func (bla *BucketLeapArray) AddCount(event MetricEvent, count int64) {
	// It might panic?
	//添加值
	bla.addCountWithTime(CurrentTimeMillis(), event, count)
}

func (bla *BucketLeapArray) BulkAddCount(eventCounts map[MetricEvent]int64) {
	now := CurrentTimeMillis()
	for event, count := range eventCounts {
		bla.addCountWithTime(now, event, count)
	}
}

//获取当前毫秒时间戳   event具体统计指标  数量
func (bla *BucketLeapArray) addCountWithTime(now uint64, event MetricEvent, count int64) {
	b := bla.currentBucketWithTime(now)
	if b == nil {
		return
	}
	b.Add(event, count)
}

func (bla *BucketLeapArray) currentBucketWithTime(now uint64) *MetricBucket {
	//获取当前的bucket
	curBucket, err := bla.data.currentBucketOfTime(now, bla)
	if err != nil {
		log.Error(err, "Failed to get current bucket in BucketLeapArray.currentBucketWithTime()", "now", now)
		return nil
	}
	if curBucket == nil {
		log.Error(errors.New("current bucket is nil"), "Nil curBucket in BucketLeapArray.currentBucketWithTime()")
		return nil
	}
	mb := curBucket.Value.Load()
	if mb == nil {
		log.Error(errors.New("nil bucket"), "Current bucket atomic Value is nil in BucketLeapArray.currentBucketWithTime()")
		return nil
	}
	b, ok := mb.(*MetricBucket)
	if !ok {
		log.Error(errors.New("fail to type assert"), "Bucket data type error in BucketLeapArray.currentBucketWithTime()", "expectType", "*MetricBucket", "actualType", reflect.TypeOf(mb).Name())
		return nil
	}
	return b
}

// Count returns the sum count for the given MetricEvent within all valid (non-expired) buckets.
func (bla *BucketLeapArray) Count(event MetricEvent) int64 {
	// it might panic?
	return bla.CountWithTime(CurrentTimeMillis(), event)
}

func (bla *BucketLeapArray) BulkCount(event ...MetricEvent) map[MetricEvent]int64 {
	now := CurrentTimeMillis()
	_, err := bla.data.currentBucketOfTime(now, bla)
	if err != nil {
		log.Error(err, "Failed to get current bucket in BucketLeapArray.CountWithTime()", "now", now)
	}

	eventResult := make(map[MetricEvent]int64, len(event))
	for _, ww := range bla.data.valuesWithTime(now) {
		mb := ww.Value.Load()
		if mb == nil {
			log.Error(errors.New("current bucket is nil"), "Failed to load current bucket in BucketLeapArray.CountWithTime()")
			continue
		}
		b, ok := mb.(*MetricBucket)
		if !ok {
			log.Error(errors.New("fail to type assert"), "Bucket data type error in BucketLeapArray.CountWithTime()", "expectType", "*MetricBucket", "actualType", reflect.TypeOf(mb).Name())
			continue
		}

		for _, e := range event {
			eventResult[e] += b.Get(e)
		}
	}
	return eventResult
}

func (bla *BucketLeapArray) CountWithTime(now uint64, event MetricEvent) int64 {
	_, err := bla.data.currentBucketOfTime(now, bla)
	if err != nil {
		log.Error(err, "Failed to get current bucket in BucketLeapArray.CountWithTime()", "now", now)
	}
	count := int64(0)
	for _, ww := range bla.data.valuesWithTime(now) {
		mb := ww.Value.Load()
		if mb == nil {
			log.Error(errors.New("current bucket is nil"), "Failed to load current bucket in BucketLeapArray.CountWithTime()")
			continue
		}
		b, ok := mb.(*MetricBucket)
		if !ok {
			log.Error(errors.New("fail to type assert"), "Bucket data type error in BucketLeapArray.CountWithTime()", "expectType", "*MetricBucket", "actualType", reflect.TypeOf(mb).Name())
			continue
		}
		count += b.Get(event)
	}
	return count
}

// Values returns all valid (non-expired) buckets.
func (bla *BucketLeapArray) Values(now uint64) []*BucketWrap {
	// Refresh current bucket if necessary.
	_, err := bla.data.currentBucketOfTime(now, bla)
	if err != nil {
		log.Error(err, "Failed to refresh current bucket in BucketLeapArray.Values()", "now", now)
	}

	return bla.data.valuesWithTime(now)
}

func (bla *BucketLeapArray) ValuesConditional(now uint64, predicate TimePredicate) []*BucketWrap {
	return bla.data.ValuesConditional(now, predicate)
}
