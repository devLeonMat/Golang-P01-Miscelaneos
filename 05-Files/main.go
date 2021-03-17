package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	readFile()
}

func readFile() {
	file, err := ioutil.ReadFile("./archivo.txt")
	if err != nil {
		fmt.Printf("There was an error %d", err)
	} else {
		fmt.Println(string(file))
	}

}
