package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	args := os.Args

	if len(args) <= 1 {
		return
	}

	fileName := os.Args[1]

	file, error := os.Open(fileName)

	if error != nil {
		fmt.Println(error)
		return
	}

	content := make([]byte, 101010)

	_, error = file.Read(content)

	if error != nil {
		fmt.Println(error)
	}

	defer file.Close()

	fmt.Println(string(content))

	newFile, _ := os.Create("other.txt")

	defer newFile.Close()

	_, error = io.Copy(newFile, file)
	if error != nil {
		fmt.Println(error)
	}

	newFile.Close()
}
