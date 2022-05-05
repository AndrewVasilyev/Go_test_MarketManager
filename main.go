package main

import (
	"fmt"
	"strings"
)

func lenSubStr(s string) int {
	size := 0
	currsize := 0
	for j := range s {
		if currsize == len(s) {
			return currsize
		}
		if currsize > size {
			size = currsize
		}
		var arr [2]byte
		idx := 0
		arr[idx] = s[j]
		idx++
		currsize = 1
		for i := j + 1; i < len(s); i++ {
			if arr[0] != s[i] && arr[1] != s[i] {
				if idx < len(arr) {
					arr[idx] = s[i]
					currsize++
					idx++
					continue
				}
				break
			}
			currsize++
		}
	}
	return size
}

func main() {

	fmt.Println("Enter string")
	var initStr string

	fmt.Scanf("%s/n", &initStr)

	initStr = strings.ToLower(initStr)

	fmt.Println(lenSubStr(initStr))
}
