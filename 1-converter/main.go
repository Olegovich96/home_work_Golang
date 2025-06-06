package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("__Вас приветствует онлайн конвертер валют!__")
	createMenu()
	// getFirstCurrency := getInputCurrency(firstCurrency, secondCurrency)

	// getSecondCurrency := getInputCurrency(getFirstCurrency, secondCurrency)

}

func createMenu() {
	const UsdRuble float64 = 78.95
	const UsdEuro float64 = 0.88
	const EuroRuble float64 = 89.88
	// countConvert, firstCurrency, secondCurrency := getCountCalculate()
	// calculateCurrency(countConvert, firstCurrency, secondCurrency)
	var firstCurrency string
	var secondCurrency string
	var countConvert float64
	for {
		fmt.Println("Выберите нужный пункт меню:")
		fmt.Println("1. Ввод валюты которую необходимо конверитировать")
		fmt.Println("2. Ввод валюты в которую необходимо конвертировать")
		fmt.Println("3. Количество валюты которую нужно конвертировать")
		fmt.Println("4. Выход")
		var userChoices int
		fmt.Scan(&userChoices)

		switch userChoices {
		case 1:
			firstCurrency = getInputCurrency("", "")
			fmt.Println("Вы выбрали валюту из которой хотите конвертировать:", firstCurrency)
		case 2:
			if firstCurrency == "" {
				fmt.Println("Сначала перейдите в пункт 1, для того чтобы заполнить валюту из которой вы хотите произвести конвертациб")
				continue
			}
			secondCurrency = getInputCurrency(firstCurrency, "")
			fmt.Println("Вы выбрали валюту в которую хотите конвертировать:", secondCurrency)
		case 3:
			if firstCurrency == "" && secondCurrency == "" {
				fmt.Println("Вы пропустили пункт №1 и пункт №2")
				continue
			}
			countConvert = getCountConvert()
			result := calculateCurrency(firstCurrency, secondCurrency, countConvert, UsdRuble, UsdEuro, EuroRuble)
			fmt.Printf("Результат конвертации %.2f %s в %s %.2f \n", countConvert, firstCurrency, secondCurrency, result)

		case 4:
			return
		default:
			fmt.Println("Такого варианта в меню нет, попробуйте еще раз....")
		}

	}
}

func getInputCurrency(firstCurrency, secondCurrency string) string {
	currencies := "usd, eur, rub"
	if firstCurrency == "" && secondCurrency == "" {
		for {
			fmt.Println("Пожалуйста введите валюту которую хотите конвертировать: %v", currencies)
			var getUserInput string
			fmt.Scan(&getUserInput)
			getUserInput = strings.ToLower(getUserInput)
			checking := checkCurrency(getUserInput)
			if checking {
				return getUserInput
			}
			fmt.Println("Введите значение валюты которое совпадает с одним значением из списка: usd, eur, rub")
		}
	} else if firstCurrency != "" && secondCurrency == "" {
		filtering := strings.Replace(currencies, firstCurrency+", ", "", -1)
		for {
			fmt.Println("Пожалуйста введите валюту которую хотите конвертировать: %v", filtering)
			var getUserInput string
			fmt.Scan(&getUserInput)
			checking := checkCurrency(getUserInput)
			if checking && firstCurrency != getUserInput {
				return getUserInput
			}
			fmt.Println("Валюты не должны совпадать! Проверьте совпадает ли название валюты со списком: usd, eur, rub")
		}
	}
	return ""
}

func checkCurrency(currency string) bool {
	if currency == "usd" || currency == "eur" || currency == "rub" {
		return true
	}
	return false
}

func getCountConvert() float64 {
	var amountConvert float64
	fmt.Println("Введите сумму которую необходимо конвертировать")
	for {
		_, err := fmt.Scan(&amountConvert)
		if err != nil {
			fmt.Println("Нерверный ввод. Введите число например 80, или 120.1")
		}
		return amountConvert
	}
}

func calculateCurrency(firstCurrency, secondCurrency string, amountConvert float64, usdRub, eurUsd, eurRub float64) float64 {
	switch {
	case firstCurrency == "usd" && secondCurrency == "rub" || firstCurrency == "rub" && secondCurrency == "usd":
		result := amountConvert * usdRub
		return result
	case firstCurrency == "eur" && secondCurrency == "usd" || firstCurrency == "usd" && secondCurrency == "eur":
		result := amountConvert * eurUsd
		return result
	case firstCurrency == "eur" && secondCurrency == "rub" || firstCurrency == "rub" && secondCurrency == "eur":
		result := amountConvert * eurRub
		return result
	default:
		return 0.0
	}
}
