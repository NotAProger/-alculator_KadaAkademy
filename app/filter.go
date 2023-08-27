package app

import (
	"fmt"
	"strings"
)

func UserInputFilter(s string) ([]string, []string, bool, error) {
	sliceArgs := strings.Split(s, " ")
	var isRome bool
	var err error

	if len(sliceArgs) > 3 {
		return nil, nil, false, fmt.Errorf("quantity of arguments and mathematical symbols are %d, need 3", len(s))
	} else if len(sliceArgs) < 3 {
		return nil, nil, false, fmt.Errorf("string integrity violated (check spaces and presence of 2 arguments and sign)")
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
	} else if r >= '1' && r <= '9' || r == '-' {
		isRome = false
	} else {
		return false, fmt.Errorf("1 user number %s is incorrect (check for numbers)", s[0])
	}

	for num := 0; num < len(s); num++ {
		for _, r := range s[num] {
			if isRome {
				if r == 'I' || r == 'V' || r == 'X' {
				} else {
					return false, fmt.Errorf("%d user number %s is not rome format", num+1, s[num])
				}
			} else {
				if r >= '0' && r <= '9' || r == '-' {
				} else {
					return false, fmt.Errorf("%d user number %s is not arab format", num+1, s[num])
				}
			}
		}
	}
	return isRome, nil
}

func checkForCorrectMathematicalSymbol(s []string) error {
	for _, r := range s {
		if r == "+" || r == "-" || r == "*" || r == "/" {
		} else {
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
