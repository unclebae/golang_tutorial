package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go say("Hay")
	wg.Add(1)
	go say("Kido")
	// wg.Wait()

	wg.Add(1)
	go say("Hi")
	wg.Wait()
	// time.Sleep(time.Second * 5)
	// fmt.Println("End")
}
