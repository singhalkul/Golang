package main

import "fmt"

func strings() {
	defer fmt.Println()
	defer fmt.Println()
	fmt.Println("Strings in go!--------------------------------------------------------------------")

	str := "the quick brown fox jumped over the lazy dog"

	fmt.Println("Entire string: ", str)
	fmt.Println("Substring from 0 to 9: ", str[0:9])
	fmt.Println("Substring from 0 to 9 by specifying only end: ", str[:9])
	fmt.Println("Substring from 14 till the end by specifying only the start: ", str[14:])

	unescaped := `the "quick" brown fox jumped over the lazy dog\n\n`
	fmt.Println("Handling double quotes in string by using backticks for creating strings: ", unescaped)

	fmt.Println("Length of the string is: ", len(str))

	// Iterating over individual characters and the indexes of the string
	for i, r := range str {
		fmt.Printf("%d %c\t", i, r)
	}
}