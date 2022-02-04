package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Please enter customer age")
	}

	age := os.Args[1]

	if age, err := strconv.Atoi(age); err != nil || age < 1 {
		fmt.Printf("%v is not a vlaid age. Please enter customer age", age)
		return
	} else if age >= 17 {
		fmt.Println("R-Rated")
	} else if age < 17 && age >= 13 {
		fmt.Println("PG-13")
	} else if age < 13 {
		fmt.Println("PG-Rated")
	}

}
