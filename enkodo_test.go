package enkodo

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"math"
	"testing"
)

const (
	testNum      = 5
	testErrorFmt = "Invalid value, expected %v and receIed %v"
)

var testVal testStruct

func TestInt(t *testing.T) {
	var (
		I8  int8
		I16 int16
		I32 int32
		I64 int64
		err error
	)

	e := newEncoder(nil)

	e.Int8(testNum)
	e.Int16(testNum)
	e.Int32(testNum)
	e.Int64(testNum)
	d := newDecoder(bytes.NewBuffer(e.bs))

	if I8, err = d.Int8(); err != nil {
		t.Fatal(err)
	} else if I8 != testNum {
		t.Fatalf(testErrorFmt, testNum, I8)
	}

	if I16, err = d.Int16(); err != nil {
		t.Fatal(err)
	} else if I16 != testNum {
		t.Fatalf(testErrorFmt, testNum, I16)
	}

	if I32, err = d.Int32(); err != nil {
		t.Fatal(err)
	} else if I32 != testNum {
		t.Fatalf(testErrorFmt, testNum, I32)
	}

	if I64, err = d.Int64(); err != nil {
		t.Fatal(err)
	} else if I64 != testNum {
		t.Fatalf(testErrorFmt, testNum, I64)
	}
}

func TestUint(t *testing.T) {
	var (
		U8  uint8
		U16 uint16
		U32 uint32
		U64 uint64
		err error
	)

	e := newEncoder(nil)

	e.Uint8(testNum)
	e.Uint16(testNum)
	e.Uint32(testNum)
	e.Uint64(testNum)
	d := newDecoder(bytes.NewBuffer(e.bs))

	if U8, err = d.Uint8(); err != nil {
		t.Fatal(err)
	} else if U8 != testNum {
		t.Fatalf(testErrorFmt, testNum, U8)
	}

	if U16, err = d.Uint16(); err != nil {
		t.Fatal(err)
	} else if U16 != testNum {
		t.Fatalf(testErrorFmt, testNum, U16)
	}

	if U32, err = d.Uint32(); err != nil {
		t.Fatal(err)
	} else if U32 != testNum {
		t.Fatalf(testErrorFmt, testNum, U32)
	}

	if U64, err = d.Uint64(); err != nil {
		t.Fatal(err)
	} else if U64 != testNum {
		t.Fatalf(testErrorFmt, testNum, U64)
	}
}

func TestFloat(t *testing.T) {
	var (
		F32 float32
		F64 float64
		err error
	)

	e := newEncoder(nil)
	e.Float32(3.33)
	e.Float64(3.33)
	d := newDecoder(bytes.NewBuffer(e.bs))

	if F32, err = d.Float32(); err != nil {
		t.Fatal(err)
	} else if F32 != 3.33 {
		t.Fatalf(testErrorFmt, 3.33, F32)
	}

	if F64, err = d.Float64(); err != nil {
		t.Fatal(err)
	} else if F64 != 3.33 {
		t.Fatalf(testErrorFmt, 3.33, F64)
	}
}

func TestBool(t *testing.T) {
	var (
		Bool bool
		err  error
	)

	e := newEncoder(nil)
	e.Bool(true)
	d := newDecoder(bytes.NewBuffer(e.bs))

	if Bool, err = d.Bool(); err != nil {
		t.Fatal(err)
	} else if !Bool {
		t.Fatalf(testErrorFmt, true, false)
	}
}

