package main

import (
	"fmt"
)

func main() {
	//printElements()
	addElements()
}

func printElements() {
	//countries := make(map[string]string)
	countries := make(map[string]string, 5)
	fmt.Println(countries)

	countries["Peru"] = "Lima"
	countries["Argentina"] = "Buenos Aires"

	fmt.Println(countries)
	fmt.Println(countries["Peru"])

}

func addElements() {
	champion := map[string]int{
		"Barcelona":   2,
		"Real madrid": 20,
		"Chivas":      10}

	fmt.Println(champion)

	// add elements to map
	champion["River plate"] = 25

	// change value
	champion["Chivas"] = 55

	//delete element
	delete(champion, "Real madrid")

	fmt.Println(champion)

	for group, point := range champion {
		fmt.Printf("El equipo %s, tiene un puntaje de : %d \n", group, point)
	}
	point, exist := champion["Mineiro"]
	fmt.Printf("El puntaje capturado es %d, el equipo existe : %t \n", point, exist)

}
