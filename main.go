package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	duration := time.Minute * 1
	timer := time.NewTicker(duration)

	fmt.Printf("Starting a tick every %s\n", duration)

	for {
		select {
		case <-timer.C:
			resp, err := http.Get("http://clock.meaningful-godiva.koyeb:8000")
			if err == nil {
				b, err := io.ReadAll(resp.Body)
				if err == nil {
					fmt.Println(resp.Status)
					fmt.Printf("%s\n", string(b))
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		}
	}
}
