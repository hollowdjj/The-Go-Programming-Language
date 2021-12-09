package ch11_Tests

import (
	"math/rand"
	"testing"
	"time"
)

//必须导入testing包

/*
以_test.go结尾的文件为测试文件。其中的测试函数必须以Test开头，可选的后缀名必须以大写字母开头，例如：
func TestSin(t *testing.T) {.....}

go test ...
go test -v ... 可打印每个测试函数的名字和运行时间
go test -v -run="French|Canal" 参数-run对应一个正则表达式
*/

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") == false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") == false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("Palindrome") {
		t.Error(`IsPalindrome("Palindrome") == true`)
	}
}

func TestFrenchPalindrome(t *testing.T) {
	//Bug: 采用了byte序列而非rune序列，从而导致非ASCII字符不能被正确处理
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") == false`)
	}
}

func TestCanalPalindrme(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	//Bug：没有忽略空格和字母的大小写
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}

//表格驱动的测试
func TestIsPalindrome(t *testing.T) {

	tests := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, test.want)
		}
	}
}

//随机测试
func TestRandomPalindrome(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
