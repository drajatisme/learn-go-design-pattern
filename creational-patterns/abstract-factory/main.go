package main

import (
	"fmt"
	"runtime"
)

type IButton interface {
	onClick()
}

type WinButton struct{}

func (w *WinButton) onClick() {
	fmt.Println("Win Button Clicked")
}

type MacButton struct{}

func (w *MacButton) onClick() {
	fmt.Println("Mac Button Clicked")
}

type ICheckbox interface {
	onClick()
}

type WinCheckbox struct{}

func (w *WinCheckbox) onClick() {
	fmt.Println("Win Checkbox Clicked")
}

type MacCheckbox struct{}

func (w *MacCheckbox) onClick() {
	fmt.Println("Mac Checkbox Clicked")
}

type IGUIFactory interface {
	createButton() IButton
	createCheckbox() ICheckbox
}

type WinGUIFactory struct{}

func (w *WinGUIFactory) createButton() IButton {
	fmt.Println("Creating Windows Button")
	return &WinButton{}
}

func (w *WinGUIFactory) createCheckbox() ICheckbox {
	fmt.Println("Creating Windows Checkbox")
	return &WinCheckbox{}
}

type MacGUIFactory struct{}

func (w *MacGUIFactory) createButton() IButton {
	fmt.Println("Creating Mac Button")
	return &WinButton{}
}

func (w *MacGUIFactory) createCheckbox() ICheckbox {
	fmt.Println("Creating Mac Checkbox")
	return &WinCheckbox{}
}

// Client
func getGUIFactory(name string) IGUIFactory {
	switch name {
	case "windows":
		return &WinGUIFactory{}
	case "darwin":
		return &MacGUIFactory{}
	default:
		return nil
	}
}

func main() {
	guiFactory := getGUIFactory(runtime.GOOS)

	button := guiFactory.createButton()
	checkbox := guiFactory.createCheckbox()

	button.onClick()
	checkbox.onClick()
}
