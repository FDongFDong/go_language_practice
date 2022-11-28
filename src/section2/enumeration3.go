package main

import "fmt"

func main() {
	const (
		_ = iota
		A
		_
		C
	)
	const (
		_ = iota + 0.75*2
		DEFAULT
		SILVER
		_
		PLATINUM
	)
	fmt.Println("Default : ", DEFAULT, "Silver : ", SILVER, "Platinum : ", PLATINUM)
}
