package solution

import "strings"

type pathStack []string

func (s pathStack) Len() int { return len(s) }

func (s pathStack) Cap() int { return cap(s) }

func (s *pathStack) Push(p string) { *s = append(*s, p) }

func (s *pathStack) Pop() string {
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value
}

func (s pathStack) Top() string {
	return s[len(s)-1]
}

func (s pathStack) IsEmpty() bool {
	return len(s) == 0
}

func simplifyPath(path string) string {
	pathslice := strings.Split(path, "/")
	pathstack := make(pathStack, 0)
	for _, v := range pathslice {
		if v == "" {
			continue
		} else if v == ".." {
			if pathstack.IsEmpty() {
				continue
			} else {
				pathstack.Pop()
			}
		} else if v == "." {
			continue
		} else {
			pathstack.Push(v)
		}
	}
	res := ""
	for _, v := range pathstack {
		res += "/"
		res += v
	}
	if res == "" {
		res += "/"
	}
	return res
}
