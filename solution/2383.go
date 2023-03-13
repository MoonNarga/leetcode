package solution

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	res, temp := 0, 0
	for i := 0; i < len(energy); i++ {
		temp += energy[i]
		if initialExperience <= experience[i] {
			res += experience[i] - initialExperience + 1
			initialExperience += experience[i] - initialExperience + 1
		}
		initialExperience += experience[i]
	}
	if initialEnergy > temp {
		return res
	}
	return res + (temp - initialEnergy + 1)
}
