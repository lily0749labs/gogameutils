package algoUtil

import (
	"reflect"
	"testing"
)

func TestUnset(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name  string
		index int
		input []int32
		want  []int32
	}{
		{name: "first", index: 0, input: []int32{1, 2, 3}, want: []int32{2, 3}},
		{name: "middle", index: 1, input: []int32{1, 2, 3}, want: []int32{1, 3}},
		{name: "last", index: 2, input: []int32{1, 2, 3}, want: []int32{1, 2}},
		{name: "negative", index: -1, input: []int32{1, 2, 3}, want: []int32{1, 2, 3}},
		{name: "out of range", index: 3, input: []int32{1, 2, 3}, want: []int32{1, 2, 3}},
		{name: "empty", index: 0, input: nil, want: nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Unset(tc.index, tc.input); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Unset() = %v, want %v", got, tc.want)
			}
		})
	}
}
