package solution

import (
	"sort"
)

type L struct {
	x      int
	top    int
	bottom int
}

type Lslice []L

func (lslice *Lslice) Less(i, j int) bool {
	if (*lslice)[i].x != (*lslice)[j].x {
		return (*lslice)[i].x < (*lslice)[j].x
	} else {
		return (*lslice)[i].top > (*lslice)[j].top
	}
}

func (lslice *Lslice) Swap(i, j int) {
	(*lslice)[i], (*lslice)[j] = (*lslice)[j], (*lslice)[i]
}

func (lslice *Lslice) Len() int {
	return len((*lslice))
}

func isRectangleCover(rectangles [][]int) bool {
	if len(rectangles) == 1 {
		return true
	}
	var LsetLeft Lslice
	var LsetRight Lslice
	for _, v := range rectangles {
		LsetLeft = append(LsetLeft, L{x: v[0], top: v[3], bottom: v[1]})
		LsetRight = append(LsetRight, L{x: v[2], top: v[3], bottom: v[1]})
	}
	sort.Sort(&LsetLeft)
	sort.Sort(&LsetRight)
	var fixLsetLeft Lslice
	var fixLsetRight Lslice
	tl := L{}
	tl = LsetLeft[0]
	for i, v := range LsetLeft {
		if i == 0 {
			continue
		}
		if tl.x != v.x {
			fixLsetLeft = append(fixLsetLeft, tl)
			tl = v
		} else {
			if tl.bottom < v.top {
				return false
			} else if tl.bottom == v.top {
				tl.bottom = v.bottom
			} else {
				fixLsetLeft = append(fixLsetLeft, tl)
				tl = v
			}
		}
		if i == len(LsetLeft)-1 {
			fixLsetLeft = append(fixLsetLeft, tl)
		}
	}
	tl = LsetRight[0]
	for i, v := range LsetRight {
		if i == 0 {
			continue
		}
		if tl.x != v.x {
			fixLsetRight = append(fixLsetRight, tl)
			tl = v
		} else {
			if tl.bottom < v.top {
				return false
			} else if tl.bottom == v.top {
				tl.bottom = v.bottom
			} else {
				fixLsetRight = append(fixLsetRight, tl)
				tl = v
			}

		}
		if i == len(LsetRight)-1 {
			fixLsetRight = append(fixLsetRight, tl)
		}
	}
	if fixLsetLeft[0].top != fixLsetRight[len(fixLsetRight)-1].top || fixLsetLeft[0].bottom != fixLsetRight[len(fixLsetRight)-1].bottom {
		return false
	}
	top := fixLsetLeft[0].top
	bottom := fixLsetLeft[0].bottom
	for i, v := range fixLsetLeft {
		if i == 0 {
			continue
		}
		if v.top > top || v.bottom < bottom {
			return false
		}
		if v != fixLsetRight[i-1] {
			return false
		}
	}
	return true
}
