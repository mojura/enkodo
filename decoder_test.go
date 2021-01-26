package enkodo

import (
	"bytes"
	"io"
	"testing"
)

func TestDecoder_Decode(t *testing.T) {
	type testcase struct {
		val int
		err error
	}

	tcs := []testcase{
		{val: 1},
		{val: 2},
		{val: 3},
		{val: 4},
		{err: io.EOF},
	}
	buf := bytes.NewBuffer(nil)
	enc := newEncoder(buf)

	var err error
	for _, tc := range tcs {
		if tc.err != nil {
			continue
		}

		if err = enc.Int(tc.val); err != nil {
			t.Fatal(err)
			return
		}
	}

	dec := newDecoder(buf)

	for _, tc := range tcs {
		var val int
		if val, err = dec.Int(); err != tc.err {
			t.Fatalf("invalid error, expected <%v> and received <%v>", tc.err, err)
		}

		if val != tc.val {
			t.Fatalf("invalid value, expected %d and received %d", tc.val, val)
		}
	}
}
