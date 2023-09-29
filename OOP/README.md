# Object Oriented Programming in golang

<h1 align="center"><img class="goldT" src="../img/oopgo.jpeg" ></h1>

## Difference between new and make
`make` allocates memory for reference data types ( slice , map , chan ) and also initializes
their underlying data structures.

`new` returns only pointers to the initialized memory and sets it to zero for type `T`,
that is, `0` for `int`, `''` for `string` and `nil` for reference types (`slice`, `map`, `chan`)

## Class. Creating a Class Object

In golang there are no such classes, but structures are used. Structures also have fields and methods, you can write
its own constructor, but it will be represented as a regular method.

```golang
type person struct {
    age    int
    height int
    weight int
    male   string
    name   string
}

func NewPerson(age, height, weight int, male, name string) *person {
    return &person{
    age:    age,
    height: height,
    weight: weight,
    male:   male,
    name:   name,
    }
}
```

## Visibility of structure and fields
`Go` uses the case rule of the first letter of the name - if the name begins with a capital letter -
this is `public` access, if with lowercase it is `private`.
From the example above we can see that our `person` structure is not exported (private),
therefore it is only available in the given package where we initialized it. But we created a constructor,
which can create this object for us according to the given instructions inside it.

## Implementation of concepts in golang:
### Inheritance

Golang does not have inheritance, but uses `inlining`.
The simplest use case for inheritance is that a child type must have access to the fields and methods of the parent type.
This is done in Go through inlining. The base structure is embedded in the child structure, after which the base fields
and methods can be directly accessed by the child structure.

More details:  
[OOP: Inheritance in GOLANG complete guide](https://golangbyexample.com/oop-inheritance-golang-complete/)

### Polymorphism
#### Compile Time Polymorphism - not possible in golang

With compile-time polymorphism, the compiler decides which function to call.
Examples of compile-time polymorphism would be:
- method/function overloading: there is more than one method/function with the same name,
  but with different signatures or perhaps different return types;
- operator overloading: the same operator is used to work with different data types.

Go does not support method overloading.
Go also does not support operator overloading.
***
#### Run Time Polymorphism - achieved through interfaces

Runtime polymorphism means that the decision about which function to call is made at runtime.
In [oop.go](https://github.com/lumorow/golang-interview-preparation/blob/main/OOP/oop.go) `interface Grower`
creates a contract with
`person` and `dog` structures, which have different implementations of the `growUp()` function.

More details:  
[OOP: Polymorphism in Go Complete Guide](https://golangbyexample.com/oop-polymorphism-in-go-complete-guide/)

### Abstract classes

An interface in Go does not contain fields, nor does it allow you to define methods within it. Any type must implement all methods of an interface in order to have that interface's type. There are situations where it is useful to have a default method implementation as well as default fields in Go. Before we understand how this can be done, let's first understand the requirements for an abstract class:

An abstract class must have default fields
An abstract class must have a default method
It should not be possible to directly instantiate an abstract class

More details:  
[Abstract Class in GO: Complete Guide](https://golangbyexample.com/go-abstract-class/)

### Encapsulation

Golang provides package-level encapsulation. There are no public, private, or protected keywords in Go. The only mechanism for controlling visibility is the use of uppercase and lowercase letters.

Identifiers with a capital letter are exported. A capital letter means it is an exported identifier.
Identifiers with lowercase letters are not exported. Lowercase letters indicate that the identifier is not exported and will only be accessible from the same package.
There are five types of IDs that can be exported.

- Structure
- Structure method
- Structure field
- Function
- Variable

More details:  
[Encapsulation in Go (Golang)](https://golangbyexample.com/encapsulation-in-go/)

## Resources

- [Объектно-ориентированное программирование в Golang](https://medium.com/nuances-of-programming/%D0%BE%D0%B1%D1%8A%D0%B5%D0%BA%D1%82%D0%BD%D0%BE-%D0%BE%D1%80%D0%B8%D0%B5%D0%BD%D1%82%D0%B8%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%BD%D0%BE%D0%B5-%D0%BF%D1%80%D0%BE%D0%B3%D1%80%D0%B0%D0%BC%D0%BC%D0%B8%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5-%D0%B2-golang-52f36f2fa837)

## README.md

- eng [English](https://github.com/lumorow/golang-interview-preparation/blob/main/OOP/README.md)
- ru [Русский](https://github.com/lumorow/golang-interview-preparation/blob/main/OOP/readme/README.ru.md)