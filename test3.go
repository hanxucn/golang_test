package main

import (
	"fmt"
	"math"
	"strconv"
)


func stableNameGenerate(value int32) string {
	randNum := simulatePseudoEncrypt(value)
	advisor := int32(math.Pow10(5))
	reminder := randNum % advisor
	name := "user" + strconv.Itoa(int(reminder))
	return name
}

func main() {
	userID := []int32{1000001, 1000002, 1000003, 1000004}
	for _, username := range userID{
		name := stableNameGenerate(username)
		fmt.Println(name)
	}
}
