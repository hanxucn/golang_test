package main

import (
	"fmt"
	"math"
	"strconv"
)


func stableNameGenerate(value int32) int32 {
	randNum := simulatePseudoEncrypt(value)
	advisor := int32(math.Pow10(5))
	reminder := randNum % advisor
	return reminder
}

func main() {
	userID := []int32{1000001, 1000002, 1000003, 1000004}
	for _, username := range userID{
		processID := stableNameGenerate(username)
		name := "user" + strconv.Itoa(int(processID))
		fmt.Println(name)
	}
}
