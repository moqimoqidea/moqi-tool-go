// Package gocron
// @author moqi
// On 2022/10/30 16:40:22
// see: https://github.com/go-co-op/gocron
package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

func hello(name string) {
	message := fmt.Sprintf("Hi, %v", name)
	fmt.Println(message)
}

func runCronJobs() {
	s := gocron.NewScheduler(time.UTC)

	_, err := s.Every(10).Seconds().Do(func() {
		hello("moqi")
	})

	if err != nil {
		return
	}

	s.StartBlocking()
}

func main() {
	runCronJobs()
	fmt.Println("start gocron")
}
