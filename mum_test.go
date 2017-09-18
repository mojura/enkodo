package mum

import (
	"bytes"
	"testing"
)

const (
	testNum      = 5
	testErrorFmt = "Invalid value, expected %v and received %v"
)

func TestIntegers(t *testing.T) {
	var (
		iv8  int8
		iv16 int16
		iv32 int32
		iv64 int64

		uv8  uint8
		uv16 uint16
		uv32 uint32
		uv64 uint64
		err  error
	)

	b := bytes.NewBuffer(nil)
	e := NewEncoder(b)
	d := NewDecoder(b)

	e.Int8(testNum)
	e.Int16(testNum)
	e.Int32(testNum)
	e.Int64(testNum)

	e.Uint8(testNum)
	e.Uint16(testNum)
	e.Uint32(testNum)
	e.Uint64(testNum)

	e.Float32(3.33)
	e.Float64(3.33)

	if iv8, err = d.Int8(); err != nil {
		t.Fatal(err)
	} else if iv8 != testNum {
		t.Fatalf(testErrorFmt, testNum, iv8)
	}

	if iv16, err = d.Int16(); err != nil {
		t.Fatal(err)
	} else if iv16 != testNum {
		t.Fatalf(testErrorFmt, testNum, iv16)
	}

	if iv32, err = d.Int32(); err != nil {
		t.Fatal(err)
	} else if iv32 != testNum {
		t.Fatalf(testErrorFmt, testNum, iv32)
	}

	if iv64, err = d.Int64(); err != nil {
		t.Fatal(err)
	} else if iv64 != testNum {
		t.Fatalf(testErrorFmt, testNum, iv64)
	}

	if uv8, err = d.Uint8(); err != nil {
		t.Fatal(err)
	} else if uv8 != testNum {
		t.Fatalf(testErrorFmt, testNum, uv8)
	}

	if uv16, err = d.Uint16(); err != nil {
		t.Fatal(err)
	} else if uv16 != testNum {
		t.Fatalf(testErrorFmt, testNum, uv16)
	}

	if uv32, err = d.Uint32(); err != nil {
		t.Fatal(err)
	} else if uv32 != testNum {
		t.Fatalf(testErrorFmt, testNum, uv32)
	}

	if uv64, err = d.Uint64(); err != nil {
		t.Fatal(err)
	} else if uv64 != testNum {
		t.Fatalf(testErrorFmt, testNum, uv64)
	}

	var (
		f32 float32
		f64 float64
	)

	if f32, err = d.Float32(); err != nil {
		t.Fatal(err)
	} else if f32 != 3.33 {
		t.Fatalf(testErrorFmt, 3.33, f32)
	}

	if f64, err = d.Float64(); err != nil {
		t.Fatal(err)
	} else if f64 != 3.33 {
		t.Fatalf(testErrorFmt, 3.33, f64)
	}

}
