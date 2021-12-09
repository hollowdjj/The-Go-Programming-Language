package ch11_Tests

import (
	"math/rand"
	"unicode"
)

//IsPalindrome 判断s是否是一个回文
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		//忽略空格
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r)) //忽略大小写
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

//randomPalindrome 随机生成一个回文
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) //在区间[0,25)之间返回一个int类型的随机数
	runes := make([]rune, n)

	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}

	return string(runes)
}
