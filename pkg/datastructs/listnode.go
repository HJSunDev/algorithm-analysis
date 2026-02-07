// Package datastructs 提供力扣常用的公共数据结构及其辅助方法。
// 导入路径: algorithm-analysis/pkg/datastructs
package datastructs

import "fmt"

// ListNode 单链表节点，用于链表类题目
type ListNode struct {
	Val  int
	Next *ListNode
}

// BuildList 从整数序列构建单链表，返回头节点。
// 传入空参数时返回 nil。
//
// 示例:
//
//	head := BuildList(1, 2, 3) // 1→2→3
func BuildList(vals ...int) *ListNode {
	// 使用哨兵节点简化头节点处理
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// BuildCycleList 构建带环的链表，cyclePos 指定环的入口索引（从 0 开始）。
// cyclePos 为 -1 时表示无环。
//
// 示例:
//
//	head := BuildCycleList([]int{3, 2, 0, -4}, 1) // 尾部连接到索引 1 的节点
func BuildCycleList(vals []int, cyclePos int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	// 构建链表
	dummy := &ListNode{}
	cur := dummy
	// nodes 保存所有节点引用，用于设置环
	nodes := make([]*ListNode, len(vals))
	for i, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
		nodes[i] = cur
	}

	// 设置环：将尾节点的 Next 指向 cyclePos 位置的节点
	if cyclePos >= 0 && cyclePos < len(nodes) {
		nodes[len(nodes)-1].Next = nodes[cyclePos]
	}

	return dummy.Next
}

// maxTraverse 是 Slice/String 遍历的节点上限，防止环形链表导致死循环
const maxTraverse = 1000

// Slice 将链表转换为整数切片，便于测试中进行断言比较。
// nil 链表返回 nil 切片。超过 1000 个节点时截断，防止环形链表死循环。
func (l *ListNode) Slice() []int {
	var res []int
	for p := l; p != nil; p = p.Next {
		res = append(res, p.Val)
		if len(res) >= maxTraverse {
			break
		}
	}
	return res
}

// String 格式化输出链表，例: [1 2 3]
// 环形链表会截断并标注提示
func (l *ListNode) String() string {
	s := l.Slice()
	if len(s) >= maxTraverse {
		return fmt.Sprintf("%v...(截断: 可能存在环)", s)
	}
	return fmt.Sprintf("%v", s)
}

// Equal 判断两个链表是否值相等
func (l *ListNode) Equal(other *ListNode) bool {
	a, b := l, other
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	// 两个链表必须同时到达末尾
	return a == nil && b == nil
}
