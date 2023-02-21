package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(time.Second * 5)

	for {
		select {
		case v := <-timer.C:
			fmt.Println(v)
		}
	}
}
