package solution

import "container/list"

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queNode := list.New()
	quePos := list.New()
	queNode.PushBack(root)
	quePos.PushBack(1)
	numLevel := 1
	numTemp := 0
	maxWidth := 1
	L, R := 0, 0
	for numLevel != 0 {
		if L == 0 {
			L = quePos.Front().Value.(int)
			R = L
		} else if quePos.Front().Value.(int) > R {
			R = quePos.Front().Value.(int)
		}
		if queNode.Front().Value.(*TreeNode).Left != nil {
			queNode.PushBack(queNode.Front().Value.(*TreeNode).Left)
			quePos.PushBack(quePos.Front().Value.(int)*2 - 1)
			numTemp++
		}
		if queNode.Front().Value.(*TreeNode).Right != nil {
			queNode.PushBack(queNode.Front().Value.(*TreeNode).Right)
			quePos.PushBack(quePos.Front().Value.(int) * 2)
			numTemp++
		}
		queNode.Remove(queNode.Front())
		quePos.Remove(quePos.Front())
		numLevel--
		if numLevel == 0 {
			if R-L+1 > maxWidth {
				maxWidth = R - L + 1
			}
			numLevel = numTemp
			numTemp = 0
			L, R = 0, 0
		}
	}
	return maxWidth
}
