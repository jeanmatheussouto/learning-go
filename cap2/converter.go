package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("*** Converter ***")
	argsLen := len(os.Args)

	if argsLen < 3 {
		fmt.Println("Use: converter <vals>,<unit>")
		os.Exit(1)
	}

	originUnit := os.Args[argsLen-1]
	originVals := os.Args[1 : argsLen-1]

	var destinyUnit string

	fmt.Println(originUnit, originVals, destinyUnit)

	if originUnit == "celsius" {
		destinyUnit = "fahrenheit"
	} else if originUnit == "kilometers" {
		destinyUnit = "miles"
	} else {
		fmt.Printf("%s isn't an unit known", originUnit)
		os.Exit(1)
	}

	for index, value := range originVals {
		valOrigin, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Printf("The value %s in index %d isn't a valid number", value, index)
			os.Exit(1)
		}

		var valDestiny float64

		if originUnit == "celsius" {
			valDestiny = valOrigin*1.8 + 32
		} else {
			valDestiny = valOrigin * 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n", valOrigin, originUnit, valDestiny, destinyUnit)
	}
}
