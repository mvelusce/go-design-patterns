package main

import (
	/*"fmt"
	"github.com/mvelusce/go-design-patterns/launch-system/rocket"
	"time"*/
	"github.com/gin-gonic/gin"
)

func main() {

	/*groundControlChannel := make(chan bool)
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
	}*/

	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//r.Run() // listen and serve on 0.0.0.0:8080
	return r
}
