package main

import (
	"errors"
	"fmt"
)

const USDinEUR = 0.88
const USDinRUB = 82.17
const EURinRUB = (1 / USDinEUR) * USDinRUB

func main() {

	for {
		firstValute, summ, secondValute, error := inputUser()
		if error != nil {
			fmt.Println("некоректный ввод")
			continue
		} else {
			fmt.Printf("Резултат конвертации = %0.2f", calculate(firstValute, summ, secondValute))
			break
		}
	}

}
func inputUser() (string, float64, string, error) {
	var firstInput string
	var secondInput float64
	var thirdInput string
	fmt.Print("Введите первую валюту EUR, USD, RUB: ")
	fmt.Scan(&firstInput)
	if firstInput != "EUR" && firstInput != "USD" && firstInput != "RUB" {
		return "", 0.0, "", errors.New("Error")
	}

	fmt.Print("Введите сумму: ")
	fmt.Scan(&secondInput)
	if secondInput < 0 {
		return "", 0.0, "", errors.New("Error")
	}

	switch firstInput {
	case "EUR":
		fmt.Print("Введите вторую валюту USD, RUB: ")
		fmt.Scan(&thirdInput)
		if thirdInput != "USD" && thirdInput != "RUB" {
			return "", 0.0, "", errors.New("Error")
		}
	case "USD":
		fmt.Print("Введите вторую валюту EUR, RUB: ")
		fmt.Scan(&thirdInput)
		if thirdInput != "EUR" && thirdInput != "RUB" {
			return "", 0.0, "", errors.New("Error")
		}
	case "RUB":
		fmt.Print("Введите вторую валюту USD, EUR: ")
		fmt.Scan(&thirdInput)
		if thirdInput != "USD" && thirdInput != "EUR" {
			return "", 0.0, "", errors.New("Error")
		}
	}
	return firstInput, secondInput, thirdInput, nil
}

func calculate(firstValute string, summ float64, secondValute string) float64 {
	first := firstValute
	chislo := summ
	two := secondValute
	var result float64

	switch first {
	case "EUR":
		switch two {
		case "USD":
			result = chislo * 1 / USDinEUR

		case "RUB":
			result = chislo * EURinRUB
		}
	case "USD":
		switch two {
		case "EUR":
			result = chislo * USDinEUR

		case "RUB":
			result = chislo * USDinRUB
		}
	case "RUB":
		switch two {
		case "EUR":
			result = chislo * 1 / EURinRUB

		case "USD":
			result = chislo * 1 / USDinRUB
		}

	}
	return result
}
