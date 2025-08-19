package main

import (
	"fmt"

	"github.com/iamritikbhardwaj/commonPractices/pracPrograms"
)

func main() {
	fmt.Println("Hello World")
	// pracPrograms.Sleeps()
	h := &pracPrograms.MinHeap{}

	h.Insert(10)
	h.Insert(4)
	h.Insert(15)
	h.Insert(20)
	h.Insert(0)
	h.Insert(8)

	fmt.Println("Heap elements:")
	h.Print()
	h.Remove()
	h.Remove()
	h.Remove()
	h.Remove()
	h.Remove()
	h.Remove()
	// pracPrograms.SlowDown()
}
