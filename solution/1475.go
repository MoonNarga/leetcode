package solution

func finalPrices(prices []int) []int {
	pr := make([]int, 0)
	indices := make([]int, 0)
	for i, v := range prices {
		for len(pr) > 0 && v <= pr[len(pr)-1] {
			prices[indices[len(indices)-1]] = pr[len(pr)-1] - v
			pr = pr[:len(pr)-1]
			indices = indices[:len(indices)-1]
		}
		pr = append(pr, v)
		indices = append(indices, i)
	}
	for i := 0; i < len(pr); i++ {
		prices[indices[i]] = pr[i]
	}
	return prices
}
