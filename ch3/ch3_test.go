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
