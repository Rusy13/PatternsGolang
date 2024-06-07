// Package strategy реализует паттерн "Стратегия" на языке Go.
package strategy

import "fmt"

// Интерфейс Стратегии
type Strategy interface {
	Execute()
}

// Конкретная Стратегия A
type ConcreteStrategyA struct{}

func (s *ConcreteStrategyA) Execute() {
	fmt.Println("Executing ConcreteStrategyA")
}

// Конкретная Стратегия B
type ConcreteStrategyB struct{}

func (s *ConcreteStrategyB) Execute() {
	fmt.Println("Executing ConcreteStrategyB")
}

// Контекст использует Стратегию
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy() {
	c.strategy.Execute()
}

func main() {
	context := &Context{}

	strategyA := &ConcreteStrategyA{}
	context.SetStrategy(strategyA)
	context.ExecuteStrategy()

	strategyB := &ConcreteStrategyB{}
	context.SetStrategy(strategyB)
	context.ExecuteStrategy()
}
