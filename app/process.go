package app

import (
	"fmt"
	"strconv"
)

func Process(s []string, b bool) (int, bool, error) {
	arrNums := []string{}
	arrNumsInt := []int{}
	arrSign := []string{}
	var err error
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			arrNums = append(arrNums, s[i])
		} else {
			arrSign = append(arrSign, s[i])
		}
	}
	fmt.Println(arrNumsInt)

	if b {
		if arrNumsInt, err = romeOperator(arrNums); err != nil {
			return 0, false, err
		}
	} else {
		if arrNumsInt, err = arabOperator(arrNums); err != nil {
			return 0, false, err
		}
	}
	ans, err := calculation(arrNumsInt, arrSign)
	return ans, b, err
}

func arabOperator(s []string) ([]int, error) {
	arrNum := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		num, err := strconv.Atoi(s[i])
		if err != nil {
			return nil, err
		}
		if num <= -10 && num >= 10 {
			return nil, fmt.Errorf(strconv.Itoa(i), " number is less -10 or more then 10")
		}
		arrNum[i] = num
	}
	return arrNum, nil
}

func romeOperator(s []string) ([]int, error) {
	arrNum := make([]int, len(s))
	var postNum, num int
	roman := make(map[byte]int)
	roman['I'] = 1
	roman['V'] = 5
	roman['X'] = 10

	for r := 0; r < len(s); r++ {
		for i := len(s[r]) - 1; i >= 0; i-- {
			num = roman[s[r][i]]
			if postNum <= num {
				arrNum[r] += num
			} else {
				arrNum[r] -= num
			}
			postNum = num
		}
	}

	for i := 0; i < len(arrNum); i++ {
		if arrNum[i] > 10 || arrNum[i] < 1 {
			return nil, fmt.Errorf("rome number is more than 10 or below 1")
		}
	}

	return arrNum, nil
}

func calculation(num []int, sign []string) (int, error) {
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
			return 0, fmt.Errorf("you cant divide by 0")
		}
	}
	return res, nil
}
