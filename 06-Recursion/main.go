package main

import "fmt"

func main() {
	expo(2)
}

func expo(number int) {
	if number > 100000 {
		return
	}
	fmt.Println(number)
	expo(number * 2)
}
