package main

import "testing"

func TestLengthOfNonRepeatingSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcbb", 3},
		{"一二三四一二三一二一", 4},
		//{"abcabcbb",4},
	}

	for _, tt := range tests {
		if actual := lengthOfNonRepeatingSubstr(tt.s); actual != tt.ans {
			t.Errorf("lengthOfNonRepeatingSubstr(%s);"+"got %d;expected %d", tt.s, actual, tt.ans)
		}
	}

}

func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans := 8

	for i := 0; i < b.N; i++ {
		if actual := lengthOfNonRepeatingSubstr(s); actual != ans {
			b.Errorf("lengthOfNonRepeatingSubstr(%s);"+"got %d;expected %d", s, actual, ans)
		}
	}

}
