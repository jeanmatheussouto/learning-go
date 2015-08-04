package main

import (
	"errors"
	"fmt"
)

func main() {
	pile := Pile{}
	fmt.Println("--------")
	fmt.Println("New pile build with size ", pile.Size())
	fmt.Println("Pile empty: ", pile.Empty())
	fmt.Println("--------")

	fmt.Println("Stack elements")

	pile.Stack("Jean")
	pile.Stack("Camila")
	pile.Stack(1)
	pile.Stack(true)
	pile.Stack(3.14)

	fmt.Println("Pile size: ", pile.Size())
	fmt.Println("Pile empty: ", pile.Empty())
	fmt.Println("--------")

	for !pile.Empty() {
		value, _ := pile.Unstack()
		fmt.Println("Unstack: ", value)
		fmt.Println("Pile size: ", pile.Size())
		fmt.Println("Pile empty: ", pile.Empty())
		fmt.Println("--------")
	}

	_, err := pile.Unstack()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

type Pile struct {
	values []interface{}
}

func (pile Pile) Size() int {
	return len(pile.values)
}

func (pile Pile) Empty() bool {
	return pile.Size() == 0
}

func (pile *Pile) Stack(value interface{}) {
	pile.values = append(pile.values, value)
}

func (pile *Pile) Unstack() (interface{}, error) {
	if pile.Empty() {
		return nil, errors.New("Empty pile")
	}
	value := pile.values[pile.Size()-1]
	pile.values = pile.values[:pile.Size()-1]
	return value, nil
}
