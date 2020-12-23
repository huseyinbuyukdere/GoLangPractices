package main

import "fmt"

//MyMap is My MAp
type MyMap map[string]string

func main() {

	//var colors map[string]string
	// colors := make(MyMap)
	// colors[1] = "#ffffff"
	// delete(colors, 10)
	// fmt.Println(colors)

	colors := MyMap{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	colors.printMap("-")

}

func (valueMap MyMap) printMap(seperator string) {

	for color, hex := range valueMap {
		fmt.Printf("%v%v%v\n", color, seperator, hex)
	}

}
