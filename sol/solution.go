package sol

func trap(height []int) int {
	leftMax, rightMax, result := 0, 0, 0
	lp, rp := 0, len(height)-1
	for lp < rp {
		if height[lp] > height[rp] {
			if height[rp] > rightMax {
				rightMax = height[rp]
			} else {
				result += rightMax - height[rp]
			}
			rp--
		} else {
			if height[lp] > leftMax {
				leftMax = height[lp]
			} else {
				result += leftMax - height[lp]
			}
			lp++
		}
	}
	return result
}
