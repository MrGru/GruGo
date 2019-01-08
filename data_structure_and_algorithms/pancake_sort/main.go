package main

import "fmt"

func main() {
	list := data{28, 11, 59, -26, 503, 158, 997, 193, -23, 44}
	fmt.Println("\n--- Unsorted ---\n\n", list)
	list.pancakeSort()
	fmt.Println("\n--- Sorted ---\n\n", list)
}

type data []int32

func (dataList data) pancakeSort() {
	for uns := len(dataList) - 1; uns > 0; uns-- {
		lx, lg := 0, dataList[0]
		for i := 1; i < uns; i++ {
			if dataList[i] > lg {
				lx, lg = i, dataList[i]
			}
		}
		dataList.flip(lx)
		dataList.flip(uns)
	}
}
func (dataList data) flip(r int) {
	for l := 0; l < r; l, r = l+1, r-1 {
		dataList[l], dataList[r] = dataList[r], dataList[l]
	}
}
