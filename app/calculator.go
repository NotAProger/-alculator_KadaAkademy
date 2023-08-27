package app

import (
	"fmt"
	"strconv"
)

func Calculation(nums []int, sign []string, isRome bool) (string, error) {
	ansver := ""
	ans, err := calculationNums(nums, sign)

	if isRome {
		ansver = roman(ans)
	} else {
		ansver = strconv.Itoa(ans)
	}

	return ansver, err
}

func calculationNums(num []int, sign []string) (int, error) {
	res := 0
	switch sign[0] {
	case "+":
		res = num[0] + num[1]
	case "-":
		res = num[0] - num[1]
	case "*":
		res = num[0] * num[1]
	case "/":
		if num[1] != 0 {
			res = num[0] / num[1]
		} else {
			return 0, fmt.Errorf("cant divide by 0")
		}
	}
	return res, nil
}

func roman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}
	return roman
}
