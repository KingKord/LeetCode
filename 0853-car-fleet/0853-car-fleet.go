func carFleet(target int, position []int, speed []int) int {
	var cars = make([][2]int, len(position))

	for i := range cars {
		cars[i] = [2]int{
			position[i],
			speed[i],
		}
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] > cars[j][0]
	})

	var max = 0

	var minTime = math.MaxFloat64
	for _, car := range cars {
		nowTime := float64((target)-car[0]) / float64(car[1])
		if max == 0 || nowTime > minTime {
			max++
			minTime = nowTime
		}
	}

	return max
}