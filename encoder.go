package enkodo

import "io"

func newEncoder(w io.Writer) *Encoder {
	var e Encoder
	e.w = w
	return &e
}

// Encoder helps to Marshal data
type Encoder struct {
	bs []byte
	w  io.Writer

	// Total written bytes
	written int64
}

func (e *Encoder) flush() (err error) {
	if e.w == nil {
		return
	}

	var n int
	if n, err = e.w.Write(e.bs); err == nil {
		e.written += int64(n)
	}

	e.bs = e.bs[:0]
	return
}

// Encode will encode an encodee
func (e *Encoder) teardown() {
	e.bs = nil
	e.w = nil
}

// Uint encodes a uint type
func (e *Encoder) Uint(v uint) (err error) {
	e.bs = encodeUint(e.bs, v)
	return e.flush()
}

// Uint8 encodes a uint8 type
func (e *Encoder) Uint8(v uint8) (err error) {
	e.bs = encodeUint8(e.bs, v)
	return e.flush()
}

// Uint16 encodes a uint16 type
func (e *Encoder) Uint16(v uint16) (err error) {
	e.bs = encodeUint16(e.bs, v)
	return e.flush()
}

// Uint32 encodes a uint32 type
func (e *Encoder) Uint32(v uint32) (err error) {
	e.bs = encodeUint32(e.bs, v)
	return e.flush()
}

// Uint64 encodes a uint64 type
func (e *Encoder) Uint64(v uint64) (err error) {
	e.bs = encodeUint64(e.bs, v)
	return e.flush()
}

// Int encodes an int type
func (e *Encoder) Int(v int) (err error) {
	e.bs = encodeInt(e.bs, v)
	return e.flush()
}

// Int8 encodes an int8 type
func (e *Encoder) Int8(v int8) (err error) {
	e.bs = encodeInt8(e.bs, v)
	return e.flush()
}

// Int16 encodes an int16 type
func (e *Encoder) Int16(v int16) (err error) {
	e.bs = encodeInt16(e.bs, v)
	return e.flush()
}

// Int32 encodes an int32 type
func (e *Encoder) Int32(v int32) (err error) {
	e.bs = encodeInt32(e.bs, v)
	return e.flush()
}

// Int64 encodes an int64 type
func (e *Encoder) Int64(v int64) (err error) {
	e.bs = encodeInt64(e.bs, v)
	return e.flush()
}

// Float32 encodes an float32 type
func (e *Encoder) Float32(v float32) (err error) {
	e.bs = encodeFloat32(e.bs, v)
	return e.flush()
}

// Float64 encodes an float64 type
func (e *Encoder) Float64(v float64) (err error) {
	e.bs = encodeFloat64(e.bs, v)
	return e.flush()
}

// Bytes will encode a byteslice to the writer
func (e *Encoder) Bytes(v []byte) (err error) {
	e.bs = encodeBytes(e.bs, v)
	return e.flush()
}

// String will encode a string to the writer
func (e *Encoder) String(v string) (err error) {
	e.bs = encodeString(e.bs, v)
	return e.flush()
}

// Bool will encode a boolean value to the writer
func (e *Encoder) Bool(v bool) (err error) {
	e.bs = encodeBool(e.bs, v)
	return e.flush()
}

// Encode will encode an encodee
func (e *Encoder) Encode(v interface{}) (err error) {
	enc, ok := v.(Encodee)
	if ok {
		return enc.MarshalEnkodo(e)
	}

	var s Schema
	if s, err = MakeSchema(v); err != nil {
		return
	}

	return s.MarshalEnkodo(e, v)
}

// Encodee is a data structure to be encoded
type Encodee interface {
	MarshalEnkodo(*Encoder) error
}
