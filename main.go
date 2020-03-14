package main

import (
	"fmt"
	"github.com/mvelusce/go-design-patterns/singleton"
	"github.com/mvelusce/go-design-patterns/state"
)

func main() {
	fmt.Println("state")
	state.InsertCoin()
	state.InsertCoin()
	state.Push()
	state.Push()
	state.InsertCoin()
	state.Push()

	fmt.Println("singleton")
	var config = singleton.GetInstance()
	fmt.Println(config.GetProperty("p1"))
	fmt.Println(config.GetProperty("p2"))

	var config1 = singleton.GetInstance()

	config.SetProperty("p1", "my value")

	fmt.Println(config1.GetProperty("p1"))
}
