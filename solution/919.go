package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "container/list"

type CBTInserter struct {
	root         *TreeNode
	hasNullChild *list.List
}

func Constructor_919(root *TreeNode) CBTInserter {
	newTree := CBTInserter{root: root}
	newTree.hasNullChild = list.New()
	newTree.hasNullChild.PushBack(root)
	for {
		temp := newTree.hasNullChild.Front().Value.(*TreeNode)
		c := 0
		if temp.Left != nil {
			newTree.hasNullChild.PushBack(temp.Left)
			c++
		}
		if temp.Right != nil {
			newTree.hasNullChild.PushBack(temp.Right)
			newTree.hasNullChild.Remove(newTree.hasNullChild.Front())
			c++
		}
		if c < 2 {
			break
		}
	}
	return newTree
}

func (this *CBTInserter) Insert(val int) int {
	queue := this.hasNullChild
	front := queue.Front().Value.(*TreeNode)
	ret := front.Val
	if front.Left == nil {
		front.Left = &TreeNode{Val: val, Left: nil, Right: nil}
		queue.PushBack(front.Left)
	} else {
		front.Right = &TreeNode{Val: val, Left: nil, Right: nil}
		queue.PushBack(front.Right)
		queue.Remove(queue.Front())
	}
	return ret
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
