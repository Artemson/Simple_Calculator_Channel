package main

import (
	"fmt"
	"strconv"
	"time"
	"unicode"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()
	calc(ch, "1+1", "2+2", "3*3")
}

func calc(ch chan int, args ...string) {
	for i := range args {
		var firstString, secondString string = "", ""
		var firstInt, secondInt int = 0, 0
		var operation string = ""
		var res int
		for _, elem := range args[i] {
			if unicode.IsDigit(elem) && operation == "" {
				firstString += string(elem)
			} else if unicode.IsDigit(elem) {
				secondString += string(elem)
			} else if string(elem) != " " {
				operation = string(elem)
			} else {
				continue
			}
		}
		firstInt, _ = strconv.Atoi(firstString)
		secondInt, _ = strconv.Atoi(secondString)
		switch operation {
		case "+":
			res = firstInt + secondInt
		case "-":
			res = firstInt - secondInt
		case "*":
			res = firstInt * secondInt
		case "/":
			res = firstInt / secondInt
		}
		ch <- res
		time.Sleep(1 * time.Second)
	}
	close(ch)
}
