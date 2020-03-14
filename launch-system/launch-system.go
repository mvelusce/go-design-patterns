package main

import (
	"fmt"
	"github.com/mvelusce/go-design-patterns/launch-system/rocket"
	"time"
)

func main() {

	groundControlChannel := make(chan bool)
	rocketChannel := make(chan string)

	go rocket.LaunchRocket(groundControlChannel, rocketChannel)

	fmt.Println("Launching rocket")
	groundControlChannel <- true

	for {
		rocketStatus := <-rocketChannel
		fmt.Println(rocketStatus)
		if rocketStatus == rocket.LandingSafeState {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}
