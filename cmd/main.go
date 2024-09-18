package main

import (
	"GoTest/pkg"
)

func main() {
	
	pkg.ThreadTest1()
}

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, x := range nums {
		if index, ok := mp[target-x]; ok {
			return []int{index, i}
		}
		mp[x] = i
	}
	return nil
}