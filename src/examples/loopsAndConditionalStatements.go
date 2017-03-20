package main

import (
	"fmt"
	"time"
)

func forLoops() {
	defer fmt.Println()
	fmt.Println("Loops in go!--------------------------------------------------------------------")

	i := 1
	//while loop
	for i <= 3 {
		fmt.Println("i", i)
		i = i + 1
	}

	// infinite while loop
	for {
		fmt.Println("Infinite")
		break
	}

	//traditional for loop
	for j := 5; j <= 9; j++ {
		fmt.Println("j", j)
	}

	//traditional for loop with continue
	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("n", n)
	}
}

func ifElse() {
	defer fmt.Println()
	fmt.Println("If Else in go!--------------------------------------------------------------------")

	//traditional if else
	if 7%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	// initializing a variable and using it in if check. The scope of the variable is only within the if and subsequent else statements
	if num := 9; num < 9 {
		fmt.Println("less than 9")
	} else if num > 10 {
		fmt.Println("greater than 9")
	} else {
		fmt.Println("is 9")
	}
}

func switchExamples() {
	defer fmt.Println()
	fmt.Println("Switch in go!--------------------------------------------------------------------")

	// traditional switch case
	i := 2
	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	}

	// multiple case check. There is no fallthrough by default, so break not needed.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	character := 'c'
	// switch with fallthrough
	switch character {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("Vowel")
	case 'c':
		fmt.Println("The character is c")
		fallthrough
	default:
		fmt.Println("Consonant")
	}

	// switch without expression. Can be used as an alternative for if else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Before noon")
	default:
		fmt.Println("After noon")
	}

	// switch on the data type
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Bool")
		case int:
			fmt.Println("Int")
		default:
			fmt.Printf("unknown %T", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hello")
}
