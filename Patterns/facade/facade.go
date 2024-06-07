// Package facade реализует паттерн "Фасад" на языке Go.
package facade

import "fmt"

// Подсистема 1
type Subsystem1 struct{}

func (s *Subsystem1) Operation1() string {
	return "Subsystem1: Operation1"
}

// Подсистема 2
type Subsystem2 struct{}

func (s *Subsystem2) Operation2() string {
	return "Subsystem2: Operation2"
}

// Фасад предоставляет упрощенный интерфейс к сложной подсистеме.
type Facade struct {
	subsystem1 *Subsystem1
	subsystem2 *Subsystem2
}

func NewFacade() *Facade {
	return &Facade{
		subsystem1: &Subsystem1{},
		subsystem2: &Subsystem2{},
	}
}

func (f *Facade) Operation() {
	fmt.Println(f.subsystem1.Operation1())
	fmt.Println(f.subsystem2.Operation2())
}

func main() {
	facade := NewFacade()
	facade.Operation()
}
