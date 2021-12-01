package ch1

//Signum Go中的switch可以不带操作对象，此时默认用true代替。这种无tag switch和switch true是等价的
func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
}
