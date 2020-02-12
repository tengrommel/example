package main

import (
	"fmt"
	"sync"
	"time"
)

//the secret structure holds
//a shared variable
//a sync.RWMutex mutex
//a sync.Mutex mutex
type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

// The Change() function modifies a shared variable, which means that you need to
// use an exclusive lock, which is the reason for using the Lock() and UnLock() functions
func Change(c *secret, pass string) {
	c.RWM.Lock()
	fmt.Println("LChange")
	time.Sleep(10 * time.Second)
	c.password = pass
	c.RWM.Unlock()
}

// The show() function uses the RLock() and RUnlock()
// function because its critical section is used for reading a shared variable.
func show(c *secret) string {
	c.RWM.RLock()
	fmt.Print("show")
	time.Sleep(3 * time.Second)
	defer c.RWM.RUnlock()
	return c.password
}

// The only difference between the code of the showWithLock() function and the code of
// the show() function is that the showWithLock() function uses an exclusive lock for reading,
// which means that only one showWithLock
func showWithLock(c *secret) string {
	c.RWM.Lock()
	fmt.Println("showWithLock")
	time.Sleep(3 * time.Second)
	defer c.RWM.Unlock()
	return c.password
}

func main() {

}
