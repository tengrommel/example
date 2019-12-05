package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

/**
在worker的循环中，为了保证total.value += i 的原子性，我们通过sync.Mutex加锁和解锁来保证该语句在
同一时刻只被一个线程访问。
*/

var totalValue uint64

func workerValue(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		atomic.AddUint64(&totalValue, uint64(i))
	}
}

type singleton struct{}

var (
	instance    *singleton
	initialized uint32
	mu          sync.Mutex
	once        sync.Once
)

func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}
	return instance
}

type Once struct {
	m    sync.Mutex
	done uint32
}

func InstanceOnce() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Println(total.value)
}
