package pracPrograms

import (
	"fmt"
	"strings"
	"time"
)

const delay = 2 * time.Second

func print(msg string) {
	for i, _ := range msg {
		for j := 0; j <= i; j++ {
			fmt.Print(strings.Split(msg, "")[i])
		}
		time.Sleep(delay)
	}
}

func SlowDown() {
	str := "hello world"
	print(str)
}
