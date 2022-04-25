package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	bCount := 0
	for _, birdsPerDay := range birdsPerDay {
		if birdsPerDay < 0 {
			continue
		}
		bCount += birdsPerDay
	}
	return bCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	birdsPerWeek := map[int]int{}
	var weekNr int

	for i := 0; i < len(birdsPerDay); i++ {
		if i == 0 || i%7 == 0 {
			weekNr++
			birdsPerWeek[weekNr] = 0
		}

		birdsPerWeek[weekNr] += birdsPerDay[i]
	}
	return birdsPerWeek[week]
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}
