package enkodo

import (
	"testing"
)

func Test_expandSlice(t *testing.T) {
	var (
		bs     []byte
		target = &bs
	)

	type testcase struct {
		expandTo int

		expectedLen int
		expectedCap int
	}

	tcs := []testcase{
		{
			expandTo:    8,
			expectedLen: 8,
			expectedCap: 8,
		},
		{
			expandTo:    2,
			expectedLen: 2,
			expectedCap: 8,
		},
		{
			expandTo:    4,
			expectedLen: 4,
			expectedCap: 8,
		},
	}

	for i, tc := range tcs {
		expandSlice(target, tc.expandTo)

		if len(*target) != tc.expectedLen {
			t.Fatalf("invalid length, expected %d and received %d (test case #%d)", tc.expectedLen, len(*target), i+1)
		}

		if cap(*target) != tc.expectedCap {
			t.Fatalf("invalid capacity, expected %d and received %d (test case #%d)", tc.expectedCap, cap(*target), i+1)
		}
	}
}
