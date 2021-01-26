package enkodo

import (
	"bytes"
	"io"
	"testing"
)

func TestReader_Decode(t *testing.T) {
	type testcase struct {
		val testStruct
		err error
	}

	tcs := []testcase{
		{val: testStruct{I64: 1}},
		{val: testStruct{I64: 2}},
		{val: testStruct{I64: 3}},
		{val: testStruct{I64: 4}},
		{err: io.EOF},
	}

	buf := bytes.NewBuffer(nil)
	w := NewWriter(buf)

	var err error
	for _, tc := range tcs {
		if tc.err != nil {
			continue
		}

		if err = w.Encode(&tc.val); err != nil {
			t.Fatal(err)
			return
		}
	}

	r := NewReader(buf)
	for _, tc := range tcs {
		var val testStruct
		if err = r.Decode(&val); err != tc.err {
			t.Fatalf("invalid error, expected <%v> and received <%v>", tc.err, err)
		}

		if !val.isMatch(&tc.val) {
			t.Fatalf("invalid value, expected %+v and received %+v", tc.val, val)
		}
	}
}
