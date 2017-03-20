package main

import (
	"fmt"
	"math"
)

func values() {
	defer fmt.Println()
	fmt.Println("Values in go!--------------------------------------------------------------------")

	fmt.Println("go" + "lang")
	fmt.Println("21+22=", 21+22)
	fmt.Println("7.0/3.0=", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func variables() {
	defer fmt.Println()
	fmt.Println("Variables in go!--------------------------------------------------------------------")

	var declared string

	declared = "initialized"
	fmt.Println(declared)

	// simple declaration with type and initialization
	var simpleDeclaration string = "string"
	fmt.Println(simpleDeclaration)

	// declaring and initializing multiple variables in one statement
	var b, c int = 1, 2
	fmt.Println(b, c)

	// simple declaration without type and initialization
	var d = true
	fmt.Println(d)

	// simple declaration with type but no initializtion which results in zero value. 0 for int, false for bool, "" for string
	var e int
	fmt.Println(e)

	// declare and initialize without using the var keyword. Type is infered and cannot be specified
	declAndInitialize := "test"
	fmt.Println(declAndInitialize)

	// multiple variables initialization in a block
	var (
		v1 = 1
		v2 = 2
	)

	fmt.Println(v1)
	fmt.Println(v2)

}

const (
	globalConstant = "a"
)

func constants() {
	defer fmt.Println()
	fmt.Println("Constants in go!--------------------------------------------------------------------")

	fmt.Println(globalConstant)

	const denominator = 500000000

	const result = 3e20 / denominator

	fmt.Println(result)
	fmt.Println(int64(result))
	fmt.Println(math.Sin(result))
}