package enkodo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"testing"
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

	e := newEncoder(nil)

	e.Int8(testNum)
	e.Int16(testNum)
	e.Int32(testNum)
	e.Int64(testNum)
	d := newDecoder(bytes.NewBuffer(e.bs))

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

	e := newEncoder(nil)

	e.Uint8(testNum)
	e.Uint16(testNum)
	e.Uint32(testNum)
	e.Uint64(testNum)
	d := newDecoder(bytes.NewBuffer(e.bs))

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

	e := newEncoder(nil)
	e.Float32(3.33)
	e.Float64(3.33)
	d := newDecoder(bytes.NewBuffer(e.bs))

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

	e := newEncoder(nil)
	e.Bool(true)
	d := newDecoder(bytes.NewBuffer(e.bs))

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

	e := newEncoder(nil)
	e.String("Hello world")
	d := newDecoder(bytes.NewBuffer(e.bs))

	if sv, err = d.String(); err != nil {
		t.Fatal(err)
	} else if sv != "Hello world" {
		t.Fatalf(testErrorFmt, "Hello world", sv)
	}
}

func Test_encodeUint64(t *testing.T) {
	var (
		bs  []byte
		err error
	)

	for i := uint64(0); i < 1_000_000; i++ {
		bs = encodeUint64(bs, i)

		var val uint64
		if val, err = decodeUint64(bytes.NewReader(bs)); err != nil {
			t.Fatal(err)
		}

		if val != i {
			t.Fatalf("invalid value, expected <%d> and received <%d>", i, val)
		}

		bs = bs[:0]
	}
}

func Test_encodeInt64(t *testing.T) {
	var (
		bs  []byte
		err error
	)

	for i := int64(-500_000); i < 500_000; i++ {
		bs = encodeInt64(bs, i)
		var val int64
		if val, err = decodeInt64(bytes.NewReader(bs)); err != nil {
			t.Fatal(err)
		}

		if val != i {
			t.Fatalf("invalid value, expected <%d> and received <%d>", i, val)
		}

		bs = bs[:0]
	}
}

func testietest(in []byte, v uint64) (buf []byte) {
	buf = in
	switch {
	case v < 1<<7-1:
		buf = append(buf, byte(v))
	case v < 1<<14-1:
		buf = append(buf, byte(v)|0x80, byte(v>>7))
	case v < 1<<21-1:
		buf = append(buf, byte(v)|0x80, byte(v>>7)|0x80, byte(v>>14))
	case v < 1<<28-1:
		buf = append(buf, byte(v)|0x80, byte(v>>7)|0x80, byte(v>>14)|0x80, byte(v>>21))
	case v < 1<<35-1:
		buf = append(buf, byte(v)|0x80, byte(v>>7)|0x80, byte(v>>14)|0x80, byte(v>>21)|0x80, byte(v>>28))
	case v < 1<<42-1:
		buf = append(buf, byte(v)|0x80, byte(v>>7)|0x80, byte(v>>14)|0x80, byte(v>>21)|0x80, byte(v>>28)|0x80, byte(v>>35))
	case v < 1<<49-1:
		buf = append(buf, byte(v)|0x80, byte(v>>7)|0x80, byte(v>>14)|0x80, byte(v>>21)|0x80, byte(v>>28)|0x80, byte(v>>35)|0x80, byte(v>>42))
	case v < 1<<56-1:
		buf = append(buf, byte(v)|0x80, byte(v>>7)|0x80, byte(v>>14)|0x80, byte(v>>21)|0x80, byte(v>>28)|0x80, byte(v>>35)|0x80, byte(v>>42)|0x80, byte(v>>49))
	default:
		buf = append(buf, byte(v)|0x80, byte(v>>7)|0x80, byte(v>>14)|0x80, byte(v>>21)|0x80, byte(v>>28)|0x80, byte(v>>35)|0x80, byte(v>>42)|0x80, byte(v>>49)|0x80, byte(v>>56))
	}

	return
}

func TestEncoderDecoder(t *testing.T) {
	var (
		a, b testStruct
		err  error
	)

	a = newTestStruct()
	inBuf := bytes.NewBuffer(nil)
	e := newEncoder(inBuf)
	if err = e.Encode(&a); err != nil {
		t.Fatal(err)
	}

	d := newDecoder(bytes.NewReader(inBuf.Bytes()))
	if err = d.Decode(&b); err != nil {
		return
	}

	if !a.isMatch(&b) {
		t.Fatal("Structs do not match")
	}
}

func BenchmarkEnkodoEncoding(b *testing.B) {
	var err error
	base := newTestStruct()
	b.ReportAllocs()
	b.ResetTimer()

	w := NewWriter(nil)
	for i := 0; i < b.N; i++ {
		if err = w.Encode(&base); err != nil {
			b.Fatal(err)
		}

		// We reset after each iteration so our buffer slice doesn't continuously grow
		w.Reset()
	}
}

func BenchmarkEnkodoDecoding(b *testing.B) {
	var err error
	base := newTestStruct()
	inBuf := bytes.NewBuffer(nil)
	e := newEncoder(inBuf)

	if err = e.Encode(&base); err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	buf := bytes.NewReader(inBuf.Bytes())
	r := NewReader(buf)
	for i := 0; i < b.N; i++ {
		fmt.Println("About to decode", inBuf.String())
		if err = r.Decode(&testVal); err != nil {
			b.Fatal(err)
		}

		if _, err = buf.Seek(0, 0); err != nil {
			b.Fatal(err)
		}
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

func BenchmarkEncodeInt64(b *testing.B) {
	w := NewWriter(nil)
	b.ReportAllocs()
	b.ResetTimer()

	var value int64
	for i := 0; i < b.N; i++ {
		w.e.Int64(value)

		// We reset after each iteration so our buffer slice doesn't continuously grow
		w.Reset()
	}
}

func BenchmarkEncoder_Uint64(b *testing.B) {
	w := NewWriter(nil)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.e.Uint64(uint64(i))

		// We reset after each iteration so our buffer slice doesn't continuously grow
		w.Reset()
	}
}

func BenchmarkEncodeUint64(b *testing.B) {
	var buf []byte
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf = encodeUint64(buf, uint64(i))
		buf = buf[:0]
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

func (t *testStruct) MarshalEnkodo(enc *Encoder) (err error) {
	enc.Int8(t.iv8)
	enc.Int16(t.iv16)
	enc.Int32(t.iv32)
	enc.Int64(t.iv64)
	enc.Uint8(t.uv8)
	enc.Uint16(t.uv16)
	enc.Uint32(t.uv32)
	enc.Uint64(t.uv64)
	enc.String(t.sv)
	enc.Bool(t.bv)
	return
}

func (t *testStruct) UnmarshalEnkodo(dec *Decoder) (err error) {
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
