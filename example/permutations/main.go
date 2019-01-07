package main

import "fmt"

func main() {
	inputArray := []string{"abc",
		"bef",
		"bcc",
		"bec",
		"bbc",
		"bdc"}
	result := permutations(inputArray)
	for index, val := range result {
		fmt.Println(index+1, ":", val)
	}
}

func permutations(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					arr[i], arr[n-1] = arr[n-1], arr[i]
				} else {
					arr[0], arr[n-1] = arr[n-1], arr[0]
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
