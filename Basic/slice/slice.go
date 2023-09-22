package main

import "fmt"

func getSlice1(sl []int) {
	sl = append(sl, 11)
}

func getSlice2(sl *[]int) {
	*sl = append(*sl, 11)
}

func main() {
	sl3 := make([]int, 0, 10) //len: 0; cap: 10

	sl1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //len: 10; cap: 10
	sl2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 1; i < 10; i++ {
		sl3[i-1] = i
	}

	getSlice1(sl1)
	getSlice1(sl3)
	getSlice2(&sl2)

	fmt.Printf("sl1: %d; cap: %d; len: %d\n", sl1, cap(sl1), len(sl1)) // sl1: [1 2 3 4 5 6 7 8 9 10]; cap: 10; len: 10
	fmt.Printf("sl2: %d; cap: %d; len: %d\n", sl2, cap(sl2), len(sl2)) // sl2: [1 2 3 4 5 6 7 8 9 10 11]; cap: 20; len: 11
	fmt.Printf("sl3: %d; cap: %d; len: %d\n", sl3, cap(sl3), len(sl3)) // sl3: [1 2 3 4 5 6 7 8 9]; cap: 10; len: 9
}

//
//func main() {
//	slice := []string{"a", "a"}
//
//	func(slice []string) {
//		slice = append(slice, "a")
//		slice[0] = "b"
//		slice[1] = "b"
//		fmt.Print(slice)
//	}(slice)
//	fmt.Print(slice) // [b b a][a a]
//}
