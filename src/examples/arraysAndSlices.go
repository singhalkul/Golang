package main

import "fmt"

func arrays() {
	defer fmt.Println()
	fmt.Println("Arrays in go!--------------------------------------------------------------------")

	wordsArrayUnbound := [...]string{"the", "quick", "brown", "fox"}
	fmt.Println(wordsArrayUnbound[2])

	wordsArrayFixed := [4]string{"the", "quick", "brown", "fox"}
	fmt.Println(wordsArrayFixed[2])

	// When arrays are passed to functions, a copy is created as other types and this could cause an issue for large arrays.
	arrayPrinter(wordsArrayUnbound)
	arrayPrinter(wordsArrayUnbound)
}

func arrayPrinter(words [4]string) {
	defer fmt.Println()
	for _, word := range words {
		fmt.Printf("%s ", word)
	}
	words[2] = "black"
}

func slices() {
	defer fmt.Println()
	fmt.Println("Slices in go!--------------------------------------------------------------------")

	words := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}

	fmt.Println("The length of the slice is: ", len(words))

	// Slices are passed by reference and that is why any modification in the function modifies the original data as well.
	// Since slice are passed by reference, using them will speed up the program.
	slicePrinter(words)
	slicePrinter(words)

	// Slices are subset of the original array, but it is not a new copy.
	fmt.Println(words[0:2])
	fmt.Println(words[:3])
	fmt.Println(words[5:len(words)])
	fmt.Println(words[5:])

	newWords := make([]string, 4, 4) // creates a slice of 100 elements with 4 having 0 value i.e. empty string
	fmt.Printf("%d %d\n", len(newWords), cap(newWords))
	fmt.Println(newWords)
	newWords[0] = "The"
	newWords[1] = "Quick"
	newWords[2] = "Brown"
	newWords[3] = "Fox"
	// If we want to add new entries in the slice, we can append them using the append function, can't do this directly
	// by saying newWords[4] = "Jumps"
	newWords = append(newWords, "Jumps")

	fmt.Printf("%d %d\n", len(newWords), cap(newWords))
	fmt.Println(newWords)

	copiedWords := make([]string, len(words))
	copy(copiedWords, words) // If length of new slice is less than the source, elements upto the length will be copied rest will not be copied
	copiedWords[2] = "pink"
	slicePrinter(words)
	slicePrinter(copiedWords)
}

func slicePrinter(words []string) {
	defer fmt.Println()
	for _, word := range words {
		fmt.Printf("%s ", word)
	}
	words[2] = "black"
}