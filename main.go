package main

import (
	"calculator_module/app"
	"fmt"
)

func main() {
	s, _ := app.Reader()
	arrS, b, _ := app.UserInputFilter(s)
	fmt.Println(app.Process(arrS, b))
}
