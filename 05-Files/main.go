package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//readFile()
	//readFile2()
	//saveFile()
	saveFile2()
}

func readFile() {
	file, err := ioutil.ReadFile("./newFile.txt")
	if err != nil {
		fmt.Printf("There was an error %d", err)
	} else {
		fmt.Println(string(file))
	}
}

func readFile2() {
	file, err := os.Open("./newFile.txt")
	if err != nil {
		fmt.Printf("There was an error %d", err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			register := scanner.Text()
			fmt.Printf("Linea > " + register + "\n")
		}
	}
	file.Close()
}

func saveFile() {
	file, err := os.Create("./newFile.txt")
	if err != nil {
		fmt.Printf("There was an error %d", err)
		return
	}
	fmt.Fprintln(file, "this is the new line")
	file.Close()
}

func saveFile2() {
	fileName := "./newFile.txt"

	if Append(fileName, "\n This is the second line") == false {
		fmt.Println("There was an error")
	}
}

func Append(file string, text string) bool {
	arch, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("There was an error")
		return false
	}
	_, err = arch.WriteString(text)
	if err != nil {
		fmt.Println("There was an error")
		return false
	}
	return true
}
