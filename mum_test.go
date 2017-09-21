package mum

import (
	"bytes"
	"encoding/json"
	"math"
	"testing"

	"github.com/missionMeteora/binny.v2"
)

const (
	testNum      = 5
	testErrorFmt = "Invalid value, expected %v and received %v"
)

var testVal testStruct

func TestInt(t *testing.T) {
	var (
		iv8  int8
		iv16 int16
		iv32 int32
		iv64 int64
		err  error
	)

	b := bytes.NewBuffer(nil)
	e := NewEncoder(b)
	d := NewDecoder(b)

	e.Int8(testNum)
	e.Int16(testNum)
	e.Int32(testNum)
	e.Int64(testNum)

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
}

func TestUint(t *testing.T) {
	var (
		uv8  uint8
		uv16 uint16
		uv32 uint32
		uv64 uint64
		err  error
	)

	b := bytes.NewBuffer(nil)
	e := NewEncoder(b)
	d := NewDecoder(b)

	e.Uint8(testNum)
	e.Uint16(testNum)
	e.Uint32(testNum)
	e.Uint64(testNum)

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
}

func TestFloat(t *testing.T) {
	var (
		f32 float32
		f64 float64
		err error
	)

	b := bytes.NewBuffer(nil)
	e := NewEncoder(b)
	d := NewDecoder(b)

	e.Float32(3.33)
	e.Float64(3.33)

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

func TestBool(t *testing.T) {
	var (
		bv  bool
		err error
	)

	b := bytes.NewBuffer(nil)
	e := NewEncoder(b)
	d := NewDecoder(b)

	e.Bool(true)

	if bv, err = d.Bool(); err != nil {
		t.Fatal(err)
	} else if !bv {
		t.Fatalf(testErrorFmt, true, false)
	}
}

func TestString(t *testing.T) {
	var (
		sv  string
		err error
	)

	b := bytes.NewBuffer(nil)
	e := NewEncoder(b)
	d := NewDecoder(b)

	e.String("Hello world")

	if sv, err = d.String(); err != nil {
		t.Fatal(err)
	} else if sv != "Hello world" {
		t.Fatalf(testErrorFmt, "Hello world", sv)
	}
}

func TestEncoderDecoder(t *testing.T) {
	var (
		a, b testStruct
		err  error

		buf = bytes.NewBuffer(nil)
	)

	a = newTestStruct()
	e := NewEncoder(buf)

	if err = e.Encode(&a); err != nil {
		return
	}

	d := NewDecoder(buf)
	if err = d.Decode(&b); err != nil {
		return
	}

	if !a.isMatch(&b) {
		t.Fatal("Structs do not match")
	}
}

func BenchmarkMumEncoding(b *testing.B) {
	var err error
	base := newTestStruct()
	buf := bytes.NewBuffer(nil)
	b.ReportAllocs()
	b.ResetTimer()

	e := NewEncoder(buf)
	for i := 0; i < b.N; i++ {
		if err = e.Encode(&base); err != nil {
			b.Fatal(err)
		}

		// We reset after each iteration so our buffer slice doesn't continuously grow
		buf.Reset()
	}
}

func BenchmarkMumDecoding(b *testing.B) {
	var err error
	base := newTestStruct()
	buf := bytes.NewBuffer(nil)

	if err = NewEncoder(buf).Encode(&base); err != nil {
		b.Fatal(err)
	}

	r := bytes.NewReader(buf.Bytes())

	b.ReportAllocs()
	b.ResetTimer()

	d := NewDecoder(r)
	for i := 0; i < b.N; i++ {
		if err = d.Decode(&testVal); err != nil {
			b.Fatal(err)
		}

		// We reset after each iteration so our buffer slice doesn't continuously grow
		r.Reset(buf.Bytes())
	}
}

func BenchmarkBinnyEncoding(b *testing.B) {
	var err error
	base := newTestStruct()
	buf := bytes.NewBuffer(nil)
	b.ReportAllocs()
	b.ResetTimer()

	e := binny.NewEncoder(buf)
	for i := 0; i < b.N; i++ {
		if err = e.Encode(&base); err != nil {
			b.Fatal(err)
		}

		// We reset after each iteration so our buffer slice doesn't continuously grow
		buf.Reset()
	}
}

func BenchmarkBinnyDecoding(b *testing.B) {
	var err error
	base := newTestStruct()
	buf := bytes.NewBuffer(nil)

	if err = binny.NewEncoder(buf).Encode(&base); err != nil {
		b.Fatal(err)
	}

	r := bytes.NewReader(buf.Bytes())
	d := binny.NewDecoder(r)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err = d.Decode(&testVal); err != nil {
			b.Fatal(err)
		}

		// We reset after each iteration so our buffer slice doesn't continuously grow
		r.Reset(buf.Bytes())
	}
}

func BenchmarkJSONEncoding(b *testing.B) {
	var err error
	base := newTestStruct()
	buf := bytes.NewBuffer(nil)
	b.ReportAllocs()
	b.ResetTimer()

	e := json.NewEncoder(buf)
	for i := 0; i < b.N; i++ {
		if err = e.Encode(&base); err != nil {
			b.Fatal(err)
		}

		// We reset after each iteration so our buffer slice doesn't continuously grow
		buf.Reset()
	}
}

