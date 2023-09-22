package main

import "fmt"

func main() {
	sl1 := make([]int, 10, 10)
	sl2 := make([]int, 10, 10)
	sl3 := make([]int, 10, 10)

	sl1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl3 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	getSlice1(sl1)
	getSlice1(sl3)
	getSlice2(&sl2)

	fmt.Printf("sl1: %d; cap: %d; len: %d\n", sl1, cap(sl1), len(sl1))
	fmt.Printf("sl2: %d; cap: %d; len: %d\n", sl2, cap(sl2), len(sl2))
	fmt.Printf("sl3: %d; cap: %d; len: %d\n", sl3, cap(sl3), len(sl3))
}

func getSlice1(sl []int) {
	sl = append(sl, 11)
}

func getSlice2(sl *[]int) {
	*sl = append(*sl, 11)
}
