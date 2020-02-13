package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func CatchSig(ch chan os.Signal, done chan bool) {
	sig := <-ch
	fmt.Println("n sig received:", sig)
	// we can set up handles for all types of sig here
	switch sig {
	case syscall.SIGINT:
		fmt.Println("handling a SIGINT now!")
	case syscall.SIGTERM:
		fmt.Println("handling a SIGTERM in an entirely different way!")
	default:
		fmt.Println("unexpected signal received")
	}
	done <- true
}

func main() {
	// initialize our channels
	signals := make(chan os.Signal)
	done := make(chan bool)
	// hook them up to the signals lib
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go CatchSig(signals, done)
	fmt.Println("Press ctrl-c to terminate...")
	<-done
	fmt.Println("Done")
}
