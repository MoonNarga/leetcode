package solution

import "container/list"

func deepestLeavesSum(root *TreeNode) int {
	queue := list.New()
	queue.PushBack(root)
	currentNum, childNum, sumLevel, sumFinal := 1, 0, 0, 0
	var head *TreeNode
	for currentNum != 0 {
		head = queue.Front().Value.(*TreeNode)
		if head.Left != nil {
			queue.PushBack(head.Left)
			childNum++
		}
		if head.Right != nil {
			queue.PushBack(head.Right)
			childNum++
		}
		sumLevel += head.Val
		queue.Remove(queue.Front())
		currentNum--
		if currentNum == 0 {
			sumFinal, currentNum, childNum, sumLevel = sumLevel, childNum, 0, 0
		}
	}
	return sumFinal
}
