package main

import (
	"fmt"
)

// 测试专用
func SingleNumber(nums []int) int {
	// TODO: implement
	map_v := make(map[int]int)
	m := 0
	for _, num := range nums {
		if _, ok := map_v[num]; !ok {
			map_v[num] = 1
		} else {
			map_v[num] += 1
		}
	}
	for k, v := range map_v {
		if v == 1 {
			m = k
			return m
		}
	}
	return 0
}

func main() {
	fmt.Println(SingleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(SingleNumber([]int{2, 2, 1}))
	fmt.Println(SingleNumber([]int{1}))
	fmt.Println(SingleNumber([]int{3, 2, 3, 4, 2}))
}
