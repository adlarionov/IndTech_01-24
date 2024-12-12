package main

import (
	"fmt"
	"math"
	"strings"
)

// 1) Сумма цифр числа
func task_1(num int) int {
	if num == 0 {
		return 0
	}
	return num%10 + task_1(num/10)
}

// 2) Преобразование температуры из Цельсия в Фаренгейты и обратно
func task_2(celsius bool, temp float64) float64 {
	if celsius {
		return temp*9/5 + 32 // из Цельсия в Фаренгейты
	}
	return (temp - 32) * 5 / 9 // из Фаренгейтов в Цельсия
}

// 3) Удвоение каждого элемента массива
func task_3(arr []int) []int {
	arr[0] = arr[0] * 2
	if len(arr) == 1 {
		return arr
	}
	return append(arr[:1], task_3(arr[1:])...)

}

// 4) Объединение строк
func task_4(strs []string) string {
	return strings.Join(strs, " ")
}

// 5) Расстояние между двумя точками в 2D пространстве
func task_5(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// 6) Проверка на четность или нечетность
func task_6(n int) string {
	if n%2 == 0 {
		return "Четное"
	}
	return "Нечетное"
}

// 7) Проверка на високосный год
func task_7(year int) string {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return "Високосный"
	}
	return "Не високосный"
}

// 8) Наибольшее из трех чисел
func task_8(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= c {
		return b
	}
	return c
}

// 9) Определение возрастной группы
func task_9(age int) string {
	if age <= 12 {
		return "Ребенок" // до 12 лет
	} else if age <= 17 {
		return "Подросток" // с 13 до 17 лет
	} else if age <= 64 {
		return "Взрослый" // с 18 до 64 лет
	}
	return "Пожилой" // 65 и старше
}

// 10) Проверка делимости на 3 и 5
func task_10(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "Делится на 3 и 5"
	}
	return "Не делится на 3 и 5"
}

// 11) Факториал числа
func task_11(factorial int) int {
	result := 1

	for factorial > 0 {
		result *= factorial
		factorial--
	}

	return result
}

// 12) Числа Фибоначчи
func task_12(n int) []int {
	fib := make([]int, n)
	fib[0] = 0
	if n > 1 {
		fib[1] = 1
	}
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

// 13) Реверс массива
func task_13(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
}

// 14) Поиск простых чисел
func task_14(limit int) []int {
	primes := []int{}
	for num := 2; num <= limit; num++ {
		isPrime := true
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, num)
		}
	}
	return primes
}

// 15) Сумма чисел в массиве
func task_15(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println("Задача 1: Сумма цифр числа =", task_1(1234))
	fmt.Println("Задача 2: Температура =", task_2(true, 36.6))
	fmt.Println("Задача 3: Удвоение массива =", task_3([]int{1, 2, 4, 8}))
	fmt.Println("Задача 4: Объединение строк =", task_4([]string{"Hello", "world"}))
	fmt.Println("Задача 5: Расстояние между точками =", task_5(1, 2, 4, 6))

	fmt.Println()

	fmt.Println("Задача 6: Проверка на четность или нечетность =", task_6(7))
	fmt.Println("Задача 7: Проверка на високосный год =", task_7(2024))
	fmt.Println("Задача 8: Наибольшее из трех чисел =", task_8(3, 7, 5))
	fmt.Println("Задача 9: Определение возрастной группы =", task_9(25))
	fmt.Println("Задача 10: Проверка делимости на 3 и 5 =", task_10(15))

	fmt.Println()

	fmt.Println("Задача 11: Факториал числа =", task_11(5))
	fmt.Println("Задача 12: Числа Фибоначчи =", task_12(7))
	fmt.Println("Задача 13: Реверс массива =", task_13([]int{1, 2, 3, 4, 5}))
	fmt.Println("Задача 14: Поиск простых чисел =", task_14(20))
	fmt.Println("Задача 15: Сумма чисел в массиве =", task_15([]int{1, 2, 3, 4, 5}))
}
