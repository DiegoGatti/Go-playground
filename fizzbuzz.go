package main

// import (
// 	"fmt"
// )

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				println("Fizz-Buzz")
			} else {
				println("Fizz")
			}
		} else if i%5 == 0 {
			println("Buzz")
		} else {
			println(i)
		}
	}
}
