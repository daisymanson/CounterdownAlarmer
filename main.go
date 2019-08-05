package main

import (
	"fmt"
	"time"
)

type Alert struct{}

func (alert *Alert) Message(t interface{}) string {
	return "time's up!"
}

func (alert *Alert) Tick(n time.Duration) {
	ticker := time.NewTicker(time.Second * n)
	for {
		select {
		case t := <-ticker.C:
			fmt.Println(alert.Message(t))
			ticker.Stop()
		}
	}
}

func main() {
	alert := Alert{}
	// 選擇秒數
	alert.Tick(5)
}
