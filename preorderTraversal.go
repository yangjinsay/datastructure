/**
 * preorder: root -> left -> right
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append([]int{root.Val}, preorderTraversal(root.Left)...), preorderTraversal(root.Right)...)
}
