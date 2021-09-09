package enkodo

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

type schemaTestLarge struct {
	A []interface{}
	B map[string]interface{}

	Medium schemaTestMedium
}

type schemaTestMedium struct {
	A []string
	B []uint64
	C []bool

	D map[string]string
	E map[string]uint64
	F map[string]bool

	Small schemaTestSmall
}

type schemaTestSmall struct {
	A string
	B uint64
	C bool
}

func TestMakeSchema(t *testing.T) {
	var val schemaTestSmall
	val.A = "foo"
	val.B = 63
	val.C = true

	var err error
	if _, err = MakeSchema(val); err != nil {
		t.Fatal(err)
	}
}
func TestSchema_Marshal_Unmarshal(t *testing.T) {
	val := schemaTestLarge{
		A: []interface{}{
			"foo",
			"bar",
			"baz",
			1337,
			true,
			map[string]interface{}{
				"foo": "1",
				"bar": true,
				"baz": 1337,
			},
		},
		B: map[string]interface{}{
			"foo": "1",
			"bar": true,
			"baz": 1337,
		},
		Medium: schemaTestMedium{
			A: []string{"foo", "bar", "baz"},
			B: []uint64{63, 126, 811},
			C: []bool{true, false, true},
			D: map[string]string{
				"foo": "1",
				"bar": "2",
				"baz": "3",
			},
			E: map[string]uint64{
				"foo": 1,
				"bar": 2,
				"baz": 3,
			},
			F: map[string]bool{
				"foo": true,
				"bar": false,
				"baz": true,
			},
			Small: schemaTestSmall{
				A: "bar",
				B: 1337,
				C: true,
			},
		},
	}

	var (
		s   Schema
		err error
	)

	if s, err = MakeSchema(val.A); err != nil {
		t.Fatal(err)
	}

	buf := bytes.NewBuffer(nil)
	enc := newEncoder(buf)

	if err = s.MarshalEnkodo(enc, val.A); err != nil {
		t.Fatal(err)
	}

	dec := newDecoder(buf)

	var ref []interface{}
	rval := reflect.ValueOf(&ref)
	if err = s.UnmarshalEnkodo(dec, &rval); err != nil {
		t.Fatal(err)
	}

	fmt.Printf("\n%+v\n%+v\n", val.A, ref)
}
