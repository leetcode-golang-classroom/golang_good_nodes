package sol

import (
	"strconv"
	"testing"
)

func CreateBinaryTree(input *[]string) *TreeNode {
	var tree *TreeNode
	arr := *input
	result := make([]*TreeNode, len(arr))
	for idx, val := range arr {
		if val != "null" {
			num, _ := strconv.Atoi(val)
			result[idx] = &TreeNode{Val: num}
		}
	}
	tree = result[0]
	for idx, node := range result {
		if 2*idx+1 < len(result) {
			node.Left = result[2*idx+1]
		}
		if 2*idx+2 < len(result) {
			node.Right = result[2*idx+2]
		}
	}
	return tree
}
func Test_goodNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "root = [3,1,4,3,null,1,5]",
			args: args{root: CreateBinaryTree(&[]string{"3", "1", "4", "3", "null", "1", "5"})},
			want: 4,
		},
		{
			name: "root = [3,3,null,4,2]",
			args: args{root: CreateBinaryTree(&[]string{"3", "3", "null", "4", "2"})},
			want: 3,
		},
		{
			name: "root = [1]",
			args: args{root: CreateBinaryTree(&[]string{"1"})},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goodNodes(tt.args.root); got != tt.want {
				t.Errorf("goodNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
