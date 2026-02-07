package datastructs

import (
	"reflect"
	"testing"
)

func TestBuildList(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{name: "空参数", vals: nil, want: nil},
		{name: "单节点", vals: []int{1}, want: []int{1}},
		{name: "多节点", vals: []int{1, 2, 3}, want: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := BuildList(tt.vals...)
			got := head.Slice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildList(%v).Slice() = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}

func TestBuildCycleList(t *testing.T) {
	t.Run("空输入", func(t *testing.T) {
		head := BuildCycleList(nil, -1)
		if head != nil {
			t.Errorf("期望 nil, 得到 %v", head)
		}
	})

	t.Run("无环", func(t *testing.T) {
		head := BuildCycleList([]int{1, 2, 3}, -1)
		got := head.Slice()
		if !reflect.DeepEqual(got, []int{1, 2, 3}) {
			t.Errorf("期望 [1 2 3], 得到 %v", got)
		}
	})

	t.Run("有环-尾接头", func(t *testing.T) {
		head := BuildCycleList([]int{1, 2, 3}, 0)
		// 验证环存在：快慢指针
		slow, fast := head, head
		hasCycle := false
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
			if slow == fast {
				hasCycle = true
				break
			}
		}
		if !hasCycle {
			t.Error("期望有环，但未检测到环")
		}
	})
}

func TestListNode_Equal(t *testing.T) {
	tests := []struct {
		name string
		a, b *ListNode
		want bool
	}{
		{name: "两个nil", a: nil, b: nil, want: true},
		{name: "一个nil", a: BuildList(1), b: nil, want: false},
		{name: "值相等", a: BuildList(1, 2, 3), b: BuildList(1, 2, 3), want: true},
		{name: "值不等", a: BuildList(1, 2), b: BuildList(1, 3), want: false},
		{name: "长度不等", a: BuildList(1, 2), b: BuildList(1, 2, 3), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Equal(tt.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_String(t *testing.T) {
	t.Run("普通链表", func(t *testing.T) {
		head := BuildList(1, 2, 3)
		got := head.String()
		want := "[1 2 3]"
		if got != want {
			t.Errorf("String() = %q, want %q", got, want)
		}
	})

	t.Run("环形链表不死循环", func(t *testing.T) {
		// 关键：环形链表调用 String() 不应卡死，应截断
		head := BuildCycleList([]int{1, 2, 3}, 0)
		got := head.String()
		if len(got) == 0 {
			t.Error("String() 返回空字符串")
		}
		// 验证截断提示存在
		if len(head.Slice()) < maxTraverse {
			t.Error("环形链表应触发截断保护")
		}
	})
}
