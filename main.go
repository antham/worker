package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	duration := time.Second * 5
	timer := time.NewTicker(duration)

	fmt.Printf("Starting a tick every %s\n", duration)

	for {
		select {
		case v := <-timer.C:
			fmt.Println(v.Format("2006-01-02 15:04:05"))
			resp, err := http.Get("demo2.meaningful-godiva.koyeb")
			if err == nil {
				b, err := io.ReadAll(resp.Body)
				if err == nil {
					fmt.Println(resp.Status)
					fmt.Printf("%s\n", b[0:50])
				}
			}
		}
	}
}
