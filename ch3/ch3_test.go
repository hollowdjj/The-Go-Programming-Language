package ch3

import "testing"

func TestBasename(t *testing.T) {
	data := []struct {
		source  string
		wanting string
	}{
		{"a", "a"},
		{"a.go", "a"},
		{"a/b/c.go", "c"},
		{"a/b.c.go", "b.c"},
	}

	for _, str := range data {
		if res := Basename(str.source); res != str.wanting {
			t.Errorf("Basename(%q) = %v", str.source, res)
		}
		if res := Basename1(str.source); res != str.wanting {
			t.Errorf("Basename1(%q) = %v", str.source, res)
		}
	}
}

func TestPrintInts(t *testing.T) {
	data := []struct {
		source  []int
		wanting string
	}{
		{[]int{}, "[]"},
		{[]int{1}, "[1]"},
		{[]int{1, 2}, "[1, 2]"},
		{[]int{1, 2, 3}, "[1, 2, 3]"},
	}

	for _, d := range data {
		if res := PrintInts(d.source); res != d.wanting {
			t.Errorf("PrintInts(%v) = %s", d.source, res)
		}
	}
}

func TestCommaInt(t *testing.T) {
	data := []struct {
		source  string
		wanting string
	}{
		{"12", "12"},
		{"123", "123"},
		{"12345", "12,345"},
		{"12345678", "12,345,678"},
	}

	for _, d := range data {
		if res := CommaInt(d.source); res != d.wanting {
			t.Errorf("CommaInt(%q) = %v", d.source, res)
		}
	}
}

func TestCommaIntUsingBuffer(t *testing.T) {
	data := []struct {
		source  string
		wanting string
	}{
		{"12", "12"},
		{"123", "123"},
		{"12345", "12,345"},
		{"123456", "123,456"},
		{"12345678", "12,345,678"},
	}

	for _, d := range data {
		if res := CommaIntUsingBuffer(d.source); res != d.wanting {
			t.Errorf("CommaInt(%q) = %v", d.source, res)
		}
	}
}

func TestPractice321(t *testing.T) {
	data := []struct {
		s1      string
		s2      string
		wanting bool
	}{
		{"aabbcc", "bbaacc", false},
		{"aabbcc", "ccaabb", true},
	}

	for _, d := range data {
		if res := Practice321(d.s1, d.s2); res != d.wanting {
			t.Errorf("Practice321(%q,%q) = %v", d.s1, d.s2, res)
		}
	}
}
