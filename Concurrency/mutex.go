package main

import (
	"fmt"
	"sync"
)

var (
	balance int
	mutex   sync.Mutex
)

func init() {
	balance = 1000
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func(val int, wg *sync.WaitGroup) {
		fmt.Println("Withdraw")
		mutex.Lock()
		fmt.Println(val, balance)
		balance -= val
		fmt.Println(val, balance)
		mutex.Unlock()
		wg.Done()
	}(700, &wg)

	go func(val int, wg *sync.WaitGroup) {
		fmt.Println("Deposit")
		mutex.Lock()
		fmt.Println(val, balance)
		balance += val
		fmt.Println(val, balance)
		mutex.Unlock()
		wg.Done()
	}(900, &wg)

	wg.Wait()
	fmt.Println("Final", balance)

}
