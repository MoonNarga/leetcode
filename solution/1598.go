package solution

func minOperations(logs []string) int {
	depth := 0
	for _, v := range logs {
		switch v {
		case "./":
			{
				break
			}
		case "../":
			{
				if depth > 0 {
					depth--
				}
			}
		default:
			{
				depth++
			}
		}
	}
	return depth
}
