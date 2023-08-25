package main

import (
	"calculator_module/app"
	"fmt"
)

func main() {
	s, _ := app.Reader()
	fmt.Println(app.UserInputFilter(s))
}
