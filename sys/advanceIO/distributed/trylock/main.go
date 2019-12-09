package main

import "sync"

/**
在某些场景，我们只是希望一个任务有单一的执行者，而不像计数器场景那样，所有Goroutine都执行成功。
后来的Goroutine在抢锁失败后，需要放弃其流程。这时候就需要尝试（try lock）了。
*/

// Lock尝试锁
type Lock struct {
	c chan struct{}
}

// NewLock生成一个尝试锁
func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

// Lock try lock, return lock result
func (l Lock) Lock() bool {
	lockResult := true
	select {
	case <-l.c:
		lockResult = true
	default:
	}
	return lockResult
}

// Unlock, Unlock the try lock
func (l Lock) Unlock() {
	l.c <- struct{}{}
}

var counter int

func main() {
	var l = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				// log error
				println("lock failed")
				return
			}
			counter++
			println("current counter", counter)
			l.Unlock()
		}()
	}
	wg.Wait()
}
