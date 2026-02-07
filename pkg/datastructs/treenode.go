package datastructs

// TreeNode 二叉树节点，用于树类题目
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NULL 表示二叉树中的空节点，用于 BuildTree 的参数。
// 值为 int 最小值，不会与正常节点值冲突。
const NULL = -1 << 31

// BuildTree 从层序遍历序列构建二叉树。
// 使用 NULL 常量表示空节点。
//
// 示例:
//
//	root := BuildTree(3, 9, 20, NULL, NULL, 15, 7)
//	构建如下二叉树:
//	      3
//	     / \
//	    9  20
//	      /  \
//	     15   7
func BuildTree(vals ...int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}

	// 创建根节点
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		// 构建左子节点
		if i < len(vals) {
			if vals[i] != NULL {
				node.Left = &TreeNode{Val: vals[i]}
				queue = append(queue, node.Left)
			}
			i++
		}

		// 构建右子节点
		if i < len(vals) {
			if vals[i] != NULL {
				node.Right = &TreeNode{Val: vals[i]}
				queue = append(queue, node.Right)
			}
			i++
		}
	}

	return root
}

// Inorder 返回中序遍历结果（左-根-右）
func (t *TreeNode) Inorder() []int {
	if t == nil {
		return nil
	}
	var res []int
	res = append(res, t.Left.Inorder()...)
	res = append(res, t.Val)
	res = append(res, t.Right.Inorder()...)
	return res
}

// Preorder 返回前序遍历结果（根-左-右）
func (t *TreeNode) Preorder() []int {
	if t == nil {
		return nil
	}
	var res []int
	res = append(res, t.Val)
	res = append(res, t.Left.Preorder()...)
	res = append(res, t.Right.Preorder()...)
	return res
}

// Postorder 返回后序遍历结果（左-右-根）
func (t *TreeNode) Postorder() []int {
	if t == nil {
		return nil
	}
	var res []int
	res = append(res, t.Left.Postorder()...)
	res = append(res, t.Right.Postorder()...)
	res = append(res, t.Val)
	return res
}

// LevelOrder 返回层序遍历结果，每层为一个切片
func (t *TreeNode) LevelOrder() [][]int {
	if t == nil {
		return nil
	}
	var res [][]int
	queue := []*TreeNode{t}
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

// MaxDepth 返回二叉树的最大深度
func (t *TreeNode) MaxDepth() int {
	if t == nil {
		return 0
	}
	left := t.Left.MaxDepth()
	right := t.Right.MaxDepth()
	if left > right {
		return left + 1
	}
	return right + 1
}
