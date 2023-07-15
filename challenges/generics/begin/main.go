// challenges/generics/begin/main.go
package main

import (
	"errors"
	"fmt"
)

// Part 1: print function refactoring

// non-generic print functions

func printString(s string) { fmt.Println(s) }

func printInt(i int) { fmt.Println(i) }

func printBool(b bool) { fmt.Println(b) }

func printAny[T string | int | bool](a T) { fmt.Println(a) }

// Part 2 sum function refactoring

// sum sums a slice of any type
func sum(numbers []interface{}) interface{} {
	var result float64
	for _, n := range numbers {
		switch n.(type) {
		case int:
			result += float64(n.(int))
		case float32, float64:
			result += n.(float64)
		default:
			continue
		}
	}
	return result
}

type numeric interface {
	~int | ~float64
}

func sumAny[T numeric](numbers ...T) (T, error) {
	var result T
	if len(numbers) == 0 {
		return result, errors.New("no numbers provided to sum")
	}
	for _, number := range numbers {
		result += number
	}
	return result, nil
}

func main() {
	// call non-generic print functions
	//printString("Hello")
	//printInt(42)
	//printBool(true)

	// call generic printAny function for each value above
	printAny("Hello")
	printAny(42)
	printAny(true)

	// call sum function
	//fmt.Println("result", sum([]interface{}{1, 2, 3}))

	// call generics sumAny function
	result, err := sumAny(1, 2, 3)
	if err != nil {
		panic("💩")
	}
	fmt.Println("result", result)

	fmt.Println("result", func() (result int) {
		result, _ = sumAny(1, 2, 3)
		return
	}())
}
