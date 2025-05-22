package my_math

func Fact(n int) int {
	f := 1
	for i := 1; i <= n; i++ {
		f = f * i
	}

	return f
}

func MaxNum(n []int) int {
	//var maxNum int
	maxNum := n[0]

	for _, i := range n {
		if i > maxNum {
			maxNum = i
		}
	}
	return maxNum
}
