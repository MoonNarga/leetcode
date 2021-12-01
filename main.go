package main

import (
	"fmt"
)

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	pmap := make(map[byte]int)
	smap := make(map[byte]int)
	if len(s) < len(p) {
		return nil
	}
	for i := 0; i < len(p); i++ {
		if _, ok := pmap[p[i]]; ok {
			pmap[p[i]]++
		} else {
			pmap[p[i]] = 1
		}
	}
	lenpmap := len(pmap)
	flag := false
	for i, j := 0, 0; i < len(s); {
		if flag {
			flag = false
			for k := range smap {
				delete(smap, k)
			}
		}
		for j = i; j < len(s); j++ {
			vp, okp := pmap[s[j]]
			if !okp {
				flag = true
				break
			}
			vs, oks := smap[s[i]]
			if !oks {
				smap[s[i]] = 0
			} else if vs >= vp {
				break
			} else {
				smap[s[i]]++
			}
			if len(smap) == lenpmap {
				res = append(res, i)
			}
		}
	}
	return res
}

func main() {
	s, p := "cbaebabacd", "abc"
	fmt.Println(findAnagrams(s, p))
}
