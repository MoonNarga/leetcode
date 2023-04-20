package solution

type bucket struct {
	min  int
	max  int
	used bool
}

func maxWidthOfVerticalArea(points [][]int) int {
	numMax, numMin := points[0][0], points[0][0]
	for i := 0; i < len(points); i++ {
		if points[i][0] > numMax {
			numMax = points[i][0]
		}
		if points[i][0] < numMin {
			numMin = points[i][0]
		}
	}
	valueRange := numMax - numMin
	bucketSize := valueRange/len(points) + 1
	buckets := make([]bucket, bucketSize)
	for i := 0; i < len(points); i++ {
		index := (points[i][0] - numMin) / bucketSize
		buckets[index].used = true
		if points[i][0] < buckets[index].min {
			buckets[index].min = points[i][0]
		}
		if points[i][0] > buckets[index].max {
			buckets[index].max = points[i][0]
		}
	}
	maxWidth := 0
	for i := 0; i < len(buckets); i++ {
		if buckets[i].used {
			if i > 0 {
				if buckets[i-1].used {
					if buckets[i].min-buckets[i-1].max > maxWidth {
						maxWidth = buckets[i].min - buckets[i-1].max
					}
				}
			}
		}
	}
	return maxWidth
}
