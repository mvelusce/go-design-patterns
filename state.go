package main

import (
	"fmt"
)

type state interface {
	insertCoin()
	push()
}

var turnstileState state = lockState{}

// InsertCoin insert coin
func InsertCoin() {
	fmt.Println("Insert coin")
	turnstileState.insertCoin()
}

// Push push
func Push() {
	fmt.Println("Push")
	turnstileState.push()
}

type lockState struct {}

type openState struct {}

func (s lockState) insertCoin() {
	fmt.Println("Unlocking turnstile")
	turnstileState = openState{}
}

func (s lockState) push() {
	fmt.Println("Turnstile is locked")
}

func (s openState) insertCoin() {
	fmt.Println("Turnstile is unlokced already")
}

func (s openState) push() {
	fmt.Println("Opening turnstile")
	turnstileState = lockState{}
}

func main() {
	InsertCoin()
	InsertCoin()
	Push()
	Push()
	InsertCoin()
	Push()
}
