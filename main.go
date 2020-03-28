package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

func dev() {
	fmt.Println("add branch dev")
}

// new code from master
func one() {
	fmt.Println("new day from master")
	fmt.Println("update by dev")
	fmt.Println("update by dev again")
}

func two() {
	fmt.Println("add by master")
}

func three() {
	fmt.Println("add by dev")
}

// add four
func four() {
	fmt.Println("add by dev")
}
