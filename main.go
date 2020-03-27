/*
 * @Description:
 * @Author: chinphy
 * @Date: 2020-03-06 14:28:26
 */
package main

import (
	"fmt"
	"leetcode-go/solution"
)

func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	solution.RotateMatrix(mat)
	fmt.Println("nums let:", mat)
}
