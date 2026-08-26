package curUtil

import (
	"math"
	"testing"
)

func TestScoreToStrCur(t *testing.T) {
	t.Parallel()

	tests := []struct {
		score int64
		want  string
	}{
		{score: 0, want: "0.00"},
		{score: 1, want: "0.01"},
		{score: 12345, want: "123.45"},
		{score: -1, want: "-0.01"},
		{score: -12345, want: "-123.45"},
		{score: math.MinInt64, want: "-92233720368547758.08"},
	}

	for _, tc := range tests {
		if got := ScoreToStrCur(tc.score); got != tc.want {
			t.Errorf("ScoreToStrCur(%d) = %q, want %q", tc.score, got, tc.want)
		}
	}
}
