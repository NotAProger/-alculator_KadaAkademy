package main

import (
	"calculator_module/app"
	"fmt"
	"os"
)

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Run() error {
	var numStr []string
	var sign []string
	var num []int
	var isRome bool
	var example string
	var err error

	if example, err = app.Reader(); err != nil {
		return err
	}
	if numStr, sign, isRome, err = app.UserInputFilter(example); err != nil {
		return err
	}
	if num, err = app.Convertation(numStr, isRome); err != nil {
		return err
	}
	if example, err = app.Calculation(num, sign, isRome); err != nil {
		return err
	}
	fmt.Println("Answer: ", example)
	return nil
}
