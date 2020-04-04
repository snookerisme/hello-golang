package main

import "hello/vars"
import "fmt"

func main() {
	name()
	println("Score 71 => Grade :", grade(71))
	array()
}

func name() {
	println("Hello ", vars.Name)
}

func init() {
	fmt.Println("initial ...")
}

func a(i float64) string {
	return "A"
}

func grade(i float64) string {
	if i > 90 && i <= 100 {
		return "A"
	} else if i > 80 && i <= 90 {
		return "B"
	} else if i > 70 && i <= 80 {
		return "C"
	} else if i > 60 && i <= 70 {
		return "D"
	}
	return "F"
}

func array() {
	var slice = [...]int{1, 2, 3}

	for _, v := range slice {
		fmt.Println(v)
	}
}
