package main

import (
	"fmt"
	"time"
)

func backgroundTimer(sent string) {

	ticker := time.NewTicker(2 * time.Second)
	for _ = range ticker.C {
		fmt.Println(sent)
	}
}

func main() {
	go backgroundTimer("Tick")
	time.Sleep(1 * time.Second)
	go backgroundTimer("Tock")

	select {}
}
