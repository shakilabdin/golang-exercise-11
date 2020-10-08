package main

import (
	"fmt"
	"sync"
	"runtime"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	fmt.Println("routines", runtime.NumGoroutine())
	
	go foo()
	go bar()
	
	fmt.Println("first line")
	fmt.Println("routines", runtime.NumGoroutine())
	
	wg.Wait()
	
	fmt.Println("routines", runtime.NumGoroutine())
}

func foo() {
	fmt.Println("Hello")
	wg.Done()
}

func bar()  {
	fmt.Println("World")
	wg.Done()
}