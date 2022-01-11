package main

import "sort"

/*
二刷
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 28. 实现 strStr()
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == len(needle)-1 {
				return i
			}
		}
	}
	return -1
}

//78. 子集
func subsets(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	list := make([]int, 0)
	subsetsHelper(0, nums, list, &result)
	return result
}

func subsetsHelper(pos int, nums []int, list []int, result *[][]int) {

	data := make([]int, len(list))
	copy(data, list)
	*result = append(*result, data)

	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		subsetsHelper(i+1, nums, list, result)
		list = list[:len(list)-1]
	}
}

// 104. 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxDepthHelper(root)
}

func maxDepthHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepthHelper(root.Left)
	right := maxDepthHelper(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

// 110. 平衡二叉树
func isBalanced(root *TreeNode) bool {
	//分治法
	if root == nil {
		return true
	}
	_, b := isBalancedHelper(root)
	return b
}

func isBalancedHelper(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftD, leftB := isBalancedHelper(root.Left)
	rightD, rightB := isBalancedHelper(root.Right)
	if !leftB || !rightB {
		return 0, false
	}
	if leftD-rightD > 1 || rightD-leftD > 1 {
		return 0, false
	}

	if leftD > rightD {
		return leftD + 1, true
	}
	return rightD + 1, true

}

// 236. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, node := helper(root, p, q, 0)
	return node
}

func helper(root, p, q *TreeNode, count int) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}

	left, leftNode := helper(root.Left, p, q, count)
	right, rightNode := helper(root.Right, p, q, count)
	if left == 2 {
		return 2, leftNode
	}
	if right == 2 {
		return 2, rightNode
	}

	count = 0
	if root == p || root == q {
		count += 1
	}
	if left == 1 {
		count += 1
	}
	if right == 1 {
		count += 1
	}
	if count == 2 {
		return count, root
	}
	return count, nil
}

// 124. 二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {
	//贡献值的做法
	maxSum := root.Val

	var maxGain func(root *TreeNode) int
	maxGain = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftGain := max(maxGain(root.Left), 0)
		rightGain := max(maxGain(root.Right), 0)

		maxSum = max(leftGain+rightGain+root.Val, maxSum)

		return max(leftGain, rightGain) + root.Val
	}
	maxGain(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 102. 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {

	levelList := make([]*TreeNode, 0)
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	levelList = append(levelList, root)

	for len(levelList) > 0 {
		data := make([]int, 0)
		l := len(levelList)
		for i := 0; i < l; i++ {
			node := levelList[i]
			data = append(data, node.Val)
			if node.Left != nil {
				levelList = append(levelList, node.Left)
			}
			if node.Right != nil {
				levelList = append(levelList, node.Right)
			}
		}
		result = append(result, data)
		levelList = levelList[l:]
	}
	return result
}

// 107. 二叉树的层序遍历 II
func levelOrderBottom(root *TreeNode) [][]int {

	result := make([][]int, 0)
	levelStack := make([]*TreeNode, 0)
	if root == nil {
		return result
	}
	levelStack = append(levelStack, root)
	for len(levelStack) > 0 {
		l := len(levelStack)
		data := make([]int, 0)
		for i := 0; i < l; i++ {
			node := levelStack[i]
			data = append(data, node.Val)
			if node.Left != nil {
				levelStack = append(levelStack, node.Left)
			}
			if node.Right != nil {
				levelStack = append(levelStack, node.Right)
			}
		}
		levelStack = levelStack[l:]
		result = append(result, data)
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}
	return result
}

//103. 二叉树的锯齿形层序遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	i := 0
	for len(stack) > 0 {
		l := len(stack)
		data := make([]int, 0)
		for i := 0; i < l; i++ {
			node := stack[i]
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			data = append(data, node.Val)
		}

		if i%2 != 0 {
			reverseSlice(data)
		}
		result = append(result, data)
		stack = stack[l:]
		i++
	}
	return result
}

func reverseSlice(data []int) []int {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return data
}

// 98. 验证二叉搜索树
/*
二叉搜索树  的中序遍历是递增的
还可以用分治法。
*/
func isValidBST(root *TreeNode) bool {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		if node.Right != nil {
			node = node.Right
			for node != nil {
				stack = append(stack, node)
				node = node.Left
			}
		}
	}
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}

//DFS 分治法
func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isTrue, _, _ := doWorkIsValidBST2(root)
	if isTrue {
		return true
	}
	return false
}

func doWorkIsValidBST2(root *TreeNode) (bool, int, int) {
	if root.Left != nil && root.Right != nil {
		isLeft, leftMin, leftMax := doWorkIsValidBST2(root.Left)
		isRight, rightMin, rightMax := doWorkIsValidBST2(root.Right)
		if isLeft && isRight && root.Val > leftMax && root.Val < rightMin {
			return true, leftMin, rightMax
		} else {
			return false, 0, 0
		}
	}

	if root.Left != nil {
		isLeft, leftMin, leftMax := doWorkIsValidBST2(root.Left)
		if isLeft && root.Val > leftMax {
			return true, leftMin, root.Val
		} else {
			return false, 0, 0
		}
	}
	if root.Right != nil {
		isRight, rightMin, rightMax := doWorkIsValidBST2(root.Right)
		if isRight && root.Val < rightMin {
			return true, root.Val, rightMax
		} else {
			return false, 0, 0
		}
	}

	return true, root.Val, root.Val

}

// 701. 二叉搜索树中的插入操作
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	head := root
	for root != nil {
		if root.Val < val {
			if root.Right == nil {
				root.Right = &TreeNode{Val: val}
				return head
			}
			root = root.Right
		}
		if root.Val > val {
			if root.Left == nil {
				root.Left = &TreeNode{Val: val}
				return head
			}
			root = root.Left
		}
	}

	return head
}

// 83. 删除排序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	headR := head
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}

	}
	return headR

}

// 82. 删除排序链表中的重复元素 II
func deleteDuplicates2(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	dummy := &ListNode{Val: -1000}
	dummy.Next = head
	node := dummy
	dupMap := make(map[int]bool, 0)

	for dummy.Next != nil {
		if _, ok := dupMap[dummy.Next.Val]; ok {
			dummy.Next = dummy.Next.Next
			continue
		}

		if dummy.Next.Next != nil {
			if dummy.Next.Val == dummy.Next.Next.Val {
				dupMap[dummy.Next.Val] = true
				dummy.Next = dummy.Next.Next
				continue
			}
		}
		dummy = dummy.Next

	}

	return node.Next
}
