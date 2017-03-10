package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	helloWorld()
	values()
	variables()
	constants()
	forLoops()
	ifElse()
	switchExamples()
	strings()
	functions()
	arrays()
	slices()
	maps()
}

func helloWorld() {
	fmt.Println("Hello, World!")
}

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

func acceptMultipleInputsOfTheSameType(m1, m2 string) {
	fmt.Printf("%s", m1)
	fmt.Printf("%s\n", m2)
}

func returnMultipleValues(m1 string) (string, int) {
	defer fmt.Println()
	fmt.Printf("%s", m1)
	return m1, len(m1)
}

func returnMultipleValuesViaNamedParameters(m1 string) (r1 string, length int) {
	defer fmt.Println()
	fmt.Printf("%s", m1)
	r1 = m1 + "\n"
	length = len(m1)
	return
}

func varagsAsParameters(format string, messages ...string) {
	// underscore means ignore the value
	for _, msg := range messages {
		fmt.Printf(format, msg)
	}
}

func functions() {
	defer fmt.Println()
	fmt.Println("Functions in go!--------------------------------------------------------------------")

	acceptMultipleInputsOfTheSameType("Hello,", " World!")
	returnMultipleValuesViaNamedParameters("Hello, world")
	msg, length := returnMultipleValues("Hello, world")
	fmt.Printf("%s %d\n", msg, length)
	varagsAsParameters("%s\n", "Hello, World!", "Greetings Earthlings!")
}

func arrayPrinter(words [4]string) {
	defer fmt.Println()
	for _, word := range words {
		fmt.Printf("%s ", word)
	}
	words[2] = "black"
}

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

func slicePrinter(words []string) {
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

func maps() {
	defer fmt.Println()
	fmt.Println("Maps in go!--------------------------------------------------------------------")

	dayMonths := make(map[string]int)
	dayMonths["Jan"] = 31
	dayMonths["Feb"] = 28
	dayMonths["Mar"] = 31
	dayMonths["Apr"] = 30
	dayMonths["May"] = 31
	dayMonths["Jun"] = 30
	dayMonths["Jul"] = 31
	dayMonths["Aug"] = 31
	dayMonths["Sep"] = 30
	dayMonths["Oct"] = 31
	dayMonths["Nov"] = 30
	dayMonths["Dec"] = 31

	fmt.Printf("Days in February: %d\n", dayMonths["Feb"])
	fmt.Printf("Days in ZZZ: %d\n", dayMonths["ZZZ"]) //returns the zero value in Go instead of Nil. So this returns 0

	// if we can't live with the zero value, we can use the below way to check if the boolean value is true indicating that
	// value was found.
	days, ok := dayMonths["ZZZ"]
	if !ok {
		fmt.Println("Can't find days for ZZZ")
	} else {
		fmt.Println("Days for month are: ", days)
	}

	for month, days := range dayMonths {
		fmt.Printf("%s %d\t", month, days)
	}
	fmt.Println()

	delete(dayMonths, "Feb")
	fmt.Printf("Days in February: %d\n", dayMonths["Feb"])
	delete(dayMonths, "Feb")
	fmt.Printf("Days in February: %d\n", dayMonths["Feb"])

	// Initializing the map inline without using make
	newMap := map[string]int{"a": 97, "b": 98}
	fmt.Println(newMap)

}
