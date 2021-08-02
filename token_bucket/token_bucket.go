package main

import (
	"fmt"
	"sync"
	"time"
)

type bucket struct {
	b int     // 桶的大小
	r float64 // 速率

	mx         sync.Mutex // 加锁
	now        time.Time  // 上次时间
	tokenCount int        // 令牌数量

}

func NewBucket(b int, r float64) *bucket {
	return &bucket{b: b, r: r}
}

func (b *bucket) Allow() bool {
	return b.AllowN(1)
}

func (b *bucket) AllowN(n int) bool {
	b.mx.Lock()
	defer b.mx.Unlock()
	if n > b.b {
		return false
	}
	now := time.Now()
	tn := now.Sub(b.now).Seconds() * b.r
	b.tokenCount += int(tn)
	if b.tokenCount > b.b {
		b.tokenCount = b.b
	}

	if b.tokenCount >= n {
		b.now = now
		b.tokenCount -= n
		return true
	}

	return false
}

func main() {
	//limiter := rate.NewLimiter(1, 5)
	bucket := NewBucket(5, 3)
	a := 0
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var wg sync.WaitGroup
	for a < 4 {
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				if bucket.Allow() {
					fmt.Println("ture --- ", i)
				} else {
					fmt.Println("false --- ", i)
				}
			}(i)
		}
		fmt.Println("===============")
		time.Sleep(time.Second)
		a++
	}

	wg.Wait()
}
