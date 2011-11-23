package nlp

func min(x int, y int) int {
if x < y {
	return x
}
return y
}

func SplitFunc(data []byte, f func(rune int) bool) [][]byte {
	start := 0
	a := make([][]byte, 0)
	if len(data) <= 1 {
		return append(a,data)
	}
	for end := 1;end<len(data);end++{
		if f(int(data[end])) {
			a = append(a, data[start:end])
			start = end + 1
			end++
		}
	}
	return a
}