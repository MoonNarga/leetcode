package solution

func maxDepth_1614(s string) int {
	top := 0
	maxtop := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			top++
		}
		if s[i] == ')' {
			top--
		}
		if top > maxtop {
			maxtop = top
		}
	}
	return maxtop
}
