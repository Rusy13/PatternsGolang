// Package factory_method реализует паттерн "Фабричный метод" на языке Go.
package factory_method

import "fmt"

// Продукт
type Product interface {
	Use()
}

// Конкретный Продукт A
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() {
	fmt.Println("Using ConcreteProductA")
}

// Конкретный Продукт B
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() {
	fmt.Println("Using ConcreteProductB")
}

// Интерфейс Создателя
type Creator interface {
	CreateProduct() Product
}

// Конкретный Создатель A
type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// Конкретный Создатель B
type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	creatorA := &ConcreteCreatorA{}
	productA := creatorA.CreateProduct()
	productA.Use()

	creatorB := &ConcreteCreatorB{}
	productB := creatorB.CreateProduct()
	productB.Use()
}