func TestString(t *testing.T) {
	var (
		str string
		err error
	)

	e := newEncoder(nil)
	e.String("Hello world")
	d := newDecoder(bytes.NewBuffer(e.bs))

	if str, err = d.String(); err != nil {
		t.Fatal(err)
	} else if str != "Hello world" {
		t.Fatalf(testErrorFmt, "Hello world", str)
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
			t.Fatalf("invalid value, expected <%d> and receIed <%d>", i, val)
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
			t.Fatalf("invalid value, expected <%d> and receIed <%d>", i, val)
		}

		bs = bs[:0]
	}
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
		if err = r.Decode(&testVal); err != nil {
			b.Fatal(err)
		}

		if _, err = buf.Seek(0, 0); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEnkodoDecoding_no_string(b *testing.B) {
	var err error
	base := newTestStruct()
	base.Str = ""
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

func BenchmarkJSONDecoding_no_string(b *testing.B) {
	var err error
	base := newTestStruct()
	base.Str = ""
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

func BenchmarkEnkodoEncoding_new_encoder(b *testing.B) {
	var err error
	base := newTestStruct()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w := NewWriter(nil)
		if err = w.Encode(&base); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEnkodoDecoding_new_decoder(b *testing.B) {
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

	for i := 0; i < b.N; i++ {
		r := NewReader(buf)
		if err = r.Decode(&testVal); err != nil {
			b.Fatal(err)
		}

		if _, err = buf.Seek(0, 0); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEnkodoDecoding_new_decoder_no_string(b *testing.B) {
	var err error
	base := newTestStruct()
	base.Str = ""
	inBuf := bytes.NewBuffer(nil)
	e := newEncoder(inBuf)

	if err = e.Encode(&base); err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	buf := bytes.NewReader(inBuf.Bytes())
	for i := 0; i < b.N; i++ {
		r := NewReader(buf)
		if err = r.Decode(&testVal); err != nil {
			b.Fatal(err)
		}

		if _, err = buf.Seek(0, 0); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGOBEncoding_new_decoder(b *testing.B) {
	var err error
	base := newTestStruct()
	buf := bytes.NewBuffer(nil)
	b.ReportAllocs()
	b.ResetTimer()

	e := gob.NewEncoder(buf)
	for i := 0; i < b.N; i++ {
		if err = e.Encode(&base); err != nil {
			b.Fatal(err)
		}

		// We reset after each iteration so our buffer slice doesn't continuously grow
		buf.Reset()
	}
}

func BenchmarkGOBDecoding_new_decoder(b *testing.B) {
	var err error
	base := newTestStruct()
	buf := bytes.NewBuffer(nil)

	if err = gob.NewEncoder(buf).Encode(&base); err != nil {
		b.Fatal(err)
	}

	r := bytes.NewReader(buf.Bytes())

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// When we re-use the gob decoder, we get a "extra data in buffer" error
		d := gob.NewDecoder(r)
		if err = d.Decode(&testVal); err != nil {
			b.Fatal(err)
		}

		// We reset after each iteration so our buffer slice doesn't continuously grow
		r.Reset(buf.Bytes())
	}
}

func BenchmarkGOBDecoding_new_decoder_no_string(b *testing.B) {
	var err error
	base := newTestStruct()
	base.Str = ""
	buf := bytes.NewBuffer(nil)

	if err = gob.NewEncoder(buf).Encode(&base); err != nil {
		b.Fatal(err)
	}

	r := bytes.NewReader(buf.Bytes())
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// When we re-use the gob decoder, we get a "extra data in buffer" error
		d := gob.NewDecoder(r)
		if err = d.Decode(&testVal); err != nil {
			b.Fatal(err)
		}

		// We reset after each iteration so our buffer slice doesn't continuously grow
		r.Reset(buf.Bytes())
	}
}

func BenchmarkJSONEncoding_new_encoder(b *testing.B) {
	var err error
	base := newTestStruct()
	buf := bytes.NewBuffer(nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e := json.NewEncoder(buf)
		if err = e.Encode(&base); err != nil {
			b.Fatal(err)
		}

		// We reset after each iteration so our buffer slice doesn't continuously grow
		buf.Reset()
	}
}

func BenchmarkJSONDecoding_new_decoder(b *testing.B) {
	var err error
	base := newTestStruct()
	buf := bytes.NewBuffer(nil)

	if err = json.NewEncoder(buf).Encode(&base); err != nil {
		b.Fatal(err)
	}

	r := bytes.NewReader(buf.Bytes())
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		d := json.NewDecoder(r)
		if err = d.Decode(&testVal); err != nil {
			b.Fatal(err)
		}

		// We reset after each iteration so our buffer slice doesn't continuously grow
		r.Reset(buf.Bytes())
	}
}

func BenchmarkJSONDecoding_new_decoder_no_string(b *testing.B) {
	var err error
	base := newTestStruct()
	base.Str = ""
	buf := bytes.NewBuffer(nil)

	if err = json.NewEncoder(buf).Encode(&base); err != nil {
		b.Fatal(err)
	}

	r := bytes.NewReader(buf.Bytes())
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		d := json.NewDecoder(r)
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
	t.I8 = math.MinInt8
	t.I16 = math.MinInt16
	t.I32 = math.MinInt32
	t.I64 = math.MinInt64

	t.U8 = math.MaxUint8
	t.U16 = math.MaxUint16
	t.U32 = math.MaxUint32
	t.U64 = math.MaxUint64

	t.F32 = math.MaxFloat32
	t.F64 = math.MaxFloat64

	t.Str = "Hello world!"
	t.Bytes = []byte(t.Str)
	t.Bool = true
	return
}

type testStruct struct {
	I8  int8
	I16 int16
	I32 int32
	I64 int64

	U8  uint8
	U16 uint16
	U32 uint32
	U64 uint64

	F32 float32
	F64 float64

	Str   string
	Bytes []byte
	Bool  bool
}

func (t *testStruct) MarshalEnkodo(enc *Encoder) (err error) {
	enc.Int8(t.I8)
	enc.Int16(t.I16)
	enc.Int32(t.I32)
	enc.Int64(t.I64)
	enc.Uint8(t.U8)
	enc.Uint16(t.U16)
	enc.Uint32(t.U32)
	enc.Uint64(t.U64)
	enc.String(t.Str)
	enc.Bytes(t.Bytes)
	enc.Bool(t.Bool)
	return
}

func (t *testStruct) UnmarshalEnkodo(dec *Decoder) (err error) {
	if t.I8, err = dec.Int8(); err != nil {
		return
	}

	if t.I16, err = dec.Int16(); err != nil {
		return
	}

	if t.I32, err = dec.Int32(); err != nil {
		return
	}

	if t.I64, err = dec.Int64(); err != nil {
		return
	}

	if t.U8, err = dec.Uint8(); err != nil {
		return
	}

	if t.U16, err = dec.Uint16(); err != nil {
		return
	}

	if t.U32, err = dec.Uint32(); err != nil {
		return
	}

	if t.U64, err = dec.Uint64(); err != nil {
		return
	}

	if t.Str, err = dec.String(); err != nil {
		return
	}

	if err = dec.Bytes(&t.Bytes); err != nil {
		return
	}

	if t.Bool, err = dec.Bool(); err != nil {
		return
	}

	return
}

func (t *testStruct) isMatch(c *testStruct) (match bool) {
	if t.I8 != c.I8 {
		return
	}

	if t.I16 != c.I16 {
		return
	}

	if t.I32 != c.I32 {
		return
	}

	if t.I64 != c.I64 {
		return
	}

	if t.U8 != c.U8 {
		return
	}

	if t.U16 != c.U16 {
		return
	}

	if t.U32 != c.U32 {
		return
	}

	if t.U64 != c.U64 {
		return
	}

	if t.Str != c.Str {
		return
	}

	if t.Bool != c.Bool {
		return
	}

	return true
}
