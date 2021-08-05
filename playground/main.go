package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}





func findNumOfValidWords(words []string, puzzles []string) []int {
	var hash = make(map[int]int)
	for _, word := range words {
		bits := getBits(word)
		hash[bits]++
	}
	result := make([]int, len(puzzles))
	for k, v := range puzzles {
		puzzleBit := getBits(v)
		first := getBits(string(v[0]))

		n := puzzleBit
		for n > 0 {
			if (n&first != 0) && (hash[n] > 0) {
				result[k] += hash[n]
			}
			n = (n - 1) & puzzleBit
		}
	}
	return result
}

// 获取字符串的位表示
func getBits(word string) (res int) {
	for _, c := range word {
		offset := c - 'a'
		status := 1 << offset
		res = res | status
	}
	return
}

func main() {
	fmt.Println(permute([]int{5, 4, 6, 2}))
}
