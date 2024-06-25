package main

import (
	"fmt"
	"strconv"
	"strings"
)

var romeDigits = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func toArabic(rome string) int {
	return romeDigits[rome]
}

func toRoman(arabic int) string {
	if arabic <= 0 {
		panic("Отрицательные числа или ноль не могут быть представлены римскими цифрами")
	}
	var result strings.Builder
	for arabic > 0 {
		for _, value := range []struct {
			value int
			digit string
		}{
			{10, "X"},
			{9, "IX"},
			{5, "V"},
			{4, "IV"},
			{1, "I"},
		} {
			if arabic >= value.value {
				arabic -= value.value
				result.WriteString(value.digit)
				break
			}
		}
	}
	return result.String()
}

func calculate(num1, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			panic("Деление на ноль")
		}
		return num1 / num2
	default:
		panic("Недопустимая арифметическая операция")
	}
}

func main() {
	var num1, num2 int
	var operator string
	var input1, input2 string

	//var array [3]string{input1, operator, input2}
	//fmt.Println(array)

	fmt.Println("Число, математическую операцию и второе число через пробел: ")
	fmt.Scanln(&input1, &operator, &input2)
	//fmt.Println("Введите второе число: ")
	//fmt.Scanln(&input2)
	//fmt.Println("Введите арифметическую операцию (+, -, *, /): ")
	//fmt.Scanln(&operator)
	//fmt.Println("Данные для рассчёта: ")
	//fmt.Scanln(&input1), (&operator), (&input2)

	if num, err := strconv.Atoi(input1); err == nil {
		num1 = num
	}

	if num, err := strconv.Atoi(input2); err == nil {
		num2 = num

	}

	if (num1 < 1 && num1 > 10) || (num2 < 1 && num2 > 10) {
		panic("Числа должны быть от 1 до 10")
	}

	result := calculate(num1, num2, operator)
	if romeDigits[input1] > 0 && romeDigits[input2] > 0 {
		num1 = romeDigits[input1]
		num2 = romeDigits[input2]
		arabicResult := calculate(num1, num2, operator)
		fmt.Println(toRoman(arabicResult))
	} else if romeDigits[input1] > 0 || romeDigits[input2] > 0 {
		panic("Нельзя использовать одновременно арабские и римские числа")
	} else {
		fmt.Println(result)
	}

}
