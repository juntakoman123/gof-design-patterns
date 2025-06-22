package main

import "fmt"

type Command interface {
	Execute()
}

type Light struct{}

func (l *Light) On() {
	fmt.Println("ライトをつけました")
}

func (l *Light) Off() {
	fmt.Println("ライトを消しました")
}

type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(cmd Command) {
	r.command = cmd
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func main() {
	light := &Light{}
	turnOn := &LightOnCommand{light: light}
	turnOff := &LightOffCommand{light: light}

	remote := &RemoteControl{}

	remote.SetCommand(turnOn)
	remote.PressButton() // ライトをつけました

	remote.SetCommand(turnOff)
	remote.PressButton() // ライトを消しました
}
