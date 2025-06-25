package main

import "fmt"

// Strategy インターフェース
type Strategy interface {
	Calculate(a, b int) int
}

// 加算戦略
type AddStrategy struct{}

func (s AddStrategy) Calculate(a, b int) int {
	return a + b
}

// 乗算戦略
type MultiplyStrategy struct{}

func (s MultiplyStrategy) Calculate(a, b int) int {
	return a * b
}

// Context
type Calculator struct {
	strategy Strategy
}

func (c *Calculator) SetStrategy(s Strategy) {
	c.strategy = s
}

func (c *Calculator) Execute(a, b int) int {
	return c.strategy.Calculate(a, b)
}

func main() {
	calc := &Calculator{}

	calc.SetStrategy(AddStrategy{})
	fmt.Println("Add:", calc.Execute(3, 4)) // 7

	calc.SetStrategy(MultiplyStrategy{})
	fmt.Println("Multiply:", calc.Execute(3, 4)) // 12
}
