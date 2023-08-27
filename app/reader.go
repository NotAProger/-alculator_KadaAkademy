package app

import (
	"bufio"
	"fmt"
	"os"
)

func Reader() (string, error) {
	fmt.Print("Enter example: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	return text, nil
}
