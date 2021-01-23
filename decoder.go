package enkodo

import (
	"bufio"
	"io"
)

func newDecoder(r io.Reader) *Decoder {
	var (
		d  Decoder
		ok bool
	)

	if d.r, ok = r.(reader); !ok {
		d.r = bufio.NewReader(r)
	}

	return &d
}

// Decoder helps to Marshal data
type Decoder struct {
	r reader
}

// Uint decodes a uint type
func (d *Decoder) Uint() (v uint, err error) {
	var v64 uint64
	if v64, err = d.Uint64(); err != nil {
		return
	}

	v = uint(v64)
	return
}

// Uint8 decodes a uint8 type
func (d *Decoder) Uint8() (v uint8, err error) {
	v, err = decodeUint8(d.r)
	return
}

// Uint16 decodes a uint16 type
func (d *Decoder) Uint16() (v uint16, err error) {
	v, err = decodeUint16(d.r)
	return
}

// Uint32 decodes a uint32 type
func (d *Decoder) Uint32() (v uint32, err error) {
	v, err = decodeUint32(d.r)
	return
}

// Uint64 decodes a uint64 type
func (d *Decoder) Uint64() (v uint64, err error) {
	v, err = decodeUint64(d.r)
	return
}

// Int decodes an int type
func (d *Decoder) Int() (v int, err error) {
	v, err = decodeInt(d.r)
	return
}

// Int8 decodes an int8 type
func (d *Decoder) Int8() (v int8, err error) {
	v, err = decodeInt8(d.r)
	return
}

// Int16 decodes an int16 type
func (d *Decoder) Int16() (v int16, err error) {
	v, err = decodeInt16(d.r)
	return
}

// Int32 decodes an int32 type
func (d *Decoder) Int32() (v int32, err error) {
	v, err = decodeInt32(d.r)
	return
}

// Int64 decodes an int64 type
func (d *Decoder) Int64() (v int64, err error) {
	v, err = decodeInt64(d.r)
	return
}

// Float32 decodes a float64 type
func (d *Decoder) Float32() (v float32, err error) {
	v, err = decodeFloat32(d.r)
	return
}

// Float64 decodes a float64 type
func (d *Decoder) Float64() (v float64, err error) {
	v, err = decodeFloat64(d.r)
	return
}

// Bool will return a decoded boolean value
func (d *Decoder) Bool() (v bool, err error) {
	v, err = decodeBool(d.r)
	return
}

// Bytes will append bytes to the inbound byteslice
func (d *Decoder) Bytes(in *[]byte) (err error) {
	return decodeBytes(d.r, in)
}

// String will return a decoded string
func (d *Decoder) String() (str string, err error) {
	return decodeString(d.r)
}

// Decode will decode a decodee
func (d *Decoder) Decode(v Decodee) (err error) {
	return v.UnmarshalEnkodo(d)
}

// Decodee is a data structure to be dedoded
type Decodee interface {
	UnmarshalEnkodo(*Decoder) error
}
