package main

import "fmt"

var table [10]int

func main() {
	initTable()
	initTable2()

}

func initTable() {
	table[0] = 1
	table[5] = 15
	fmt.Println(table)

}
func initTable2() {
	table2 := [10]int{5, 6, 8, 69, 7, 14, 2, 5, 1, 2}
	fmt.Println(table2)
}
