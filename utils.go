package mum

import (
	"encoding/binary"
	"unsafe"
)

// BinaryWriter will write numbers as binary bytes
type BinaryWriter struct {
	buf [8]byte
}

// Uint8 will return the byteslice representation of a uint8 value
func (b *BinaryWriter) Uint8(v uint8) []byte {
	b.buf[0] = v
	return b.buf[:1]
}

// Uint16 will return the byteslice representation of a uint16 value
func (b *BinaryWriter) Uint16(v uint16) []byte {
	binary.LittleEndian.PutUint16(b.buf[:2], v)
	return b.buf[:2]
}

// Uint32 will return the byteslice representation of a uint32 value
func (b *BinaryWriter) Uint32(v uint32) []byte {
	binary.LittleEndian.PutUint32(b.buf[:4], v)
	return b.buf[:4]
}

// Uint64 will return the byteslice representation of a uint64 value
func (b *BinaryWriter) Uint64(v uint64) []byte {
	binary.LittleEndian.PutUint64(b.buf[:], v)
	return b.buf[:]
}

// Int8 will return the byteslice representation of a int8 value
func (b *BinaryWriter) Int8(v int8) []byte {
	b.buf[0] = uint8(v)
	return b.buf[:1]
}

// Int16 will return the byteslice representation of a int16 value
func (b *BinaryWriter) Int16(v int16) []byte {
	return b.Uint16(uint16(v))
}

// Int32 will return the byteslice representation of a int32 value
func (b *BinaryWriter) Int32(v int32) []byte {
	return b.Uint32(uint32(v))
}

// Int64 will return the byteslice representation of a int64 value
func (b *BinaryWriter) Int64(v int64) []byte {
	return b.Uint64(uint64(v))
}

// BinaryReader will read numbers from binary bytes
type BinaryReader struct{}

// Uint16 will return the uint16 value from a provided byteslice
func (b *BinaryReader) Uint16(bs []byte) (v uint16, err error) {
	if len(bs) != 2 {
		err = ErrInvalidLength
		return
	}

	v = binary.LittleEndian.Uint16(bs)
	return
}

// Uint32 will return the uint32 value from a provided byteslice
func (b *BinaryReader) Uint32(bs []byte) (v uint32, err error) {
	if len(bs) != 4 {
		err = ErrInvalidLength
		return
	}

	v = binary.LittleEndian.Uint32(bs)
	return
}

// Uint64 will return the uint64 value from a provided byteslice
func (b *BinaryReader) Uint64(bs []byte) (v uint64, err error) {
	if len(bs) != 8 {
		err = ErrInvalidLength
		return
	}

	v = binary.LittleEndian.Uint64(bs)
	return
}

// Int16 will return the int16 value from a provided byteslice
func (b *BinaryReader) Int16(bs []byte) (v int16, err error) {
	if len(bs) != 2 {
		err = ErrInvalidLength
		return
	}

	v = int16(binary.LittleEndian.Uint16(bs))
	return
}

// Int32 will return the int32 value from a provided byteslice
func (b *BinaryReader) Int32(bs []byte) (v int32, err error) {
	if len(bs) != 4 {
		err = ErrInvalidLength
		return
	}

	v = int32(binary.LittleEndian.Uint32(bs))
	return
}

// Int64 will return the int64 value from a provided byteslice
func (b *BinaryReader) Int64(bs []byte) (v int64, err error) {
	if len(bs) != 8 {
		err = ErrInvalidLength
		return
	}

	v = int64(binary.LittleEndian.Uint64(bs))
	return
}

func getStringBytes(str string) []byte {
	return *((*[]byte)(unsafe.Pointer(&str)))
}

// Marshal will encode a value
func Marshal(v Encodee) (bs []byte) {
	return MarshalAppend(v, nil)
}

// MarshalAppend will encode a value to a provided slice
func MarshalAppend(v Encodee, buffer []byte) (bs []byte) {
	enc := newEncoder(buffer)
	enc.Encode(v)
	return enc.bs
}

// Unmarshal will decode a value
func Unmarshal(bs []byte, v Decodee) (err error) {
	dec := newDecoder(bs)
	return dec.Decode(v)
}
