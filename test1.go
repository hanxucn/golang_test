package main

import "fmt"

func twoSum(a []int, b []int, v int) bool {
	if a == nil || b == nil{
		return false
	}
	for i := 0; i < len(a); i++{
		num := v - a[i]
		if isContain(b, num){
			return true
		}
	}
	return false
}

func isContain(arr []int, item int) bool {
	for _, singleItem := range arr {
		if singleItem == item {
			return true
		}

	}
	return false
}

// map version
//func twoSum(a []int, b []int, v int) bool {
//	if a == nil || b == nil{
//		return false
//	}
//	numMap := make(map[int]int)
//	for i := 0; i < len(a); i++{
//		num := a[i]
//		numMap[num] = 1
//	}
//	for j := 0; j<len(b); j++{
//		key := v - b[j]
//		if numMap[key] == 1{
//			return true
//		}
//	}
//	return false
//}

func main() {
	a := []int{10, 40, 5, 200}
	b := []int{234, 5, 2, 148, 23}
	res1 := twoSum(a, b, 42)
	fmt.Println(res1)

	a1 := []int{7, 6, 100, 22, 3, 7}
	b1 := []int{6, 2}
	//a2 := []int{}
	res2 := twoSum(a1, b1, 11)
	fmt.Println(res2)
}
