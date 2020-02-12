package main

import "fmt"

// ex one
func twoSum(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	sliceIn := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(sliceIn, target))

}
