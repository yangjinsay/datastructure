package main

import "fmt"

/**
 * In-older: left -> root -> right
 * [1,null,2,3]
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
func main(){
	root := &TreeNode{}
	root.Left = nil
	root.Val = 1
	right := &TreeNode{}
	left  := &TreeNode{}
	root.Right = right
	root.Right.Left = left
	right.Left.Val = 3
	root.Right.Val = 2

	r := inorderTraversal(root)
	fmt.Println(r)
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	func inorderTree(root *TreeNode){
		inorderTree(root.Left)
		result = append(result, root.Val)
		inorderTree(root.Right)
	}
	inorderTree(root)
	return result
}
