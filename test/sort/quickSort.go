package main

import "fmt"

func quickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	x := arr[(l+r)/2]
	i := l - 1
	j := r + 1
	for {
		i++
		for arr[i] < x {
			i++
		}
		j--
		for arr[j] > x {
			j--
		}
		if i >= j {
			break
		}
		swap(arr, i, j)
	}
	quickSort(arr, l, j)
	quickSort(arr, j+1, r)
}

// swap 交换数组中两个元素的位置
func swap(arr []int, l int, r int) {
	arr[l], arr[r] = arr[r], arr[l]
}

func main() {
	arr := []int{1, 2, 10, 3, 9, 2}
	quickSort(arr, 0, len(arr)-1)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
