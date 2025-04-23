package main

import (
	"fmt"
)

func main() {
	currency := map[string]map[string]float64{} // создаем мар в котором будут храниться курсы валют в виде мар
	currency["USD"] = map[string]float64{       // добавляем в мар другую мап в которой хранятся курсы USD по отношению к другим валютам
		"EUR": 0.87,  // в одном долларе 0.87 евро
		"RUB": 81.35, // в одном долларе 81.35 рублей
	}
	currency["EUR"] = map[string]float64{
		"USD": 1.15,  // в оддном евро 1.15 доллара
		"RUB": 93.26, // в оддном евро 93.26 рублей
	}
	currency["RUB"] = map[string]float64{
		"USD": 0.012, // в одном рубле 0.012 долларов
		"EUR": 0.011, // в одном рубле 0.011 евро
	}
	firstCurrency := inputCurrency()                     // считываем первую валюту
	sum := inputSum()                                    // считываем сумму
	secondCurrency := inputSecondCurrency(firstCurrency) // считываем вторую валюту

	fmt.Printf("Резултат конвертации = %0.2f", calculate(&currency, firstCurrency, sum, secondCurrency))

}

func inputCurrency() string {
	var inputCurrency string // определяем переменную которую будем возвращать
	for {
		fmt.Print("Введите первую валюту EUR, USD, RUB: ")
		fmt.Scan(&inputCurrency)                                                        // сканируем ввод пользователя и записываем в переменную
		if inputCurrency != "EUR" && inputCurrency != "USD" && inputCurrency != "RUB" { // проверка правильности введеной валюты
			fmt.Println("Неверно введена валюта")
			continue

		}
		return inputCurrency // возвращаем переменную
	}
}
func inputSum() float64 {
	var inputSum float64
	for {
		fmt.Print("Введите сумму: ")
		fmt.Scan(&inputSum)
		if inputSum < 0 {
			fmt.Println("Сумма не может быть меньше 0")
			continue
		}
		return inputSum
	}
}
func inputSecondCurrency(firstCurrency string) string {
	var inputSecondCurrency string
first: // ставим лэйбл для того чтобы могли при неправильном вводе второй валюты повторно запросить ввод, при помощи continue first
	for {
		switch firstCurrency { // реализуем switch для того чтобы выводдить пользователю доступные варианты
		case "USD":
			fmt.Println("Введите вторую валюту EUR,RUB")
			fmt.Scan(&inputSecondCurrency)
			if inputSecondCurrency != "EUR" && inputSecondCurrency != "RUB" {
				fmt.Println("Неверно введена валюта")
				continue first
			}
		case "EUR":
			fmt.Println("Введите вторую валюту USD,RUB")
			fmt.Scan(&inputSecondCurrency)
			if inputSecondCurrency != "USD" && inputSecondCurrency != "RUB" {
				fmt.Println("Неверно введена валюта")
				continue first
			}
		case "RUB":
			fmt.Println("Введите вторую валюту EUR,USD")
			fmt.Scan(&inputSecondCurrency)
			if inputSecondCurrency != "USD" && inputSecondCurrency != "EUR" {
				fmt.Println("Неверно введена валюта")
				continue first
			}
		}
		return inputSecondCurrency
	}
}

func calculate(currency *map[string]map[string]float64, firstCurrency string, summ float64, secondCurrency string) float64 {
	// принемаем мар внутри другой мар, в качестве ключей к ним используем firstCurrency и secondCurrency
	result := summ * (*currency)[firstCurrency][secondCurrency] // расчитываем результат путем умножения суммы которую получили от пользователя на
	// курс который храниться в мар
	return result
}
