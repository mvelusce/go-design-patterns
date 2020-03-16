package main

import (
	/*"fmt"
	"github.com/mvelusce/go-design-patterns/launch-system/rocket"
	"time"*/
	"github.com/gin-gonic/gin"
	"net/http"
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
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	//router.Run() // listen and serve on 0.0.0.0:8080
	return router
}
