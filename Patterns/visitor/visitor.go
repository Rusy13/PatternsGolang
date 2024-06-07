// Package visitor реализует паттерн "Посетитель" на языке Go.
package visitor

import "fmt"

// Интерфейс Посетителя
type Visitor interface {
	VisitConcreteElementA(*ConcreteElementA)
	VisitConcreteElementB(*ConcreteElementB)
}

// Интерфейс Элемента
type Element interface {
	Accept(Visitor)
}

// Конкретный Элемент A
type ConcreteElementA struct{}

func (e *ConcreteElementA) Accept(v Visitor) {
	v.VisitConcreteElementA(e)
}

func (e *ConcreteElementA) OperationA() string {
	return "ConcreteElementA"
}

// Конкретный Элемент B
type ConcreteElementB struct{}

func (e *ConcreteElementB) Accept(v Visitor) {
	v.VisitConcreteElementB(e)
}

func (e *ConcreteElementB) OperationB() string {
	return "ConcreteElementB"
}

// Конкретный Посетитель
type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(e *ConcreteElementA) {
	fmt.Println("Visited", e.OperationA())
}

func (v *ConcreteVisitor) VisitConcreteElementB(e *ConcreteElementB) {
	fmt.Println("Visited", e.OperationB())
}

func main() {
	elements := []Element{
		&ConcreteElementA{},
		&ConcreteElementB{},
	}

	visitor := &ConcreteVisitor{}
	for _, element := range elements {
		element.Accept(visitor)
	}
}
