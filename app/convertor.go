package app

import (
	"fmt"
	"strconv"
)

func Convertation(arrNums []string, isRome bool) ([]int, error) {
	arrNumsInt := make([]int, len(arrNums))
	var err error

	_ = arrNumsInt
	if isRome {
		if arrNumsInt, err = romeConvertor(arrNums); err != nil {
			return nil, err
		}
	} else {
		if arrNumsInt, err = arabConvertor(arrNums); err != nil {
			return nil, err
		}
	}

	return arrNumsInt, err
}

func arabConvertor(s []string) ([]int, error) {
	arrNum := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		num, err := strconv.Atoi(s[i])
		if err != nil {
			return nil, err
		}
		if num < -10 || num > 10 {
			return nil, fmt.Errorf("%s number is less -10 or more then 10", s[i])
		}
		arrNum[i] = num
	}

	return arrNum, nil
}

func romeConvertor(s []string) ([]int, error) {
	arrNum := make([]int, len(s))
	var postNum, num int
	roman := make(map[byte]int)
	roman['I'] = 1
	roman['V'] = 5
	roman['X'] = 10

	for charNum := 0; charNum < len(s); charNum++ {
		for i := len(s[charNum]) - 1; i >= 0; i-- {
			num = roman[s[charNum][i]]
			if postNum <= num {
				arrNum[charNum] += num
			} else {
				arrNum[charNum] -= num
			}
			postNum = num
		}
		postNum, num = 0, 0
	}

	for i := 0; i < len(s); i++ {
		if arrNum[i] > 10 || arrNum[i] < 1 {
			return nil, fmt.Errorf("%s rome number is more than 10 or below 1", s[i])
		}
	}

	return arrNum, nil
}
