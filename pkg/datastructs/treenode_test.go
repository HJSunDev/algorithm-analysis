package datastructs

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	t.Run("空参数", func(t *testing.T) {
		root := BuildTree()
		if root != nil {
			t.Errorf("期望 nil, 得到 %v", root)
		}
	})

	t.Run("NULL根节点", func(t *testing.T) {
		// 修复验证：BuildTree(NULL) 应返回 nil
		root := BuildTree(NULL)
		if root != nil {
			t.Errorf("BuildTree(NULL) 应返回 nil, 得到 Val=%d", root.Val)
		}
	})

	t.Run("单节点", func(t *testing.T) {
		root := BuildTree(1)
		if root == nil || root.Val != 1 {
			t.Errorf("期望 Val=1, 得到 %v", root)
		}
		if root.Left != nil || root.Right != nil {
			t.Error("单节点不应有子节点")
		}
	})

	t.Run("完整二叉树", func(t *testing.T) {
		// 构建:    3
		//         / \
		//        9  20
		//           / \
		//          15  7
		root := BuildTree(3, 9, 20, NULL, NULL, 15, 7)
		if root.Val != 3 {
			t.Errorf("根节点 Val 期望 3, 得到 %d", root.Val)
		}
		if root.Left.Val != 9 {
			t.Errorf("左子节点 Val 期望 9, 得到 %d", root.Left.Val)
		}
		if root.Right.Val != 20 {
			t.Errorf("右子节点 Val 期望 20, 得到 %d", root.Right.Val)
		}
		if root.Right.Left.Val != 15 || root.Right.Right.Val != 7 {
			t.Error("第三层节点值不正确")
		}
	})
}

func TestTreeNode_Traversals(t *testing.T) {
	// 构建:    1
	//         / \
	//        2   3
	//       / \
	//      4   5
	root := BuildTree(1, 2, 3, 4, 5)

	t.Run("中序遍历", func(t *testing.T) {
		got := root.Inorder()
		want := []int{4, 2, 5, 1, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Inorder() = %v, want %v", got, want)
		}
	})

	t.Run("前序遍历", func(t *testing.T) {
		got := root.Preorder()
		want := []int{1, 2, 4, 5, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Preorder() = %v, want %v", got, want)
		}
	})

	t.Run("后序遍历", func(t *testing.T) {
		got := root.Postorder()
		want := []int{4, 5, 2, 3, 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Postorder() = %v, want %v", got, want)
		}
	})

	t.Run("层序遍历", func(t *testing.T) {
		got := root.LevelOrder()
		want := [][]int{{1}, {2, 3}, {4, 5}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("LevelOrder() = %v, want %v", got, want)
		}
	})

	t.Run("nil树遍历", func(t *testing.T) {
		var nilTree *TreeNode
		if nilTree.Inorder() != nil {
			t.Error("nil 树的中序遍历应返回 nil")
		}
		if nilTree.LevelOrder() != nil {
			t.Error("nil 树的层序遍历应返回 nil")
		}
	})
}

func TestTreeNode_MaxDepth(t *testing.T) {
	tests := []struct {
		name string
		tree *TreeNode
		want int
	}{
		{name: "nil树", tree: nil, want: 0},
		{name: "单节点", tree: BuildTree(1), want: 1},
		{name: "三层树", tree: BuildTree(3, 9, 20, NULL, NULL, 15, 7), want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.MaxDepth(); got != tt.want {
				t.Errorf("MaxDepth() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestTreeNode_Equal(t *testing.T) {
	tests := []struct {
		name string
		a, b *TreeNode
		want bool
	}{
		{name: "两个nil", a: nil, b: nil, want: true},
		{name: "一个nil", a: BuildTree(1), b: nil, want: false},
		{name: "结构相同", a: BuildTree(1, 2, 3), b: BuildTree(1, 2, 3), want: true},
		{name: "值不同", a: BuildTree(1, 2, 3), b: BuildTree(1, 2, 4), want: false},
		{name: "结构不同", a: BuildTree(1, 2), b: BuildTree(1, NULL, 2), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Equal(tt.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_String(t *testing.T) {
	t.Run("nil树", func(t *testing.T) {
		var nilTree *TreeNode
		if got := nilTree.String(); got != "[]" {
			t.Errorf("nil 树 String() = %q, want \"[]\"", got)
		}
	})

	t.Run("正常树", func(t *testing.T) {
		root := BuildTree(1, 2, 3)
		got := root.String()
		if len(got) == 0 {
			t.Error("String() 返回空字符串")
		}
	})
}
