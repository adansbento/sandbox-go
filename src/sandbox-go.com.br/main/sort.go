package main

import (
	"fmt"
)

//Adans Bento
func main() {

	numbers := []int{8, 2, 4, 6, 9, 3, 1, 7, 9}
	fmt.Println(numbers)
	fmt.Println("============== Disorderly ==================")
	sortNum(numbers)
	fmt.Println("============== Sort ==================")
	fmt.Println(numbers)

}

//
//Implement quickSort
//method calls sortNum1 set the left and right of ordination
//
func sortNum(values []int) {
	sortNum1(values, 0, len(values)-1)
}

//
//Method receives an int array, left and right with the size of the array,
//if the left is greater than equal to the right, it leaves the method without doing anything,
//I chose the first value of the array will be the pivot the first time the method is called,
//runs the array if the pivot is greater than the current number, will be exchanged positions,
// then will again be called the method, changing the left to order the entire array
//
func sortNum1(values []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := values[l]
	i := l + 1

	for j := l; j <= r; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[l], values[i-1] = values[i-1], pivot

	sortNum1(values, l, i-2)
	sortNum1(values, i, r)
}
