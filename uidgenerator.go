package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(randIntRange(10, 5))
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randIntRange(1000, 1000000))
	}
	return string(bytes)
}

func randInt() uint {
	return (1103515245*uint(time.Now().UTC().UnixNano()) + 12345) % (1 << 31)
}

func randIntRange(max, min uint) uint {
	return (randInt() % (max - min + 1)) + min
}
