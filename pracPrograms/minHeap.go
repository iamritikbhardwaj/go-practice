package pracPrograms

import "fmt"

func MinHeap(nums []int) {
	fmt.Println("MinHeap")
	i := len(nums) - 1

	for {
		if i <= 0 {
			break
		}
		if nums[i] < nums[i/2] {
			nums[i], nums[i/2] = nums[i/2], nums[i]
		} else {
			i--
		}
	}

	fmt.Println(nums)
}
