package solution

import "math"

func asteroidCollision(asteroids []int) []int {
	res := make([]int, len(asteroids))
	leng := 0
	lenAst := len(asteroids)
	for i := 0; i < lenAst; {
		if leng == 0 {
			res[leng] = asteroids[i]
			i++
			leng++
			continue
		}
		if res[leng-1] > 0 && asteroids[i] < 0 {
			diff := math.Abs(float64(res[leng-1])) - math.Abs(float64(asteroids[i]))
			switch {
			case diff == 0:
				{
					leng--
					i++
				}
			case diff > 0:
				{
					i++
				}
			case diff < 0:
				{
					leng--
				}
			}
		} else {
			res[leng] = asteroids[i]
			i++
			leng++
		}
	}
	return res[:leng]
}
