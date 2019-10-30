package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	minVal := 0
	val := root.Val
	maxVal := 0
	left := root
	right := root
	for {
		if left.Left != nil {
			if left.Left.Val > left.Right.Val {
				return  false
			}
			left = left.Left
		} else {
			minVal = left.Val
			break
		}
	}
	for {
		if right.Right != nil {
		if right.Left.Val > right.Right.Val {
			return false
		}
			right = right.Right
		} else {
			maxVal = right.Val
			break
		}
	}

	if minVal < val && val < maxVal {
		return true
	}
	return false
}

func main() {
	left := &TreeNode{}
	right := &TreeNode{}
	root := &TreeNode{}
	left.Val = 1
	right.Val = 3
	root.Val = 2

	root.Left = left
	root.Right = right
	re := isValidBST(root)
	fmt.Println(re)
}
