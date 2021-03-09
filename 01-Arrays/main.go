package main

import "fmt"

var table [10]int
var matriz [5][7]int

func main() {
	//initTable()
	//initTable2()
	//initMatriz()
	//initSlices()
	//initSlices2()
	//initSlices3()
	initSlices4()

}

// slices
func initSlices() {
	matSlice := []int{2, 5, 4}
	fmt.Println(matSlice)
}
func initSlices2() {
	elements := [5]int{1, 2, 3, 4, 5}
	partial := elements[3:]
	fmt.Println(partial)
}
func initSlices3() {
	elements := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elements), cap(elements))
}
func initSlices4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))
}

// matriz and arrays
func initMatriz() {
	matriz[3][5] = 1

	fmt.Println(matriz)
}

func initTable() {
	table[0] = 1
	table[5] = 15
	fmt.Println(table)

}

func initTable2() {
	table2 := [10]int{5, 6, 8, 69, 7, 14, 2, 5, 1, 2}
	for i := 0; i < len(table2); i++ {
		fmt.Println(table2[i])
	}
}
