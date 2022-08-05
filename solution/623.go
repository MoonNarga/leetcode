package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		newRoot := &TreeNode{Val: val, Left: root, Right: nil}
		return newRoot
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root
	levelNum := 1
	tempLevelNum := 0
	currentLevel := 1
	for ; currentLevel != depth-1; currentLevel++ {
		for ; levelNum != 0; levelNum-- {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
				tempLevelNum++
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
				tempLevelNum++
			}
			queue = queue[1:]
		}
		levelNum, tempLevelNum = tempLevelNum, 0
	}
	for _, v := range queue {
		v.Left = &TreeNode{Val: val, Left: v.Left, Right: nil}
		v.Right = &TreeNode{Val: val, Left: nil, Right: v.Right}
	}
	return root
}
