// Package builder реализует паттерн "Строитель" на языке Go.
package builder

import "fmt"

// Продукт
type Product struct {
	Part1 string
	Part2 string
	Part3 string
}

// Интерфейс Строителя
type Builder interface {
	BuildPart1()
	BuildPart2()
	BuildPart3()
	GetResult() Product
}

// Конкретный Строитель
type ConcreteBuilder struct {
	product Product
}

func (b *ConcreteBuilder) BuildPart1() {
	b.product.Part1 = "Part1"
}

func (b *ConcreteBuilder) BuildPart2() {
	b.product.Part2 = "Part2"
}

func (b *ConcreteBuilder) BuildPart3() {
	b.product.Part3 = "Part3"
}

func (b *ConcreteBuilder) GetResult() Product {
	return b.product
}

// Директор
type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) Construct() {
	d.builder.BuildPart1()
	d.builder.BuildPart2()
	d.builder.BuildPart3()
}

func main() {
	builder := &ConcreteBuilder{}
	director := NewDirector(builder)
	director.Construct()
	product := builder.GetResult()
	fmt.Println(product)
}
