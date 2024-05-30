package stat

import (
	"github.com/go-kratos/kratos/v2/log"
	"testing"
	"time"
)

func Test_log(t *testing.T) {

	log.Error("Current bucket atomic Value is nil in BucketLeapArray.currentBucketWithTime()")
}

func Test_Add(t *testing.T) {

	u := uint32(time.Second.Seconds() * 10)
	leapArray := NewBucketLeapArray(10, u)
	leapArray.AddCount(0, 1)
	leapArray.AddCount(0, 1)
	leapArray.AddCount(0, 1)

	time.Sleep(time.Second)
	leapArray.AddCount(0, 1)

	time.Sleep(time.Second * 20)
	leapArray.AddCount(0, 4)

	t.Log(leapArray.Count(0))
}

func Test_BulkAdd(t *testing.T) {
	u := uint32(time.Second.Seconds() * 10)
	leapArray := NewBucketLeapArray(10, u)
	leapArray.BulkAddCount(map[MetricEvent]int64{1: 1, 0: 20})
	leapArray.BulkAddCount(map[MetricEvent]int64{0: 1, 1: 2, 2: 12})
	result := leapArray.BulkCount(8, 3, 5)
	t.Log(result)
}

func Test_duartion(t *testing.T) {
	duartion := 0.012188077

	res := (time.Millisecond * 500).Seconds() > duartion
	t.Log((time.Millisecond * 500).Seconds(), res)

}