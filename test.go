package main

import "fmt"

func double(num *int) {
	*num *= 2
}

func main() {
	num := 6
	double(&num)
	fmt.Println(num)
}
