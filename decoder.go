package mum

import (
	"io"
)

// NewDecoder will return a new Decoder
func NewDecoder(r io.Reader) *Decoder {
	var d Decoder
	d.r = r
	// Initialize with a 64 byte buffer
	d.buf = make([]byte, 64)
	return &d
}

// Decoder helps to Marshal data
type Decoder struct {
	r   io.Reader
	buf []byte
	br  BinaryReader
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
	// Note: We do not have to bounds check because we know our buffer is
	// always going to be at LEAST 64 bytes

	// Read one byte from the reader
	if _, err = io.ReadAtLeast(d.r, d.buf[:1], 1); err != nil {
		return
	}

	v = uint8(d.buf[0])
	return
}

// Uint16 decodes a uint16 type
func (d *Decoder) Uint16() (v uint16, err error) {
	// Note: We do not have to bounds check because we know our buffer is
	// always going to be at LEAST 64 bytes

	// Read two bytes from the reader
	if _, err = io.ReadAtLeast(d.r, d.buf[:2], 2); err != nil {
		return
	}

	return d.br.Uint16(d.buf[:2])

}

// Uint32 decodes a uint32 type
func (d *Decoder) Uint32() (v uint32, err error) {
	// Note: We do not have to bounds check because we know our buffer is
	// always going to be at LEAST 64 bytes

	// Read four bytes from the reader
	if _, err = io.ReadAtLeast(d.r, d.buf[:4], 4); err != nil {
		return
	}

	return d.br.Uint32(d.buf[:4])

}

// Uint64 decodes a uint64 type
func (d *Decoder) Uint64() (v uint64, err error) {
	// Note: We do not have to bounds check because we know our buffer is
	// always going to be at LEAST 64 bytes

	// Read eight bytes from the reader
	if _, err = io.ReadAtLeast(d.r, d.buf[:8], 8); err != nil {
		return
	}

	return d.br.Uint64(d.buf[:8])
}

// Int decodes an int type
func (d *Decoder) Int() (v int, err error) {
	var v64 int64
	if v64, err = d.Int64(); err != nil {
		return
	}

	v = int(v64)
	return
}

// Int8 decodes an int8 type
func (d *Decoder) Int8() (v int8, err error) {
	// Note: We do not have to bounds check because we know our buffer is
	// always going to be at LEAST 64 bytes

	if _, err = io.ReadAtLeast(d.r, d.buf[:1], 1); err != nil {
		return
	}

	v = int8(d.buf[0])
	return
}

// Int16 decodes an int16 type
func (d *Decoder) Int16() (v int16, err error) {
	// Note: We do not have to bounds check because we know our buffer is
	// always going to be at LEAST 64 bytes

	bb := d.buf[:2]
	// Read eight bytes from the reader
	if _, err = io.ReadAtLeast(d.r, bb, 2); err != nil {
		return
	}

	return d.br.Int16(d.buf[:2])
}

// Int32 decodes an int32 type
func (d *Decoder) Int32() (v int32, err error) {
	// Note: We do not have to bounds check because we know our buffer is
	// always going to be at LEAST 64 bytes

	// Read eight bytes from the reader
	if _, err = io.ReadAtLeast(d.r, d.buf[:4], 4); err != nil {
		return
	}

	return d.br.Int32(d.buf[:4])
}

// Int64 decodes an int64 type
func (d *Decoder) Int64() (v int64, err error) {
	// Note: We do not have to bounds check because we know our buffer is
	// always going to be at LEAST 64 bytes

	// Read eight bytes from the reader
	if _, err = io.ReadAtLeast(d.r, d.buf[:8], 8); err != nil {
		return
	}

	return d.br.Int64(d.buf[:8])
}
