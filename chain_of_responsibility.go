package main

import (
	"fmt"
)

type handler interface {
	SetNext(handler handler)
	Handle()
}

type handler1 struct {
	next handler
}

func (h *handler1) SetNext(handler handler) {
	h.next = handler
}

func (h *handler1) Handle() {
	fmt.Println("Print handler1")
}

type handler2 struct {
	next handler
}

func (h handler2) SetNext(handler handler) {
	h.next = handler
}

func (h handler2) Handle() {
	fmt.Println("Print handler2")
}

func runChain() {
	h1 := &handler1{}
	h2 := &handler2{}
	h1.SetNext(h2)

	start := h1
	start.Handle()
	next := start.next
	if next != nil {
		next.Handle()
	}
}
