package main

import (
	"fmt"
	"time"
)

func main() {
	duration := time.Second * 5
	timer := time.NewTicker(duration)

	fmt.Printf("Starting a tick every %s\n", duration)

	for {
		select {
		case v := <-timer.C:
			fmt.Println(v.Format("2006-02-01 15:04:05"))
		}
	}
}
