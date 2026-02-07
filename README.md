# Algorithm Analysis

> Go 语言算法与数据结构学习项目

## 项目简介

使用 Go 语言系统性地学习和练习算法与数据结构。通过 AI 辅助实现题目分析、代码验证、总结归档的完整学习闭环。

## 项目结构

```
algorithm-analysis/
├── .cursor/rules/           # AI 行为规则
├── pkg/
│   └── datastructs/         # 公共数据结构 (ListNode, TreeNode 等)
├── solutions/               # 题解目录
│   ├── README.md            #   分桶索引入口
│   └── XXXX-XXXX/           #   分桶目录（每 100 题一桶）
│       ├── README.md        #     本桶题目索引
│       └── XXXX-slug/       #     具体题目
│           ├── README.md    #       题目分析文档
│           ├── solution.go  #       解题代码
│           └── solution_test.go
├── topics/                  # 知识点分类导航
│   ├── README.md            #   知识点总览
│   └── {topic}.md           #   各知识点题目索引（按需生成）
├── go.mod
└── README.md                # 本文件
```

## 导航

| 维度     | 入口                           | 说明                           |
| -------- | ------------------------------ | ------------------------------ |
| 按知识点 | [topics/](topics/README.md)       | 数组、链表、动态规划等分类索引 |
| 按题号   | [solutions/](solutions/README.md) | 按题号分桶，每桶独立索引       |

## 公共数据结构

`pkg/datastructs` 提供解题常用数据结构：

```go
import ds "algorithm-analysis/pkg/datastructs"
```

| 类型         | 构建方法                               | 说明                        |
| ------------ | -------------------------------------- | --------------------------- |
| `ListNode` | `ds.BuildList(1, 2, 3)`              | 单链表                      |
| `ListNode` | `ds.BuildCycleList([]int{1,2,3}, 1)` | 带环链表，参数为环入口索引  |
| `TreeNode` | `ds.BuildTree(1, ds.NULL, 2, 3)`     | 二叉树，`NULL` 表示空节点 |

`TreeNode` 内置遍历方法：`Inorder()`, `Preorder()`, `Postorder()`, `LevelOrder()`

## 常用命令

```bash
# 运行全部测试
go test ./solutions/...

# 运行特定题目测试
go test -v ./solutions/0000-0099/0001-two-sum/

# 静态检查
go vet ./solutions/...
```

## 进度

- 已完成: 0
- Easy: 0 | Medium: 0 | Hard: 0
