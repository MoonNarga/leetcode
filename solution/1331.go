package solution

func sortTwoArr(data, index []int, L, R int) {
	if L >= R {
		return
	}
	data[L], data[(L+R)/2] = data[(L+R)/2], data[L]
	index[L], index[(L+R)/2] = index[(L+R)/2], index[L]
	valWatcher := data[L]
	indWatcher := index[L]
	var i, j int
	for i, j = L, R; i != j; {
		for i != j {
			if data[j] < valWatcher {
				break
			} else {
				j--
			}
		}
		data[i], index[i] = data[j], index[j]
		for i != j {
			if data[i] > valWatcher {
				break
			} else {
				i++
			}
		}
		data[j], index[j] = data[i], index[i]
	}
	data[i], index[i] = valWatcher, indWatcher
	sortTwoArr(data, index, L, i-1)
	sortTwoArr(data, index, i+1, R)
}

func arrayRankTransform(arr []int) []int {
	l := len(arr)
	ind := make([]int, l)
	res := make([]int, l)
	if l == 0 {
		return res
	}
	for i := 0; i < l; i++ {
		ind[i] = i
	}
	sortTwoArr(arr, ind, 0, l-1)
	rank := 1
	temp := arr[0]
	for i := 0; i < l; i++ {
		if arr[i] != temp {
			rank++
			temp = arr[i]
		}
		res[ind[i]] = rank
	}
	return res
}
