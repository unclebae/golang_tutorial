package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup", r)
	}

}

func say(s string) {

	// defer wg.Done()
	defer cleanup()
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
		if i == 2 {
			panic("Oh dear, a 2")
		}
	}
}

func foo() {
	defer fmt.Println("Done!")
	defer fmt.Println("Are we done?")
	fmt.Println("Doing some stuff, who knows what?")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	// foo()

	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("There")

	wg.Wait()
}
