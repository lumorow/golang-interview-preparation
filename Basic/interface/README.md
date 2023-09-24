## Interface in GO

<h1 align="center"><img class="goldT" src="../../img/golang-any-interface-fs8.webp" ></h1>

The interface in the `Go` language is structured like this:
```golang
// https://golang.org/src/runtime/runtime2.go

type iface struct {
    tab  *itab
    data unsafe.Pointer
}
```
In the `tab` field we store information about the specific type of object that has been converted to an interface. And in the `data` field - a link to the real memory area,
which contains the data of the original object.

An interface implements that object that defines all the methods specified in the interface.

An empty interface type does not define methods. He has no rules. And therefore any object satisfies the empty interface.
The empty interface type `interface{}` is a kind of wild card. If you met him in an ad
(variable, function parameter or structure field), then you can use an object of any type.

## Benefits of using interfaces

- Interfaces help reduce duplication, that is, the amount of boilerplate code.
- They make it easier to use stubs instead of real objects in unit tests.
- As an architectural tool, interfaces help decouple parts of your codebase.

## Type of object wrapped in interface
You can determine the type of an object wrapped in an empty interface using the `switch` construct
```golang
func do(obj interface{}) {
    switch v := obj.(type) {
    case int:
        fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
        fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
        fmt.Printf("I don't know about type %T!\n", v)
    }
}
```
You can also use the `reflect` library. When we try to get
The actual source type of the interface object, reflect goes to the `tab` field and gets the type information there.
```golang
reflect.TypeOf(obj).String()
```
We can convert our interface object back to its original type using type assertion syntax:
```golang
obj, ok := interfaceName.(structName)  
fmt.Printf("%#v\n", obj)
fmt.Printf("%t\n", ok)
```

If we compare interface with `nil`, it will be `true` until we wrap an object in it. Even if we put an empty interface
pointer to an object, then interface will no longer be `nil`, since one of the fields of the structure will point to the type of the object (`tab`)
## Resources
***
- [Разбираемся с интерфейсами в Go (Блог компании VK)](https://habr.com/ru/companies/vk/articles/463063/)
- [Интерфейсы в Go — как красиво выстрелить себе в ногу](https://habr.com/ru/articles/597461/)
- [Тайные знания о GoLang, которые от вас скрывали (Николай Тузов — Golang)](https://www.youtube.com/watch?v=-cX0CqG6rgA&ab_channel=%D0%9D%D0%B8%D0%BA%D0%BE%D0%BB%D0%B0%D0%B9%D0%A2%D1%83%D0%B7%D0%BE%D0%B2%E2%80%94Golang)

## README.md
***

- eng [English](https://github.com/lumorow/golang-interview-preparation/blob/main/Basic/interface/README.md)
- ru [Русский](https://github.com/lumorow/golang-interview-preparation/blob/main/Basic/interface/README.ru.md)

