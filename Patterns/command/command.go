// Package command реализует паттерн "Команда" на языке Go.
package command

import "fmt"

// Интерфейс Команды
type Command interface {
	Execute()
}

// Конкретная Команда A
type ConcreteCommandA struct {
	receiver *Receiver
}

func (c *ConcreteCommandA) Execute() {
	c.receiver.ActionA()
}

// Конкретная Команда B
type ConcreteCommandB struct {
	receiver *Receiver
}

func (c *ConcreteCommandB) Execute() {
	c.receiver.ActionB()
}

// Получатель
type Receiver struct{}

func (r *Receiver) ActionA() {
	fmt.Println("ActionA executed")
}

func (r *Receiver) ActionB() {
	fmt.Println("ActionB executed")
}

// Инициатор
type Invoker struct {
	commands []Command
}

func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) ExecuteCommands() {
	for _, command := range i.commands {
		command.Execute()
	}
}

func main() {
	receiver := &Receiver{}
	commandA := &ConcreteCommandA{receiver: receiver}
	commandB := &ConcreteCommandB{receiver: receiver}

	invoker := &Invoker{}
	invoker.StoreCommand(commandA)
	invoker.StoreCommand(commandB)

	invoker.ExecuteCommands()
}
