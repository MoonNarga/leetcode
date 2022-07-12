package solution

func slowestKey(releaseTimes []int, keysPressed string) byte {
	alphatemp := keysPressed[0]
	maxlast := releaseTimes[0]
	for i := 1; i < len(releaseTimes); i++ {
		if releaseTimes[i]-releaseTimes[i-1] > maxlast || releaseTimes[i]-releaseTimes[i-1] == maxlast && keysPressed[i] > alphatemp {
			alphatemp = keysPressed[i]
			maxlast = releaseTimes[i] - releaseTimes[i-1]
		}
	}
	return alphatemp
}
