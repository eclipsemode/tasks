package main

import "fmt"

//Условие задачи
//На Авито размещено множество товаров, каждый из которых представлен числом.
//У каждого покупателя есть потребность в товаре, также выраженная числом.
//Если точного товара нет, покупатель выбирает ближайший по значению товар, что вызывает неудовлетворённость, равную разнице между его потребностью и купленным товаром.
//Количество каждого товара не ограничено, и один товар могут купить несколько покупателей. Рассчитайте суммарную неудовлетворённость всех покупателей.
//
//Нужно написать функцию, которая примет на вход два массива: массив товаров и массив потребностей покупателей, вычислит сумму неудовлетворённостей всех покупателей и вернет результат в виде числа.
//
//Пример
//# Пример
//# ввод
//goods = [8, 3, 5]
//buyerNeeds = [5, 6]
//# вывод
//res = 1 # первый покупатель покупает товар 5 и его неудовлетворённость = 0, второй также покупает товар 5 и его неудовлетворённость = 6-5 = 1

func main() {
	res := getBuyersFrustration([]int{8, 3, 7, 5}, []int{5, 6, 6, 6})
	fmt.Println("RESULT", res)
}

func getBuyersFrustration(goods []int, needs []int) int {
	r := 0
	for _, n := range needs {
		pivot := abs(n - goods[0])    // опорный элемент каждую итерацию по needs всегда ставим на нулевой индекс goods
		for _, g := range goods[1:] { // пропускаем нулевой индекс так как в опорной переменной
			if abs(n-g) < pivot {
				pivot = abs(n - g)
			}
		}
		r += pivot // прибавляем лучший результат из возможных сравнив все значения между собой
	}

	return r
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Ниже код в котором я пытался сделать вместо квадратичной сложности - линейно логарифмическую через быструю сортировку и бинарный поиск с доп сравнениями, но так и не дошел до результата))

//func getBuyersFrustration(goods []int, needs []int) int {
//	sortedArr := quickSort(goods) // Сортируем товары
//
//	res := 0
//
//	for _, v := range needs {
//		b := binarySearch(sortedArr, v) // Ищем ближайший товар для каждого needs
//		res += b                        // Суммируем фрустрацию
//	}
//
//	return res
//}

//func binarySearch(arr []int, val int) int {
//	minI := 0
//	maxI := len(arr) - 1
//
//	for minI < maxI {
//		mid := minI + (maxI-minI)/2
//
//		if arr[mid] == val {
//			return 0
//		} else if arr[mid] > val {
//			maxI = mid - 1 // Ищем в левой половине
//		} else {
//			minI = mid + 1 // Ищем в правой половине
//		}
//	}
//
//	if arr[minI] == val {
//		return 0
//	}
//
//
//	maxValue := 0
//	if minI == 0 {
//		for _, v := range arr[minI : minI+1] {
//			if abs(val-v) > maxValue {
//				fmt.Println("FIRST", abs(val-v))
//				maxValue = abs(val - v)
//			}
//		}
//	} else if minI == len(arr)-1 {
//		for _, v := range arr[minI-1 : minI] {
//			if abs(val-v) > maxValue {
//				fmt.Println("SECOND", abs(val-v))
//				maxValue = abs(val - v)
//			}
//		}
//	} else {
//		for _, v := range arr[minI-1 : minI+1] {
//			if abs(val-v) > maxValue {
//				fmt.Println("THIRD", abs(val-v))
//				maxValue = abs(val - v)
//			}
//		}
//	}
//
//	return maxValue
//}
//
//func quickSort(nums []int) []int {
//	if len(nums) < 2 {
//		return nums // Базовый случай: массив из 0 или 1 элемента уже отсортирован
//	}
//
//	pivot := nums[0] // Опорный элемент
//	var less, greater []int
//
//	for _, v := range nums[1:] {
//		if v <= pivot {
//			less = append(less, v) // Элементы меньше или равные pivot
//		} else {
//			greater = append(greater, v) // Элементы больше pivot
//		}
//	}
//
//	// Рекурсивно сортируем less и greater, затем объединяем
//	result := append(quickSort(less), pivot)
//	result = append(result, quickSort(greater)...)
//
//	return result
//}
