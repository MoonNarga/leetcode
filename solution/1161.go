package solution

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 1
	}
	nodeNumOfLevel := 1
	tempNodeNum := 0
	maxSum := root.Val
	maxLevel := 1
	currentLevel := 1
	tempSum := 0
	levelList := make([]*TreeNode, 0)
	levelList = append(levelList, root)
	for ; nodeNumOfLevel != 0; nodeNumOfLevel-- {
		tempSum += levelList[0].Val
		if levelList[0].Left != nil {
			levelList = append(levelList, levelList[0].Left)
			tempNodeNum++
		}
		if levelList[0].Right != nil {
			levelList = append(levelList, levelList[0].Right)
			tempNodeNum++
		}
		levelList = levelList[1:]
		if nodeNumOfLevel == 1 {
			if maxSum < tempSum {
				maxSum = tempSum
				maxLevel = currentLevel
			}
			if tempNodeNum != 0 {
				nodeNumOfLevel, tempNodeNum = tempNodeNum+1, 0
				tempSum = 0
				currentLevel++
			}
		}
	}
	return maxLevel
}
