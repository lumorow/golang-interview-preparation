## Срезы в Go

Срезы это структура, о которой мы поговорим далее.
<h1 align="center"><img class="goldT" src="../../img/gophslice.svg" width="400" height="250"></h1>

### Обявление
***
- Первый вариант
```golang
planets := []string{
    "Меркурий",
    "Венера",
    "Земля",
    "Марс",
    "Юпитер",
    "Сатурн",
    "Уран",
    "Нептун",
}
```

- Второй вариант
```golang
planets := make([]string, len, cap) // через make
```

- Третий вариант
```golang
nums := [...]int{
1,
2,
3,
4,
5,
6
}

numsSl := nums[:] // через композитный литерал объявляем срез из массива
```
### Определение
***
От массива срез отличает следующее:  

- Срез является ссылкой на массив;
- Срез может динамически аллоцировать память и менять свой размер;
- Срез представляется следующей структурой в исходниках `GO`

```golang
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
```
`len (length, длина)` — текущая длина среза, `cap` (capacity, вместимость) — длина внутреннего массива.

Оба эти параметра можно задать при вызове функции make:
```golang
sl := make(
[]int,
10, // len
10, // cap
)
```
`Cap` — ключевой параметр для аллокации памяти, влияет на производительность вставки в срез.

### Рассмотрим поведение среза при увеличении его размера.
***
```golang
a := []int{99}
b := a[0:1]
b[0] = 0

fmt.Println(a)
// [0]

a[0] = 99

fmt.Println(b)
// [99]
```

Мы получили срез `b` из среза `a`. Далее видим, что изменение в одном срезе влияют на другой. Оба среза ссылаются на один и тот же массив.
Теперь дополним пример вставкой элементов в срез `a`:
Рассмотрим поведение среза при увеличении его размера.
```golang
a := []int{99}
b := a[0:1]

a = append(a, 100)
a[0] = 0

fmt.Println(a, b)
// [0 100] [99]
```
Теперь значения у элементов срезов разные. Теперь срезы ссылаются на разный массив.

За это отвечает функция `growslice` из исходников.
В результате изменения `cap` среза всегда произойдет копирование данных массива:
```golang
memmove(p, old.array, lenmem)
```
Изменяться `cap` будет следующем образом
```golang
const threshold = 256
    if old.cap < threshold {
        newcap = doublecap
    } else {
        // Check 0 < newcap to detect overflow
        // and prevent an infinite loop.
        for 0 < newcap && newcap < cap {
        // Transition from growing 2x for small slices
        // to growing 1.25x for large slices. This formula
        // gives a smooth-ish transition between the two.
        newcap += (newcap + 3*threshold) / 4
	}
```

## Удаление элемента из среза
***
Чтобы удалить `i-ый` элемент из среза, нужно выделить элементы до него и после него,  
а потом объединить два новых среза в один.
```golang
slice = append(slice[:i], slice[i+1:]...)
```

## Дополнительный материал
***
- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
- [Массивы и срезы в Go](https://www.digitalocean.com/community/tutorials/understanding-arrays-and-slices-in-go-ru)
- [GoLang Slice в деталях, простым языком (Николай Тузов — Golang)](https://www.youtube.com/watch?v=10LW7NROfOQ&ab_channel=%D0%9D%D0%B8%D0%BA%D0%BE%D0%BB%D0%B0%D0%B9%D0%A2%D1%83%D0%B7%D0%BE%D0%B2%E2%80%94Golang)

## README.md
***

- eng [English](https://github.com/lumorow/golang-interview-preparation/blob/main/Basic/slice/README.md)
- ru [Русский](https://github.com/lumorow/golang-interview-preparation/blob/main/Basic/slice/README.ru.md)