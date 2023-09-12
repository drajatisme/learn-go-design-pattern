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

## Kapan?

 - Ketika kode Anda perlu bekerja dengan berbagai kelompok produk terkait, namun Anda tidak ingin kode tersebut bergantung pada kelas konkret dari produk tersebut—mereka mungkin tidak diketahui sebelumnya atau Anda hanya ingin memungkinkan perluasan di masa mendatang.
 - Anda memiliki sebuah kelas yang harus melakukan satu pekerjaan utama (tanggung jawab utama), tetapi juga melakukan pekerjaan lain yang tidak seharusnya menjadi tanggung jawabnya.

## Referensi
- https://refactoring.guru/design-patterns/abstract-factory
- https://golangbyexample.com/abstract-factory-design-pattern-go/
