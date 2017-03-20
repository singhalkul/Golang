package main

import "fmt"

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