package mum

import (
	"io"
	"math"
)

// NewEncoder will return a new Encoder
func NewEncoder(w io.Writer) *Encoder {
	var e Encoder
	e.w = w
	return &e
}

// Encoder helps to Marshal data
type Encoder struct {
	w  io.Writer
	bw BinaryWriter
}

// Uint encodes a uint type
func (e *Encoder) Uint(v uint) (err error) {
	return e.Uint64(uint64(v))
}

// Uint8 encodes a uint8 type
func (e *Encoder) Uint8(v uint8) (err error) {
	_, err = e.w.Write(e.bw.Uint8(v))
	return
}

// Uint16 encodes a uint16 type
func (e *Encoder) Uint16(v uint16) (err error) {
	_, err = e.w.Write(e.bw.Uint16(v))
	return
}

// Uint32 encodes a uint32 type
func (e *Encoder) Uint32(v uint32) (err error) {
	_, err = e.w.Write(e.bw.Uint32(v))
	return
}

// Uint64 encodes a uint64 type
func (e *Encoder) Uint64(v uint64) (err error) {
	_, err = e.w.Write(e.bw.Uint64(v))
	return
}

// Int encodes an int type
func (e *Encoder) Int(v int) (err error) {
	return e.Int64(int64(v))
}

// Int8 encodes an int8 type
func (e *Encoder) Int8(v int8) (err error) {
	_, err = e.w.Write(e.bw.Int8(v))
	return
}

// Int16 encodes an int16 type
func (e *Encoder) Int16(v int16) (err error) {
	_, err = e.w.Write(e.bw.Int16(v))
	return
}

// Int32 encodes an int32 type
func (e *Encoder) Int32(v int32) (err error) {
	_, err = e.w.Write(e.bw.Int32(v))
	return
}

// Int64 encodes an int64 type
func (e *Encoder) Int64(v int64) (err error) {
	_, err = e.w.Write(e.bw.Int64(v))
	return
}

// Float32 encodes an float32 type
func (e *Encoder) Float32(v float32) (err error) {
	return e.Uint32(math.Float32bits(v))
}

// Float64 encodes an float64 type
func (e *Encoder) Float64(v float64) (err error) {
	return e.Uint64(math.Float64bits(v))
}

// Bytes will encode a byteslice to the writer
func (e *Encoder) Bytes(v []byte) (err error) {
	if err = e.Int(len(v)); err != nil {
		return
	}

	_, err = e.w.Write(v)
	return
}

// String will encode a string to the writer
func (e *Encoder) String(v string) (err error) {
	return e.Bytes(getStringBytes(v))
}

// Bool will encode a boolean value to the writer
func (e *Encoder) Bool(v bool) (err error) {
	if v {
		return e.Uint8(1)
	}

	return e.Uint8(0)
}

// Encode will encode an encodee
func (e *Encoder) Encode(v Encodee) (err error) {
	return v.MarshalMum(e)
}

// Encodee is a data structure to be encoded
type Encodee interface {
	MarshalMum(*Encoder) error
}
