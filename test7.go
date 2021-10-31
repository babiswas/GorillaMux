package main

import "fmt"

func generator() func() int {
	var i = 0
	return func() int {
		i++
		return i
	}
}

func main() {
	originalgenerator := generator()
	for i := 0; i < 5; i++ {
		fmt.Println(originalgenerator())
	}
}
