# golang_good_nodes

Given a binary tree `root`, a node *X* in the tree is named **good** if in the path from root to *X* there are no nodes with a value *greater than* X.

Return the number of **good** nodes in the binary tree.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2020/04/02/test_sample_1.png](https://assets.leetcode.com/uploads/2020/04/02/test_sample_1.png)

```
Input: root = [3,1,4,3,null,1,5]
Output: 4
Explanation: Nodes in blue are good.
Root Node (3) is always a good node.
Node 4 -> (3,4) is the maximum value in the path starting from the root.
Node 5 -> (3,4,5) is the maximum value in the path
Node 3 -> (3,1,3) is the maximum value in the path.
```

**Example 2:**

![https://assets.leetcode.com/uploads/2020/04/02/test_sample_2.png](https://assets.leetcode.com/uploads/2020/04/02/test_sample_2.png)

```
Input: root = [3,3,null,4,2]
Output: 3
Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.
```

**Example 3:**

```
Input: root = [1]
Output: 1
Explanation: Root is considered as good.
```

**Constraints:**

- The number of nodes in the binary tree is in the range $`[1, 10^5]`$.
- Each node's value is between $`[-10^4, 10^4]`$.

## 解析

題目給一個二元樹的根結點 root 

定義一個結點 X 是 good 代表從 root 到 X 的值都不大於 X.Val

要累計這個二元樹是 good 的結點個數

 關鍵點在於透過 DFS 還有累計當下最大的值來做符合條件的檢驗

![](https://i.imgur.com/jVNFfo0.png)

## 程式碼

```go
package sol

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	max := math.MinInt32
	CountGoodNodes(root, max, &count)
	return count
}

func CountGoodNodes(root *TreeNode, max int, count *int) {
	if root == nil {
		return
	}
	curMax := max
	if curMax <= root.Val {
		*count += 1
		curMax = root.Val
	}
	CountGoodNodes(root.Left, curMax, count)
	CountGoodNodes(root.Right, curMax, count)
}

```
## 困難點

1. 理解結點是 good 的條件
2. 理解如何做 DFS

## Solve Point

- [x]  Understand what problem to solve
- [x]  Analysis Complexity