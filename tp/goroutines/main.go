package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var nbGoroutines int = 1000000
	// var wg sync.WaitGroup
	// wg.Add(nbGoroutines)
	// for i := 0; i < nbGoroutines; i++ {
	// 	go DoSomethingGreat(&wg)
	// }
	// wg.Wait()
	fmt.Println("start")
	for i := 0; i < nbGoroutines; i++ {
		go DoSomethingGreatWithoutWaitgroup()
	}
	fmt.Println("end")
}

func DoSomethingGreat(wg *sync.WaitGroup) {
	fmt.Println("Hayyy !")
	time.Sleep(time.Second * 3)
	fmt.Println("this is done")
	wg.Done()
}

func DoSomethingGreatWithoutWaitgroup() {
	fmt.Println("Hayyy !")
	time.Sleep(time.Second * 3)
	fmt.Println("this is done")
}
