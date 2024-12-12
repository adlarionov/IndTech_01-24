package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// 16) Перевод числа из одной системы счисления в другую
func task_16(number string, fromBase, toBase int) (string, error) {
	decimalValue, err := strconv.ParseInt(number, fromBase, 64)
	if err != nil {
		return "", err
	}

	if toBase == 10 {
		return fmt.Sprintf("%d", decimalValue), nil
	}

	var result string
	for decimalValue > 0 {
		remainder := decimalValue % int64(toBase)
		result = string('0'+remainder) + result
		decimalValue /= int64(toBase)
	}
	return result, nil
}

// 17) Нахождение корней квадратного уравнения ax^2 + bx + c = 0
func task_17(a, b, c float64) (interface{}, interface{}) {
	discriminant := b*b - 4*a*c

	if discriminant > 0 {
		x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		return x1, x2
	} else if discriminant == 0 {
		x1 := -b / (2 * a)
		return x1, nil
	} else {
		realPart := -b / (2 * a)
		imaginaryPart := math.Sqrt(-discriminant) / (2 * a)
		return fmt.Sprintf("%.2f + %.2fi", realPart, imaginaryPart), fmt.Sprintf("%.2f - %.2fi", realPart, imaginaryPart)
	}
}

// 18) Сортировка массива по абсолютным значениям
func task_18(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return math.Abs(float64(arr[i])) < math.Abs(float64(arr[j]))
	})
	return arr
}

// 19) Объединение двух отсортированных массивов
func task_19(arr1, arr2 []int) []int {
	result := make([]int, len(arr1)+len(arr2))
	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result[k] = arr1[i]
			i++
		} else {
			result[k] = arr2[j]
			j++
		}
		k++
	}

	for i < len(arr1) {
		result[k] = arr1[i]
		i++
		k++
	}

	for j < len(arr2) {
		result[k] = arr2[j]
		j++
		k++
	}

	return result
}

// 20) Поиск первой позиции вхождения одной строки в другую
func task_20(haystack, needle string) int {
	return strings.Index(haystack, needle)
}

// 21) Математические операции (+, -, *, /, ^, %)
func task_21(a, b float64, operator string) (float64, error) {
	var result float64
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		result = a / b
	case "^":
		result = math.Pow(a, b)
	case "%":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		result = math.Mod(a, b)
	default:
		return 0, fmt.Errorf("недопустимая операция: %s", operator)
	}
	return result, nil
}

// 22) Проверка, является ли строка палиндромом
func task_22(s string) bool {
	re := regexp.MustCompile(`[^\w]`)
	s = re.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// 23) Определение области пересечения трех отрезков
func task_23(a1, b1, a2, b2, a3, b3 int) bool {
	left := math.Max(float64(a1), math.Max(float64(a2), float64(a3)))
	right := math.Min(float64(b1), math.Min(float64(b2), float64(b3)))

	return left <= right
}

// 24) Поиск самого длинного слова в предложении
func task_24(sentence string) string {
	re := regexp.MustCompile(`[^\w\s]`)
	sentence = re.ReplaceAllString(sentence, "")
	words := strings.Fields(sentence)

	longestWord := ""
	for _, word := range words {
		if len(word) > len(longestWord) {
			longestWord = word
		}
	}
	return longestWord
}

// 25) Проверка, является ли год високосным
func task_25(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	}
	return false
}

// 26) Вывод всех чисел Фибоначчи, не превышающих заданного значения
func task_26(limit int) []int {
	fib := []int{0, 1}
	for {
		next := fib[len(fib)-1] + fib[len(fib)-2]
		if next > limit {
			break
		}
		fib = append(fib, next)
	}
	return fib
}

// 27) Поиск простых чисел в диапазоне между двумя числами
func task_27(start, end int) []int {
	primes := []int{}
	for num := start; num <= end; num++ {
		if num < 2 {
			continue
		}
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
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

// 28) Числа Армстронга в заданном диапазоне
func task_28(start, end int) []int {
	armstrongNumbers := []int{}
	for num := start; num <= end; num++ {
		digits := strconv.Itoa(num)
		numDigits := len(digits)
		sum := 0
		for _, digit := range digits {
			digitValue, _ := strconv.Atoi(string(digit))
			sum += int(math.Pow(float64(digitValue), float64(numDigits)))
		}
		if sum == num {
			armstrongNumbers = append(armstrongNumbers, num)
		}
	}
	return armstrongNumbers
}

// 29) Реверс строки без встроенных функций
func task_29(s string) string {
	var result strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		result.WriteByte(s[i])
	}
	return result.String()
}

// 30) Алгоритм Евклида для нахождения НОД с использованием цикла
func task_30(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	result1, err := task_16("1010", 2, 10)
	if err != nil {
		fmt.Println("Ошибка в задаче 1:", err)
	} else {
		fmt.Println("Задача 16: Перевод числа из системы счисления 2 в 10 =", result1)
	}

	x1, x2 := task_17(1, -3, 2)
	fmt.Println("Задача 17: Корни квадратного уравнения = ", x1, x2)
	fmt.Println("Задача 18: Сортировка массива по абсолютным значениям =", task_18([]int{-10, 5, -3, 7, -2}))
	fmt.Println("Задача 19: Объединение двух отсортированных массивов =", task_19([]int{1, 3, 5, 7}, []int{2, 4, 6, 8}))
	fmt.Println("Задача 20: Первая позиция вхождения строки =", task_20("hello world", "world"))

	fmt.Println()

	result21, err := task_21(10, 2, "+")
	if err != nil {
		fmt.Println("Ошибка в задаче 21:", err)
	} else {
		fmt.Println("Задача 21: 10 + 2 =", result21)
	}
	fmt.Println("Задача 22: Палиндром =", task_22("A man, a plan, a canal, Panama"))
	fmt.Println("Задача 23: Пересечение отрезков =", task_23(1, 5, 3, 7, 4, 6))
	fmt.Println("Задача 24: Самое длинное слово =", task_24("The quick brown fox jumped over the lazy dog!"))
	fmt.Println("Задача 25: Високосный год =", task_25(2024))

	fmt.Println()

	fmt.Println("Задача 26: Числа Фибоначчи, не превышающие 100 =", task_26(100))
	fmt.Println("Задача 27: Простые числа от 10 до 50 =", task_27(10, 50))
	fmt.Println("Задача 28: Числа Армстронга от 100 до 1000 =", task_28(100, 1000))
	fmt.Println("Задача 29: Реверс строки 'Hello, World!' =", task_29("Hello, World!"))
	fmt.Println("Задача 30: НОД чисел 56 и 98 =", task_30(56, 98))
}
