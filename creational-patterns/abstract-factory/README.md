Abstract Factory Method memungkinkan Anda untuk membuat objek-objek yang saling terkait dan bergantung satu sama lain tanpa perlu mengungkapkan detail konkretnya. Ini menciptakan antarmuka abstrak (interface atau kelas abstrak) yang mendefinisikan serangkaian metode untuk membuat objek-objek yang terkait. Kemudian, Anda dapat membuat implementasi konkret dari antarmuka ini untuk menghasilkan kelompok objek yang berbeda sesuai dengan kebutuhan.

Diagram konsep:

```mermaid
classDiagram
    direction LR
    class Client{
        - getFactory : IFactory
    }

    class IFactory{
        <<Interface>>
        + createProductA() : IProductA
        + createProductB() : IProductB
    }

    class ConcreteFactoryA {
        + createProductA() : IProductA
        + createProductB() : IProductB
    }

    class ConcreteFactoryB {
        + createProductA() : IProductA
        + createProductB() : IProductB
    }

    class IProductA{
        
    }
    
    class IProductB{
        
    }

    class ConcreteProductA1{

    }

    class ConcreteProductB1{
        
    }
    
    class ConcreteProductA2{

    }

    class ConcreteProductB2{
        
    }

    Client --> IFactory
    IFactory <|.. ConcreteFactoryA
    IFactory <|.. ConcreteFactoryB

    ConcreteFactoryA <.. IProductA
    ConcreteFactoryA <.. IProductB

    ConcreteFactoryB <.. IProductA
    ConcreteFactoryB <.. IProductB

    IProductA <|.. ConcreteProductA1 
    IProductA <|.. ConcreteProductA2 
    IProductB <|.. ConcreteProductB1 
    IProductB <|.. ConcreteProductB2 

```

Diagram implementasi:

```mermaid
classDiagram
direction LR
    class Client{
        - getGUIFactory(factody string) : IGUIFactory
    }

    class IGUIFactory{
        <<Interface>>
        + createButton() : IButton
        + createCheckbox() : ICheckbox
    }

    class MacGUIFactory {
        + createButton() : IButton
        + createCheckbox() : ICheckbox
    }

    class WinGUIFactory {
        + createButton() : IButton
        + createCheckbox() : ICheckbox
    }

    class IButton{
        <<Interface>>
        + onClick()
    }
    
    class ICheckbox{
        <<Interface>>
        + onClick()
    }

    class MacButton{
        + onClick()
    }

    class WinCheckbox{
        + onClick()
    }
    
    class WinButton{
        + onClick()
    }

    class MacCheckbox{
        + onClick()
    }

    Client --> IGUIFactory
    IGUIFactory <|.. MacGUIFactory
    IGUIFactory <|.. WinGUIFactory

    MacGUIFactory <.. IButton
    MacGUIFactory <.. ICheckbox

    WinGUIFactory <.. IButton
    WinGUIFactory <.. ICheckbox

    IButton <|.. MacButton
    IButton <|.. WinButton
    ICheckbox <|.. MacCheckbox
    ICheckbox <|.. WinCheckbox
```

```go
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

```

## Kapan?

 - Ketika kode Anda perlu bekerja dengan berbagai kelompok produk terkait, namun Anda tidak ingin kode tersebut bergantung pada kelas konkret dari produk tersebut—mereka mungkin tidak diketahui sebelumnya atau Anda hanya ingin memungkinkan perluasan di masa mendatang.
 - Anda memiliki sebuah kelas yang harus melakukan satu pekerjaan utama (tanggung jawab utama), tetapi juga melakukan pekerjaan lain yang tidak seharusnya menjadi tanggung jawabnya.

## Referensi
- https://refactoring.guru/design-patterns/abstract-factory
- https://golangbyexample.com/abstract-factory-design-pattern-go/
