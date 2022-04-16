package main

import (
	"fmt"
	"time"
)

func main() {
	myChannel := make(chan int)
	fmt.Println("[>] Running async methods")

	go func() {
		asyncAction1(myChannel)
	}()

	go func() {
		asyncAction2(myChannel)
	}()

	go func() {
		asyncAction3(myChannel)
	}()

	go func() {
		asyncAction4(myChannel)
	}()

	total := 0
	for {
		fmt.Println("[>] Main function will wait for a process to finish")
		finished := <-myChannel
		fmt.Println("[>] Main function received end signal from action ", finished)
		if total++; total >= 4 {
			break
		}
	}

	fmt.Println("[.] The main function has finished.")
}

func asyncAction1(myChannel chan int) {
	lock := GetLock()
	for !lock.IsFree() {
		fmt.Println("  Action 1 is waiting...")
		time.Sleep(1 * time.Second)
	}
	lock.TakeLock()
	fmt.Println("  Action 1 took the lock!")
	time.Sleep(3 * time.Second)
	fmt.Println("  Action 1 executed! ")
	lock.ReleaseLock()
	fmt.Println("  Action 1 released the lock!")

	myChannel <- 1
}

func asyncAction2(myChannel chan int) {
	lock := GetLock()
	for !lock.IsFree() {
		fmt.Println("  Action 2 is waiting...")
		time.Sleep(1 * time.Second)
	}
	lock.TakeLock()
	fmt.Println("  Action 2 took the lock!")
	time.Sleep(3 * time.Second)
	fmt.Println("  Action 2 executed! ")
	lock.ReleaseLock()
	fmt.Println("  Action 2 released the lock!")

	myChannel <- 2
}

func asyncAction3(myChannel chan int) {
	lock := GetLock()
	for !lock.IsFree() {
		fmt.Println("  Action 3 is waiting...")
		time.Sleep(1 * time.Second)
	}
	lock.TakeLock()
	fmt.Println("  Action 3 took the lock!")
	time.Sleep(3 * time.Second)
	fmt.Println("  Action 3 executed! ")
	lock.ReleaseLock()
	fmt.Println("  Action 3 released the lock!")

	myChannel <- 3
}

func asyncAction4(myChannel chan int) {
	lock := GetLock()
	for !lock.IsFree() {
		fmt.Println("  Action 4 is waiting...")
		time.Sleep(1 * time.Second)
	}
	lock.TakeLock()
	fmt.Println("  Action 4 took the lock!")
	time.Sleep(3 * time.Second)
	fmt.Println("  Action 4 executed! ")
	lock.ReleaseLock()
	fmt.Println("  Action 4 released the lock!")

	myChannel <- 4
}

type lock struct {
	free bool
}

func (l *lock) TakeLock() {
	l.free = false
}

func (l *lock) ReleaseLock() {
	l.free = true
}

func (l *lock) IsFree() bool {
	return l.free == true
}

var instance *lock

func GetLock() *lock {
	if instance != nil {
		return instance
	}

	instance = new(lock)
	instance.free = true
	return instance
}
