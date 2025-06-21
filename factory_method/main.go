package main

import "fmt"

type Product interface {
	Use()
}

type ProductA struct{}

func (p *ProductA) Use() {
	fmt.Println("ProductA を使う")
}

type ProductB struct{}

func (p *ProductB) Use() {
	fmt.Println("ProductB を使う")
}

// Creator インターフェース（Factory Methodを定義）
type Creator interface {
	CreateProduct() Product
}

type CreatorA struct{}

func (c *CreatorA) CreateProduct() Product {
	return &ProductA{}
}

type CreatorB struct{}

func (c *CreatorB) CreateProduct() Product {
	return &ProductB{}
}

func main() {
	var creator Creator

	creator = &CreatorA{}
	product1 := creator.CreateProduct()
	product1.Use()

	creator = &CreatorB{}
	product2 := creator.CreateProduct()
	product2.Use()
}
