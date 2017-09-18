package mum

import (
	"io"
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
	_, err = e.w.Write([]byte{v})
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
	_, err = e.w.Write([]byte{byte(v)})
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
