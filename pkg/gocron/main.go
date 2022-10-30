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
	// 3
	s := gocron.NewScheduler(time.UTC)

	// 4
	_, err := s.Every(1).Seconds().Do(func() {
		hello("moqi")
	})

	if err != nil {
		return
	}

	// 5
	s.StartBlocking()
}

// 6
func main() {
	runCronJobs()
	fmt.Println("start gocron")
}
