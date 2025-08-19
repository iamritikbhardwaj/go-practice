package main

import (
	"fmt"

	"github.com/iamritikbhardwaj/commonPractices/pracPrograms"
)

func main() {
	fmt.Println("Hello World")
	// pracPrograms.Sleeps()
	nums := []int{5, 4, 3, 2, 1}
	pracPrograms.MinHeap(nums)
	// pracPrograms.SlowDown()
}
