package solution

func subarraySum(nums []int, k int) int {
	sumCount := make(map[int]int, 0)
	sum := 0
	res := 0
	sumCount[0] = 1
	for _, v := range nums {
		sum += v
		if cnt, ok := sumCount[sum-k]; ok {
			res += cnt
		}
		sumCount[sum]++
	}
	return res
}
