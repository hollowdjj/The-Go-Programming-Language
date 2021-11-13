package ch4

import (
	"bytes"
	"testing"
)

func TestPractice43(t *testing.T) {
	data := [4]int{1, 2, 3, 4}
	d := data
	Practice43(&data)
	if data != [4]int{4, 3, 2, 1} {
		t.Errorf("Practice(%q) = %v", d, data)
	}
}

func TestPractice44(t *testing.T) {
	data := []struct {
		source  []int
		num     int
		wanting []int
	}{
		{[]int{1, 2, 3}, 0, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 1, []int{2, 3, 1}},
		{[]int{1, 2, 3}, 2, []int{3, 1, 2}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
	}

	for _, d := range data {
		var c []int
		c = append(c, d.source...)
		if Practice44(d.source, d.num); !equal(d.wanting, d.source) {
			t.Errorf("Practice44(%v,%d) = %v", c, d.num, d.source)
		}
	}
}

func TestPractice45(t *testing.T) {
	data := []struct {
		source  []string
		wanting []string
	}{
		{[]string{"a", "a", "a", "b", "b"}, []string{"a", "b"}},
		{[]string{"a", "a", "c", "b"}, []string{"a", "c", "b"}},
		{[]string{"a", "b", "b", "b", "c", "a", "b", "b"}, []string{"a", "b", "c", "a", "b"}},
	}

	for _, v := range data {
		var c []string
		c = append(c, v.source...)
		if res := Practice45(v.source); !equalStringSlice(res, v.wanting) {
			t.Errorf("Practice45(%v) = %v", c, res)
		}
	}
}

func TestPractice46(t *testing.T) {
	data := []struct {
		source  []byte
		wanting []byte
	}{
		{[]byte("a b  c d"), []byte("a b c d")},
		{[]byte("a b 焯  草 "), []byte("a b 焯 草 ")},
	}

	for _, v := range data {
		var c []byte
		c = append(c, v.source...)
		if res := Practice46(v.source); string(res) != string(v.wanting) {
			t.Errorf("Practice46(%v) = %v", c, res)
		}
	}
}

func TestPractice47(t *testing.T) {
	data := []struct {
		source  []byte
		wanting []byte
	}{
		{[]byte("Hello,世界a"), []byte("a界世,olleH")},
		{[]byte("Hello,世界"), []byte("界世,olleH")},
		{[]byte("你Hello,世界a"), []byte("a界世,olleH你")},
	}

	for _, v := range data {
		var c []byte
		c = append(c, v.source...)
		if res := Practice47(v.source); !bytes.Equal(res, v.wanting) {
			t.Errorf("Practice(%s) = %s", c, v.source)
		}
	}

}
