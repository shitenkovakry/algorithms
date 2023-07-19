package main

import "fmt"

func BinarySearch(array []int, hiddenNumber int, lowValue int, hihgValue int) int {
	if lowValue > hihgValue {
		return -1
	}

	middleElement := (lowValue + hihgValue) / 2

	if array[middleElement] == hiddenNumber {
		return middleElement
	} else if array[middleElement] > hiddenNumber {
		return BinarySearch(array, hiddenNumber, lowValue, hihgValue-1) // ищем в левой половине
	} else {
		return BinarySearch(array, hiddenNumber, middleElement+1, hihgValue) // ищем в правой половине
	}
}

func main() {
	array := []int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	hiddenNumber := 7
	result := BinarySearch(array, hiddenNumber, 0, len(array)-1)

	if result != -1 {
		fmt.Println("элемент ", hiddenNumber, "найден по индексу ", result)
	} else {
		fmt.Println("элемент ", hiddenNumber, "не найден в массиве")
	}
}
