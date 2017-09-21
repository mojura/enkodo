package mum

import (
	"encoding/binary"
	"unsafe"

	"github.com/missionMeteora/toolkit/errors"
)

const (
	// ErrInvalidLength is returned when a byteslice has an invalid length for it's desired primitive
	ErrInvalidLength = errors.Error("invalid length")
)

// BinaryWriter will write numbers as binary bytes
type BinaryWriter struct {
	buf [8]byte
}

func (b *BinaryWriter) Uint8(v uint8) []byte {
	b.buf[0] = v
	return b.buf[:1]
}

func (b *BinaryWriter) Uint16(v uint16) []byte {
	binary.LittleEndian.PutUint16(b.buf[:2], v)
	return b.buf[:2]
}

func (b *BinaryWriter) Uint32(v uint32) []byte {
	binary.LittleEndian.PutUint32(b.buf[:4], v)
	return b.buf[:4]
}

func (b *BinaryWriter) Uint64(v uint64) []byte {
	binary.LittleEndian.PutUint64(b.buf[:], v)
	return b.buf[:]
}

func (b *BinaryWriter) Int8(v int8) []byte {
	b.buf[0] = uint8(v)
	return b.buf[:1]
}

func (b *BinaryWriter) Int16(v int16) []byte {
	return b.Uint16(uint16(v))
}

func (b *BinaryWriter) Int32(v int32) []byte {
	return b.Uint32(uint32(v))
}

func (b *BinaryWriter) Int64(v int64) []byte {
	return b.Uint64(uint64(v))
}

// BinaryReader will read numbers from binary bytes
type BinaryReader struct{}

func (b *BinaryReader) Uint16(bs []byte) (v uint16, err error) {
	if len(bs) != 2 {
		err = ErrInvalidLength
		return
	}

	v = binary.LittleEndian.Uint16(bs)
	return
}

func (b *BinaryReader) Uint32(bs []byte) (v uint32, err error) {
	if len(bs) != 4 {
		err = ErrInvalidLength
		return
	}

	v = binary.LittleEndian.Uint32(bs)
	return
}

func (b *BinaryReader) Uint64(bs []byte) (v uint64, err error) {
	if len(bs) != 8 {
		err = ErrInvalidLength
		return
	}

	v = binary.LittleEndian.Uint64(bs)
	return
}

func (b *BinaryReader) Int16(bs []byte) (v int16, err error) {
	if len(bs) != 2 {
		err = ErrInvalidLength
		return
	}

	v = int16(binary.LittleEndian.Uint16(bs))
	return
}

func (b *BinaryReader) Int32(bs []byte) (v int32, err error) {
	if len(bs) != 4 {
		err = ErrInvalidLength
		return
	}

	v = int32(binary.LittleEndian.Uint32(bs))
	return
}

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
