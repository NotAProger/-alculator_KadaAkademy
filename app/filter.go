package app

import (
	"fmt"
	"strings"
)

// qwantity of arguments in example f.e. "3 + 5" is 3 arg
const maxArgs int = 3

func UserInputFilter(s string) ([]string, []string, bool, error) {
	sliceArgs := strings.Split(s, " ")
	var isRome bool
	var err error

	if err = isEnafArgs(len(sliceArgs)); err != nil {
		return nil, nil, false, err
	}

	arrNums, arrSign := sortingStrs(sliceArgs)

	if isRome, err = checkForCorrectNumbers(arrNums); err != nil {
		return nil, nil, false, err
	}

	if err = checkForCorrectMathematicalSymbol(arrSign); err != nil {
		return nil, nil, false, err
	}

	return arrNums, arrSign, isRome, nil
}

func checkForCorrectNumbers(s []string) (bool, error) {
	var isRome bool

	r := s[0][0]
	if r == 'I' || r == 'V' || r == 'X' {
		isRome = true
	}

	for num := 0; num < len(s); num++ {
		if isRome {
			for _, r := range s[num] {
				if r != 'I' && r != 'V' && r != 'X' {
					return false, fmt.Errorf("%d user number %s is not rome format", num+1, s[num])
				}
			}
		} else {
			for _, r := range s[num] {
				if r < '0' || r > '9' && r != '-' {
					return false, fmt.Errorf("%d user number %s is not arab format", num+1, s[num])
				}
			}
		}
	}
	return isRome, nil
}

func checkForCorrectMathematicalSymbol(s []string) error {
	for _, r := range s {
		if r != "+" && r != "-" && r != "*" && r != "/" {
			return fmt.Errorf("%s mathematical symbol is incorrect or is not supported", r)
		}
	}
	return nil
}

func sortingStrs(s []string) ([]string, []string) {
	arrNums := []string{}
	arrSign := []string{}

	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			arrNums = append(arrNums, s[i])
		} else {
			arrSign = append(arrSign, s[i])
		}
	}

	return arrNums, arrSign
}

func isEnafArgs(len int) error {
	if len > maxArgs {
		return fmt.Errorf("quantity of arguments and mathematical symbols are %d, need 3", maxArgs)
	}
	if len < maxArgs {
		return fmt.Errorf("string integrity violated (check spaces and presence of arguments and sign)")
	}
	return nil
}
