package solution

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if destination < start {
		start, destination = destination, start
	}
	forward := 0
	sum := 0
	for i, v := range distance {
		if i >= start && i < destination {
			forward += v
		}
		sum += v
	}
	backward := sum - forward
	if forward > backward {
		return backward
	}
	return forward
}
