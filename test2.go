package main

import (
	"fmt"
	"strconv"
)


func simulatePseudoEncrypt(value int32) int32 {
	var l1 int32 = (value >> 16) & 65535
	var r1 int32 = value & 65535
	var l2, r2 int32
	var i int = 0
	for i < 3 {
		l2 = r1
		r2 = l1 ^ int32((float64((1366*r1 + 150889) % 714025) / 714025.0) * 32767)
		l1 = l2
		r1 = r2
		i = i + 1
	}
	return (r1 << 16) + l1
}

func userNameGenerate(userId int32) string {
	nameId := simulatePseudoEncrypt(userId)
	name := "user" + strconv.Itoa(int(nameId))
	return name
}

func main() {
	userID := []int32{1000001, 1000002, 1000003, 1000004}
	for _, username := range userID{
		res := userNameGenerate(username)
		fmt.Println(res)
	}
}
