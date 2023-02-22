package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	URL := "http://clock.meaningful-godiva.koyeb:8000"
	duration := time.Minute * 10
	timer := time.NewTicker(duration)

	fmt.Printf("Starting a tick every %s\n", duration)

	for {
		select {
		case <-timer.C:
			resp, err := http.Get(URL)
			if err == nil {
				b, err := io.ReadAll(resp.Body)
				if err == nil {
					fmt.Printf("Call : %s, status code : %s\n", URL, resp.Status)
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
