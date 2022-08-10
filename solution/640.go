package solution

import (
	"fmt"
	"strings"
)

func parseNum(buf []byte) (num, len int, va bool) {
	num = 0
	neg := false
	hasDigit := false
	pos := 0
	va = false
	if buf[0] == '-' {
		neg = true
		pos++
	}
	if buf[0] == '+' {
		pos++
	}
	for {
		if buf[pos] <= '9' && buf[pos] >= '0' {
			num = num*10 + int(buf[pos]-'0')
			pos++
			hasDigit = true
		} else {
			if buf[pos] == 'x' {
				va = true
				pos++
			}
			break
		}
	}
	if num == 0 {
		if hasDigit {
			num = 0
		} else {
			num = 1
		}
	}
	if neg {
		num = -num
	}
	len = pos
	return
}

func solveEquation(equation string) string {
	exprs := strings.Split(equation, "=")
	exprLeft := []byte(exprs[0])
	exprRight := []byte(exprs[1])
	exprLeft = append(exprLeft, ' ')
	exprRight = append(exprRight, ' ')
	numLeft, numRight, xLeft, xRight := 0, 0, 0, 0
	lenLeft, lenRight := len(exprs[0]), len(exprs[1])
	for i := 0; i < lenLeft; {
		tNum, tLen, tVa := parseNum(exprLeft[i:])
		i += tLen
		if tVa {
			xLeft += tNum
		} else {
			numLeft += tNum
		}
	}
	for i := 0; i < lenRight; {
		tNum, tLen, tVa := parseNum(exprRight[i:])
		i += tLen
		if tVa {
			xRight += tNum
		} else {
			numRight += tNum
		}
	}
	xLeft -= xRight
	numRight -= numLeft
	if xLeft == 0 {
		if numRight == 0 {
			return "Infinite solutions"
		} else {
			return "No solution"
		}
	} else {
		return "x=" + fmt.Sprint(numRight/xLeft)
	}
}
