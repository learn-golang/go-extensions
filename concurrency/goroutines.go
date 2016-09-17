package main

import (
	"sync"
	"time"
)


func sleeper_nested(wg *sync.WaitGroup, first time.Duration, second time.Duration) {
	go func() {
		time.Sleep(first)
		wg.Done()
	}()
	go func() {
		time.Sleep(second)
		wg.Done()
	}()
	defer wg.Done()
}

func sleeper(wg *sync.WaitGroup, times time.Duration) {
	time.Sleep(times)
	defer wg.Done()
}


func main() {
	var wg sync.WaitGroup
	//wg.Add(3)
	//go sleeper_nested(&wg, time.Second * 5, time.Second * 2)
	wg.Add(2)
	go sleeper(&wg, time.Second * 5)
	go sleeper(&wg, time.Second * 2)
	wg.Wait()
}
