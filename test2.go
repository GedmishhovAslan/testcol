package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var imputString, firstElement, operator, secondElement, result string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	imputString = scanner.Text()
	parts := strings.Split(imputString, "\"")

	if imputString[0] != 34 {
		err := errors.New("Первое значение не подходит")
		fmt.Println("Ошибка:", err)
		return
	}

	switch len(parts) {
	case 3:
		test := strings.Fields(parts[2])
		firstElement = parts[1]
		operator = test[0]
		secondElement = test[1]
		firstElement = strings.Replace(firstElement, "\"", "", -1)
		secondElement = strings.Replace(secondElement, "\"", "", -1)
		result = "\""

		switch operator {
		case "*":
			counter, err1 := strconv.Atoi(secondElement)
			if err1 != nil || counter > 10 || counter < 1 {
				fmt.Println("Не корректное значение3")
				return
			}
			for i := 0; i < counter; i++ {
				result += firstElement
			}
			if len(result) > 40 {
				result = result[:40]
				result += "..."
			}
		case "/":
			counter, err1 := strconv.Atoi(secondElement)
			if err1 != nil || counter > 10 || counter < 1 {
				fmt.Println("Не корректное значение4")
				return
			}
			counter = len(firstElement) / counter
			result += firstElement[:counter]
			if len(result) > 40 {
				result = result[:40]
				result += "..."
			}
		default:
			err := errors.New("Не корректный оператор")
			fmt.Println("Ошибка:", err)
			return
		}

	case 5:
		firstElement = parts[1]
		operator = parts[2]
		secondElement = parts[3]
		firstElement = strings.Replace(firstElement, "\"", "", -1)
		secondElement = strings.Replace(secondElement, "\"", "", -1)
		result = "\""
		if len(firstElement) > 10 || len(secondElement) > 10 {
			fmt.Println("Длинна строки превышает 10 символов")
			return
		}

		switch operator {
		case " + ":
			result += firstElement + secondElement
			if len(result) > 40 {
				result = result[:40]
				result += "..."
			}

		case " - ":
			result += strings.Replace(firstElement, secondElement, "", -1)
			if len(result) > 40 {
				result = result[:40]
				result += "..."
			}
		default:
			err := errors.New("Не корректный оператор")
			fmt.Println("Ошибка:", err)
			return
		}
	default:
		err := errors.New("Не правильный ввод")
		fmt.Println("Ошибка:", err)
		return

	}

	result += "\""
	fmt.Println(result)
}
