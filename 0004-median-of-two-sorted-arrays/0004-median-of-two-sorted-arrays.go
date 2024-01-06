
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	var (
		len1        = len(nums1)
		len2        = len(nums2)
		total       = len1 + len2
		half        = total / 2
		firstRight  = len1 - 1
		secondLeft  = 0
		secondRight = len2 - 1
	)

	if len2 == 1 && len1 != 0 {
		if len1 == 1 {
			return float64(nums1[0]+nums2[0]) / 2
		}
		if (len1+len2)%2 == 0 {
			if nums1[half] >= nums2[0] {
				return float64(Max(nums1[half-2], nums2[0])+nums1[half-1]) / 2
			} else {
				return float64(Min(nums1[half], nums2[0])+nums1[half-1]) / 2
			}
		}
		if nums1[half-1] <= nums2[0] && nums2[0] <= nums1[half] {
			return float64(nums2[0])
		} else if nums2[0] <= nums1[half-1] {
			return float64(nums1[half-1])
		} else {
			return float64(nums1[half])
		}
	}

	if len2 == 0 {
		if (len1+len2)%2 == 0 {
			return float64(nums1[half-1]+nums1[half]) / 2
		}
		return float64(nums1[half])
	}

	if nums1[0] >= nums2[len2-1] {
		if half == len2 {
			if (len1+len2)%2 == 0 {
				return float64(nums1[0]+nums2[len2-1]) / 2
			}
			return float64(nums1[0])
		}
		if (len1+len2)%2 == 0 {
			return float64(nums1[len1-half-1]+nums1[len1-half]) / 2
		}
		return float64(nums1[len1-half-1])
	} else if nums2[0] > nums1[len1-1] {
		if half == len2 {
			if (len1+len2)%2 == 0 {
				return float64(nums1[len1-1]+nums2[0]) / 2
			}
			return float64(nums1[len1-1])
		}
		if (len1+len2)%2 == 0 {
			return float64(nums1[half-1]+nums1[half]) / 2
		}
		return float64(nums1[half])
	}
	for {
		middleSecond := (secondLeft + secondRight) / 2

		middleFirst := half - 2 - middleSecond

		if middleFirst == -1 {
			firstRight = -1
			break
		}
		if middleSecond+1 >= len2 {
			if (len1+len2)%2 == 0 {
				return float64(Max(nums1[middleFirst], nums2[middleSecond])+nums1[middleFirst+1]) / 2
			}
			return float64(nums1[middleFirst+1])
		}
		if nums1[middleFirst+1] >= nums2[middleSecond] && nums2[middleSecond+1] >= nums1[middleFirst] {
			firstRight = middleFirst
			secondRight = middleSecond
			break
		} else if nums2[middleSecond] >= nums1[middleFirst+1] {
			secondRight = middleSecond
		} else if nums1[middleFirst] >= nums2[middleSecond+1] {
			secondLeft = middleSecond + 1
		}

		if secondRight == 0 {
			if (len2+len1)%2 == 0 {
				return (float64(nums1[half-1]) + float64(Min(nums1[half], nums2[0]))) / 2
			}
			return float64(Min(nums1[half], nums2[0]))
		}
	}

	fst := 0.0
	snd := 0.0
	if firstRight == -1 {
		fst = float64(nums2[secondRight])
		snd = float64(nums1[0])
	} else {
		fst = float64(Min(nums1[firstRight+1], nums2[secondRight+1]))
		snd = float64(Max(nums1[firstRight], nums2[secondRight]))
	}
	if (len2+len1)%2 == 0 {
		return (fst + snd) / 2
	}
	if firstRight+1 >= len1 {
		return float64(nums2[0])
	} else if secondRight+1 >= len2 {
		return float64(nums1[0])
	}
	fst = float64(Min(nums1[firstRight+1], nums2[secondRight+1]))
	return fst
}
