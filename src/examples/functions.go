package main

import "fmt"

func functions() {
	defer fmt.Println()
	fmt.Println("Functions in go!--------------------------------------------------------------------")

	acceptMultipleInputsOfTheSameType("Hello,", " World!")
	returnMultipleValuesViaNamedParameters("Hello, world")
	msg, length := returnMultipleValues("Hello, world")
	fmt.Printf("%s %d\n", msg, length)
	varagsAsParameters("%s\n", "Hello, World!", "Greetings Earthlings!")
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