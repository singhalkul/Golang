package main

import (
	"fmt"
	"errors"
	"os"
)

func errorsExample() {
	defer fmt.Println()
	fmt.Println("Errors in go!--------------------------------------------------------------------")
	if err := testPrinter("Hello World!"); err != nil {
		os.Exit(1)
	}

	if err := errorsUsingFmtPackage(""); err != nil {
		fmt.Println(err)
	}

	if err := errorUsingCustomErrors(""); err != nil {
		if err == errorEmptyString {
			fmt.Printf("Custom errorEmptyString error occurred")
		} else {
			fmt.Println(err)
		}
	}

	panicExample("")
}

func testPrinter(msg string) error {
	_, err := fmt.Printf("%s\n", msg)
	return err
}

func errorsUsingFmtPackage(msg string) error {
	if msg == "" {
		return fmt.Errorf("Won't print an empty string...")
	}
	return testPrinter(msg)
}

var (
	errorEmptyString = errors.New("Won't print an empty string...")
)

func errorUsingCustomErrors(msg string) error {
	if msg == "" {
		return errorEmptyString
	}
	return testPrinter(msg)
}

func panicExample(msg string) error {
	if msg == "" {
		panic(errorEmptyString)
	}
	return testPrinter(msg)
}