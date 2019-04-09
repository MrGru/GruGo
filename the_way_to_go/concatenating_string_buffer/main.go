package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var buffer bytes.Buffer
	for {
		if s, ok := getNextString(); ok {
			buffer.WriteString(s)
		} else {
			break
		}
	}
	fmt.Println(buffer.String(), "\n")
}

func getNextString() (string, bool) {
	rand.Seed(time.Now().Unix())
	if rand.Int()%2 == 0 {
		return "Test", true
	}
	return "", false
}