func BenchmarkJSONDecoding(b *testing.B) {
	var err error
	base := newTestStruct()
	buf := bytes.NewBuffer(nil)

	if err = json.NewEncoder(buf).Encode(&base); err != nil {
		b.Fatal(err)
	}

	r := bytes.NewReader(buf.Bytes())
	d := json.NewDecoder(r)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err = d.Decode(&testVal); err != nil {
			b.Fatal(err)
		}

		// We reset after each iteration so our buffer slice doesn't continuously grow
		r.Reset(buf.Bytes())
	}
}

func newTestStruct() (t testStruct) {
	t.iv8 = math.MinInt8
	t.iv16 = math.MinInt16
	t.iv32 = math.MinInt32
	t.iv64 = math.MinInt64

	t.uv8 = math.MaxUint8
	t.uv16 = math.MaxUint16
	t.uv32 = math.MaxUint32
	t.uv64 = math.MaxUint64

	t.f32 = math.MaxFloat32
	t.f64 = math.MaxFloat64

	t.sv = "Hello world!"
	t.bv = true
	return
}

type testStruct struct {
	iv8  int8
	iv16 int16
	iv32 int32
	iv64 int64

	uv8  uint8
	uv16 uint16
	uv32 uint32
	uv64 uint64

	f32 float32
	f64 float64

	sv string
	bv bool
}

func (t *testStruct) MarshalMum(enc *Encoder) (err error) {
	if err = enc.Int8(t.iv8); err != nil {
		return
	}

	if err = enc.Int16(t.iv16); err != nil {
		return
	}

	if err = enc.Int32(t.iv32); err != nil {
		return
	}

	if err = enc.Int64(t.iv64); err != nil {
		return
	}

	if err = enc.Uint8(t.uv8); err != nil {
		return
	}

	if err = enc.Uint16(t.uv16); err != nil {
		return
	}

	if err = enc.Uint32(t.uv32); err != nil {
		return
	}

	if err = enc.Uint64(t.uv64); err != nil {
		return
	}

	if err = enc.String(t.sv); err != nil {
		return
	}

	if err = enc.Bool(t.bv); err != nil {
		return
	}

	return
}

func (t *testStruct) UnmarshalMum(dec *Decoder) (err error) {
	if t.iv8, err = dec.Int8(); err != nil {
		return
	}

	if t.iv16, err = dec.Int16(); err != nil {
		return
	}

	if t.iv32, err = dec.Int32(); err != nil {
		return
	}

	if t.iv64, err = dec.Int64(); err != nil {
		return
	}

	if t.uv8, err = dec.Uint8(); err != nil {
		return
	}

	if t.uv16, err = dec.Uint16(); err != nil {
		return
	}

	if t.uv32, err = dec.Uint32(); err != nil {
		return
	}

	if t.uv64, err = dec.Uint64(); err != nil {
		return
	}

	if t.sv, err = dec.String(); err != nil {
		return
	}

	if t.bv, err = dec.Bool(); err != nil {
		return
	}

	return
}

func (t *testStruct) MarshalBinny(enc *binny.Encoder) (err error) {
	if err = enc.WriteInt8(t.iv8); err != nil {
		return
	}

	if err = enc.WriteInt16(t.iv16); err != nil {
		return
	}

	if err = enc.WriteInt32(t.iv32); err != nil {
		return
	}

	if err = enc.WriteInt64(t.iv64); err != nil {
		return
	}

	if err = enc.WriteUint8(t.uv8); err != nil {
		return
	}

	if err = enc.WriteUint16(t.uv16); err != nil {
		return
	}

	if err = enc.WriteUint32(t.uv32); err != nil {
		return
	}

	if err = enc.WriteUint64(t.uv64); err != nil {
		return
	}

	if err = enc.WriteString(t.sv); err != nil {
		return
	}

	if err = enc.WriteBool(t.bv); err != nil {
		return
	}

	return
}

func (t *testStruct) UnmarshalBinny(dec *binny.Decoder) (err error) {
	if t.iv8, err = dec.ReadInt8(); err != nil {
		return
	}

	if t.iv16, err = dec.ReadInt16(); err != nil {
		return
	}

	if t.iv32, err = dec.ReadInt32(); err != nil {
		return
	}

	if t.iv64, err = dec.ReadInt64(); err != nil {
		return
	}

	if t.uv8, err = dec.ReadUint8(); err != nil {
		return
	}

	if t.uv16, err = dec.ReadUint16(); err != nil {
		return
	}

	if t.uv32, err = dec.ReadUint32(); err != nil {
		return
	}

	if t.uv64, err = dec.ReadUint64(); err != nil {
		return
	}

	if t.sv, err = dec.ReadString(); err != nil {
		return
	}

	if t.bv, err = dec.ReadBool(); err != nil {
		return
	}

	return
}

func (t *testStruct) isMatch(c *testStruct) (match bool) {
	if t.iv8 != c.iv8 {
		return
	}

	if t.iv16 != c.iv16 {
		return
	}

	if t.iv32 != c.iv32 {
		return
	}

	if t.iv64 != c.iv64 {
		return
	}

	if t.uv8 != c.uv8 {
		return
	}

	if t.uv16 != c.uv16 {
		return
	}

	if t.uv32 != c.uv32 {
		return
	}

	if t.uv64 != c.uv64 {
		return
	}

	if t.sv != c.sv {
		return
	}

	if t.bv != c.bv {
		return
	}

	return true
}
