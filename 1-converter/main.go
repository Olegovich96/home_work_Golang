package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("__Вас приветствует онлайн конвертер валют!__")
	const UsdRuble float64 = 78.95
	const UsdEuro float64 = 0.88
	const EuroRuble float64 = 89.88
	countConvert, firstCurrency, secondCurrency := getCountCalculate()
	calculateCurrency(countConvert, firstCurrency, secondCurrency)
}

func getCountCalculate() (float64, string, string) {
	var countConvert float64
	var firstCurrency, secondCurrency string
	fmt.Println("Введите сумму, которую хотите конвертировать?")
	fmt.Scan(&countConvert)
	fmt.Println("Введите валюту, которую хотите конверивать: usd, euro, ruble")
	fmt.Scan(&firstCurrency)
	fmt.Println("Введите валюту, в которую конвертировать: usd, euro, ruble")
	fmt.Scan(&secondCurrency)

	firstCurrency = strings.ToLower(firstCurrency)
	secondCurrency = strings.ToLower(firstCurrency)

	return countConvert, firstCurrency, secondCurrency
}

func calculateCurrency(countExchange float64, firstCurrency, secondCurrency string) float64 {
	return 0.0
}
