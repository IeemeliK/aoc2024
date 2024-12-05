package validator

func absVal(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func MinDiff(list []int, min int) bool {
	for i := 1; i <= len(list)-1; i++ {
		num := absVal(list[i] - list[i-1])
		if num < min {
			return false
		}
	}
	return true
}

func MaxDiff(list []int, max int) bool {
	for i := 1; i <= len(list)-1; i++ {
		num := absVal(list[i] - list[i-1])
		if num > max {
			return false
		}
	}
	return true
}

func Decreasing(list []int) bool {
	for i := 1; i <= len(list)-1; i++ {
		if list[i] > list[i-1] {
			return false
		}
	}
	return true
}

func Increasing(list []int) bool {
	for i := 1; i <= len(list)-1; i++ {
		if list[i] < list[i-1] {
			return false
		}
	}
	return true
}
