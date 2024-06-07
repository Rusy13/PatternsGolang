// Package chain_of_responsibility реализует паттерн "Цепочка вызовов" на языке Go.
package chain_of_responsibility

import "fmt"

// Интерфейс Обработчика
type Handler interface {
	SetNext(Handler) Handler
	Handle(request string)
}

// Базовый Обработчик
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(next Handler) Handler {
	h.next = next
	return next
}

func (h *BaseHandler) Handle(request string) {
	if h.next != nil {
		h.next.Handle(request)
	}
}

// Конкретный Обработчик A
type ConcreteHandlerA struct {
	BaseHandler
}

func (h *ConcreteHandlerA) Handle(request string) {
	if request == "A" {
		fmt.Println("ConcreteHandlerA handled request")
	} else {
		h.BaseHandler.Handle(request)
	}
}

// Конкретный Обработчик B
type ConcreteHandlerB struct {
	BaseHandler
}

func (h *ConcreteHandlerB) Handle(request string) {
	if request == "B" {
		fmt.Println("ConcreteHandlerB handled request")
	} else {
		h.BaseHandler.Handle(request)
	}
}

func main() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	handlerA.SetNext(handlerB)

	handlerA.Handle("A")
	handlerA.Handle("B")
	handlerA.Handle("C")
}
