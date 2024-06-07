// Package state реализует паттерн "Состояние" на языке Go.
package state

import "fmt"

// Интерфейс Состояния
type State interface {
	Handle(context *Context)
}

// Контекст
type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle(c)
}

// Конкретное Состояние A
type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle(context *Context) {
	fmt.Println("ConcreteStateA handling request")
	context.SetState(&ConcreteStateB{})
}

// Конкретное Состояние B
type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle(context *Context) {
	fmt.Println("ConcreteStateB handling request")
	context.SetState(&ConcreteStateA{})
}

func main() {
	context := &Context{state: &ConcreteStateA{}}

	context.Request()
	context.Request()
	context.Request()
	context.Request()
}
