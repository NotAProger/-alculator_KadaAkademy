package app

import (
	"bufio"
	"fmt"
	"os"
)

func Reader() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter mathematical example / Введите математический пример")
	example, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return example, nil
}
