package main

import "fmt"

func bubblesort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 3, 2, 1, 0}
	fmt.Println(bubblesort(arr))

}
