package main

import "fmt"

type State interface {
	PressSwitch(c *Context)
}

type Context struct {
	state State
}

func (c *Context) SetState(s State) {
	c.state = s
}

func (c *Context) PressSwitch() {
	c.state.PressSwitch(c)
}

type OffState struct{}

func (s *OffState) PressSwitch(c *Context) {
	fmt.Println("スイッチオン：電源が入りました")
	c.SetState(&OnState{})
}

type OnState struct{}

func (s *OnState) PressSwitch(c *Context) {
	fmt.Println("スイッチオフ：電源を切りました")
	c.SetState(&OffState{})
}

func main() {
	ctx := &Context{}
	ctx.SetState(&OffState{})

	ctx.PressSwitch() // オンになる
	ctx.PressSwitch() // オフになる
	ctx.PressSwitch() // オンになる
}
