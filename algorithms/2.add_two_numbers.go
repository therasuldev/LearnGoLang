package main

import "fmt"

func main2() {
	var l1 = []int{1, 2, 3}
	var l2 = []int{1, 2, 3}
	fmt.Println(addTwoNumbers(l1, l2))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 []int, l2 []int) []int {
	var result []int
	for i := 0; i < len(l1); i++ {
		for j := 0; j < len(l2); j++ {
			if i != j {
				continue
			}
			result = append(result, l1[i]+l2[j])

		}
	}
	for i2, j2 := 0, len(result)-1; i2 < j2; i2, j2 = i2+1, j2-1 {
		result[i2], result[j2] = result[j2], result[i2]
	}
	return result

}
