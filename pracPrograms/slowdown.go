package pracPrograms

import (
	"fmt"
	"time"
)

const delay = 7 * time.Second

func slowDown(str string) {
	for i, _ := range str {
		print(string(str[i]))
	}
}

func print(msg string) {
	fmt.Print(msg)
	time.Sleep(delay)
}

func SlowDown() {
	str := "hello world"
	slowDown(str)
}
