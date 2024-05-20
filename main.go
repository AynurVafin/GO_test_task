package main

import (
	"fmt"
	"sort"
)

func Equal(tot_num, num_col []int64) bool {
	for i, v := range tot_num {
		if v != num_col[i] {
			return false
		}
	}
	return true
}

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	tot_num := make([]int64, n)
	num_col := make([]int64, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var x int64
			_, _ = fmt.Scan(&x)
			tot_num[i] += x
			num_col[j] += x
		}
	}
	sort.Slice(tot_num, func(i, j int) bool { return tot_num[i] < tot_num[j] })
	sort.Slice(num_col, func(i, j int) bool { return num_col[i] < num_col[j] })
	if Equal(tot_num, num_col) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
