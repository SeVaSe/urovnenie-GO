package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	for {
		var x float64

		// Запросить значение X у пользователя
		fmt.Print("Введите значение X (или 'exit' для выхода): ")
		// Считываем введенную строку
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			continue
		}

		// Преобразуем введенную строку в число
		x, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Ошибка преобразования в число:", err)
			continue
		}

		// Проверить ОДЗ (Область допустимых значений)
		if x == 4 || x == 0 {
			fmt.Println("X не может быть равен 4 или 0, так как это приведет к делению на ноль.")
			continue
		}

		// Вычислить выражение
		y := math.Pow(2*x, 5) + 5*x/(x-4) + 1/x + math.Sqrt(x+7)

		// Проверить ошибки при вычислениях
		if math.IsNaN(y) || math.IsInf(y, 0) {
			fmt.Println("Ошибка при вычислении выражения. Результат неопределен.")
			continue
		}

		// Вывести результат
		fmt.Printf("При X = %.2f, значение y = %.2f\n", x, y)

		// Проверяем, если введено "exit", то завершаем программу
		if input == "exit" {
			break
		}

	}
}
