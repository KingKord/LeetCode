
func minEatingSpeed(piles []int, h int) int {
	sum := 0
	for i := 0; i < len(piles); i++ {
		sum += piles[i]
	}

	var (
		L = sum / h
		R = sum
	)
    if L == 0 {
		L++
	}
	for L < R {
		m := (L + R) / 2

		now := 0
		for i := 0; i < len(piles); i++ {
			nowHours := piles[i] / m
			if piles[i]%m != 0 {
				nowHours++
			}
			now += nowHours
		}

		if now <= h {
			R = m
		} else {
			L = m + 1
		}
	}
	return L
}
