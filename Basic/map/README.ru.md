## Карты в GO
***
Хэш таблица - это структура данных, которая позволяет хранить пары ключ-значение, и, как правило, обладающая функциями:

- Маппинга: map(key) → value
- Вставки: insert(map, key, value)
- Удаления: delete(map, key)
- Поиска: lookup(key) → value

### Обявление
***
```golang
m := make(map[key_type]value_type)
```
```golang
m := new(map[key_type]value_type)
```
```golang
var m map[key_type]value_type
```
```golang
m := map[key_type]value_type{key1: val1, key2: val2}
```

### Основные операции
***
- Вставка:
```golang
m[key] = value
```
- Удаление:
```golang
delete(m, key)
```
- Поиск:
```golang
value = m[key]
value, ok = m[key]
```

### Обход таблицы в go
***
Мапа в Go unordered, то есть не упорядоченная, при переборе будет всегда разный порядок элементов.  
Место поиска определяется рандомно (исходник кода рантайма языка):
```golang
// mapiterinit initializes the hiter struct used for ranging over maps.
func mapiterinit(t *maptype, h *hmap, it *hiter) {...
// decide where to start
r := uintptr(fastrand())
...
it.startBucket = r & bucketMask(h.B)...}
```
Пример:
```golang
package main

import "fmt"

func main() {
    m := map[int]bool{}
    for i := 0; i < 5; i++ {
        m[i] = ((i % 2) == 0)
    }
    for k, v := range m {
        fmt.Printf("key: %d, value: %t\n", k, v)
    }
}
```
Запуск 1:

key: 3, value: false  
key: 4, value: true  
key: 0, value: true  
key: 1, value: false  
key: 2, value: true  

Запуск 2:

key: 4, value: true  
key: 0, value: true  
key: 1, value: false  
key: 2, value: true  
key: 3, value: false  

### Поиск в таблице Go
Благодаря `«multiple assignment»` мы можем проверять с помощью второй переменной `ok` наличие ключа в мапе.  
При отсутствии ключа будет возвращаться `«нулевое значение типа»` и `false`.
***
```golang
package main

import (
"fmt"
)

func main() {
    m := map[int]int{0: 50, 1: -100}
    m2, ok := m[2]
	m3 := m[3]
    if !ok {
        m2 = 20
    }
    fmt.Println(m, m[0], m[1], m2, m3) // map[0:0 1:10] 0 10 20 0
}
```

## Дополнительный материал
***
- [Go maps in action](https://go.dev/blog/maps)
- [Знакомство с картами в Go](https://www.digitalocean.com/community/tutorials/understanding-maps-in-go-ru)
- [Хэш таблицы в Go. Детали реализации](https://habr.com/ru/articles/457728/)
- [Как на самом деле устроен тип Map в Golang? (Николай Тузов — Golang)](https://www.youtube.com/watch?v=P_SXTUiA-9Y&ab_channel=%D0%9D%D0%B8%D0%BA%D0%BE%D0%BB%D0%B0%D0%B9%D0%A2%D1%83%D0%B7%D0%BE%D0%B2%E2%80%94Golang)

## README.md
***

- eng [English](https://github.com/lumorow/golang-interview-preparation/blob/main/Basic/map/README.md)
- ru [Русский](https://github.com/lumorow/golang-interview-preparation/blob/main/Basic/map/README.ru.md)