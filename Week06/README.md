# Week06 作业题目：

参考 Hystrix 实现一个滑动窗口计数器。

```go
import (
	"sync"
	"time"
)

const (
	DefaultRollSize = 10 // 默认记录长度，秒级
)

type roll struct {
	buckets map[int64]*bucket
	size    int64 // buckets记录时间长度
	mutex   sync.RWMutex
}

type bucket struct {
	Value int
}

func NewRoll(size ...int) *roll {
	r := &roll{
		buckets: make(map[int64]*bucket),
		mutex:   sync.RWMutex{},
		size:    DefaultRollSize,
	}
	if len(size) > 0 {
		r.size = int64(size[0])
	}
	return r
}

func (r *roll) getCurrentBucket() *bucket {
	now := time.Now().Unix()
	b := r.buckets[now]
	if b == nil {
		b = &bucket{}
		r.buckets[now] = b
	}
	return b
}

func (r *roll) removeExpiredBucket() {
	before := time.Now().Unix() - r.size

	for timestamp := range r.buckets {
		if timestamp <= before {
			delete(r.buckets, timestamp)
		}
	}
}

func (r *roll) Increment(i int) {
	if i == 0 {
		return
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	b := r.getCurrentBucket()
	b.Value += i
	r.removeExpiredBucket()
}

// Sum 计算 now 时间戳向前 window 秒的数量
//     windows 默认为 roll.size
func (r *roll) Sum(now time.Time, window ...int) (sum int) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	before := now.Unix()
	if len(window) > 0 {
		before -= int64(window[0])
	} else {
		before -= r.size
	}
	for timestamp, bucket := range r.buckets {
		if timestamp >= before {
			sum += bucket.Value
		}
	}
	return sum
}

// Sum 计算 now 时间戳向前 window 秒的均值
//     windows 默认为 roll.size
func (r *roll) Avg(now time.Time, window ...int) int {
	if len(window) == 0 {
		return int(int64(r.Sum(now)) / r.size)
	} else {
		return r.Sum(now, window[0]) / window[0]
	}
}
```