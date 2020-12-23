package main

import (
	"fmt"
	"math"
	"sort"
)

type variable struct {
	x              int32
	y              int32
	className      string
	targetDistance float64
}

func main() {

	variables := []variable{
		{
			x:         2,
			y:         4,
			className: "Sınıf1",
		},
		{
			x:         3,
			y:         7,
			className: "Sınıf1",
		},
		{
			x:         5,
			y:         4,
			className: "Sınıf1",
		},
		{
			x:         3,
			y:         8,
			className: "Sınıf2",
		},
		{
			x:         5,
			y:         9,
			className: "Sınıf2",
		},
		{
			x:         7,
			y:         10,
			className: "Sınıf2",
		},
		{
			x:         6,
			y:         8,
			className: "Sınıf2",
		},
	}

	targetVariable := variable{
		x: 4,
		y: 7,
	}

	fmt.Printf(algorithmKNN(variables, 3, targetVariable))

}

func algorithmKNN(variables []variable, k int, targetVariable variable) string {

	for index, variable := range variables {

		variables[index].targetDistance = euclid(variable.x, variable.y, targetVariable.x, targetVariable.y)
	}

	sort.Slice(variables, func(i, j int) bool {
		return variables[i].targetDistance < variables[j].targetDistance
	})

	classMap := map[string]int{}

	for i := 0; i < k; i++ {
		val, ok := classMap[variables[i].className]
		if ok == true {
			classMap[variables[i].className] = val + 1
		}
		if ok == false {
			classMap[variables[i].className] = 1
		}
	}

	maxClass := ""
	maxNumber := 0
	for key, value := range classMap {

		if value > maxNumber {
			maxNumber = value
			maxClass = key
			continue
		}
	}

	return maxClass

}

func euclid(x1 int32, y1 int32, x2 int32, y2 int32) float64 {

	result := math.Sqrt(float64((x1-x2)*(x1-x2)) + float64((y1-y2)*(y1-y2)))
	return result

}
