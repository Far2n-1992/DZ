package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	for {
		numbers, err := inputUser()
		if err != nil {
			fmt.Println("Ошибка вводда данных, попробуйте снова")
			continue
		}
		result, err := calculate(numbers)
		if err != nil {
			fmt.Println("Неверно выбрали операцию, попробуйте снова")
		} else {
			fmt.Println("Результат операции =", result)
		}
	}
}

func inputUser() ([]float64, error) {
	var input string
	numbers := []float64{}
	fmt.Print("Введите данные через запятую: ")
	fmt.Scan(&input)
	stringNumber := strings.Split(input, ",")
	for _, value := range stringNumber {
		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, errors.New("Error_input")
		}
		numbers = append(numbers, number)
	}
	return numbers, nil

}

func calculate(numbers []float64) (float64, error) {
	var operation string
	var result float64
	fmt.Print("Введите операцию AVG,SUM,MED: ")
	fmt.Scan(&operation)
	if operation != "AVG" && operation != "SUM" && operation != "MED" {
		return 0, errors.New("Error_operation")
	}
	switch operation {
	case "AVG":
		var summ float64
		for _, value := range numbers {
			summ += value
			result = summ / float64(len(numbers))
		}

	case "SUM":
		for _, value := range numbers {
			result += value
		}

	case "MED":
		sort.Float64s(numbers)
		allNumbers := len(numbers)
		if allNumbers%2 == 1 {
			result = numbers[allNumbers/2]
		} else {
			result = (numbers[allNumbers/2-1] + numbers[allNumbers/2]) / 2
		}
	}
	return result, nil
}
