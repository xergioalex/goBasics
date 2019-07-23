package numbers

func GetVariables() (int, int32, int64) {
	return 1, 2147000, 90033223233
}

func GetFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

func Sum(a int, b int) int {
	return a + b
}