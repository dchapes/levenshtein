package levenshtein

import "testing"

func TestSanity(t *testing.T) {
	tests := []struct {
		a, b string
		want int
	}{
		{"", "hello", 5},
		{"hello", "", 5},
		{"hello", "hello", 0},
		{"ab", "aa", 1},
		{"ab", "aaa", 2},
		{"bbb", "a", 3},
		{"kitten", "sitting", 3},
		{"distance", "difference", 5},
		{"levenshtein", "frankenstein", 6},
		{"resume and cafe", "resumes and cafes", 2},
		{"resumé and café", "resumés and cafés", 2},
		{"resume and cafe", "resumé and café", 2},
	}
	for i, d := range tests {
		n, err := ComputeDistance(d.a, d.b)
		if err != nil {
			t.Errorf("Test[%d]: ComputeDistance(%q,%q), error: %v", i, d.a, d.b, err)
			continue
		}
		if n != d.want {
			t.Errorf("Test[%d]: ComputeDistance(%q,%q) returned %v, want %v",
				i, d.a, d.b, n, d.want)
		}
	}
}
