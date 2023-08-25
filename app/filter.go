package app

import (
	"fmt"
	"strings"
)

func UserInputFilter(s string) ([]string, bool, error) {
	sliceArgs := strings.Split(s, " ")
	var firstIsRome bool
	var err error

	if len(sliceArgs) > 3 {
		err = fmt.Errorf("quantity of arguments and mathematical symbols more then 2")
	} else if len(sliceArgs) < 3 {
		err = fmt.Errorf("string integrity violated (check spaces and presence of 2 arguments and sign)")
	}
	if err != nil {
		return nil, false, err
	}

	if firstIsRome, err = checkForCorrectNumbers(sliceArgs[0], sliceArgs[2]); err != nil {
		return nil, false, err
	}

	if err = checkForCorrectMathematicalSymbol(sliceArgs[1]); err != nil {
		return nil, false, err
	}

	return sliceArgs, firstIsRome, nil
}

func checkForCorrectNumbers(s1, s2 string) (bool, error) {
	var firstIsRome bool

	r := s1[0]
	if r >= '1' && r <= '9' {
		firstIsRome = false
	} else if r == 'I' || r == 'V' || r == 'X' {
		firstIsRome = true
	} else {
		return false, fmt.Errorf("first user number is incorrect (check for numbers)")
	}

	r = s2[0]
	if r >= '1' && r <= '9' {
		if !firstIsRome {
			return false, nil
		} else {
			return false, fmt.Errorf("second number is not an arabic")
		}
	} else if r == 'I' || r == 'V' || r == 'X' {
		if firstIsRome {
			return true, nil
		} else {
			return false, fmt.Errorf("second number is not an arabic")
		}
	} else {
		return false, fmt.Errorf("second user number is incorrect (check for numbers)")
	}
}

func checkForCorrectMathematicalSymbol(s string) error {
	switch s {
	case "+":
		return nil
	case "-":
		return nil
	case "*":
		return nil
	case "/":
		return nil
	default:
		return fmt.Errorf("mathematical symbol is incorrect or is not supported")
	}
}
