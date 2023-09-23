package main

import "fmt"

func fn1(m map[int]int) {
	m = make(map[int]int)
	fmt.Println("m == nil in fn?:", m == nil) // m == nil in fn?: false
}

func fn2(m map[string]int) {
	m["c"] = 3
	m["d"] = 4
	m["e"] = 5
	m["f"] = 6
	m["g"] = 7
	m["h"] = 8
}

//func main() {
//	var m map[int]int
//	fn1(m)
//	fmt.Println("m == nil in main?:", m == nil) // m == nil in main?: true
//}

func main() {
	m := map[string]int{"a": 1, "b": 2}
	fn2(m)
	fmt.Println(m)      // map[a:1 b:2 c:3 d:4 e:5 f:6 g:7 h:8]
	fmt.Println(len(m)) // 8
}
