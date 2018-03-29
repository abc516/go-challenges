package main

import (
	"fmt"
	"strconv"
)

func main() {
	i:= 1
	for i=1; i <100 ;i++  {
		if i % 3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i % 5 == 0 {
			fmt.Println("Buzz")
			continue
		} else {
		    fmt.Println(strconv.Itoa(i))
		}

	}
}