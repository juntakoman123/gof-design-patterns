package main

import "fmt"

type Button interface {
	Render()
}

type Checkbox interface {
	Check()
}

type WindowsButton struct{}

func (b *WindowsButton) Render() {
	fmt.Println("Render Windows Button")
}

type WindowsCheckbox struct{}

func (c *WindowsCheckbox) Check() {
	fmt.Println("Check Windows Checkbox")
}

type MacButton struct{}

func (b *MacButton) Render() {
	fmt.Println("Render Mac Button")
}

type MacCheckbox struct{}

func (c *MacCheckbox) Check() {
	fmt.Println("Check Mac Checkbox")
}

type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

type WindowsFactory struct{}

func (f *WindowsFactory) CreateButton() Button {
	return &WindowsButton{}
}

func (f *WindowsFactory) CreateCheckbox() Checkbox {
	return &WindowsCheckbox{}
}

type MacFactory struct{}

func (f *MacFactory) CreateButton() Button {
	return &MacButton{}
}

func (f *MacFactory) CreateCheckbox() Checkbox {
	return &MacCheckbox{}
}

func renderUI(factory GUIFactory) {
	button := factory.CreateButton()
	checkbox := factory.CreateCheckbox()

	button.Render()
	checkbox.Check()
}

func main() {
	os := "mac" // ここを "windows" に変えればUI一式が切り替わる

	var factory GUIFactory
	switch os {
	case "windows":
		factory = &WindowsFactory{}
	case "mac":
		factory = &MacFactory{}
	default:
		panic("unsupported platform")
	}

	renderUI(factory)
}
