package main

import (
	"os"
	"strconv"
	"fmt"
	"../utils"
)

func main() {
	if argsCount := len(os.Args); argsCount > 1 {
		value, err := strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			fmt.Println("Couldn't parse provided value")
		} else {
			fmt.Println(utils.NeotainSqrt(value))
		}
	} else {
		fmt.Println("please enter a value")
	}
}
