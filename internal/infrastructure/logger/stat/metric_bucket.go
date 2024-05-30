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
	"sync/atomic"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

// MetricBucket represents the entity to record metrics per minimum time unit (i.e. the bucket time span).
// Note that all operations of the MetricBucket are required to be thread-safe.
type MetricBucket struct {
	// Value of statistic
	counter [MetricEventTotal]int64
	rt      [MetricEventTotal]int64
}

func NewMetricBucket() *MetricBucket {
	mb := &MetricBucket{}

	return mb
}

// Add statistic count for the given metric event.
func (mb *MetricBucket) Add(event MetricEvent, count int64) {
	if event < 0 {
		log.Error(errors.Errorf("Unknown metric event: %v", event), "")
		return
	}

	mb.addCount(event, count)
}

func (mb *MetricBucket) AddRt(event MetricEvent, rt int64) {
	atomic.AddInt64(&mb.counter[event], rt)
}

func (mb *MetricBucket) addCount(event MetricEvent, count int64) {
	atomic.AddInt64(&mb.counter[event], count)
}

// Get current statistic count of the given metric event.
func (mb *MetricBucket) Get(event MetricEvent) int64 {
	if event < 0 {
		log.Error(errors.Errorf("Unknown metric event: %v", event), "")
		return 0
	}
	return atomic.LoadInt64(&mb.counter[event])
}
