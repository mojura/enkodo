package enkodo

import (
	"fmt"
	"io"
	"math"
	"unsafe"
)

const (
	byte1Subtractor = (1 << 7)
	byte2Subtractor = (1<<7 + 1<<14)
	byte3Subtractor = (1<<7 + 1<<14 + 1<<21)
	byte4Subtractor = (1<<7 + 1<<14 + 1<<21 + 1<<28)
	byte5Subtractor = (1<<7 + 1<<14 + 1<<21 + 1<<28 + 1<<35)
	byte6Subtractor = (1<<7 + 1<<14 + 1<<21 + 1<<28 + 1<<35 + 1<<42)
	byte7Subtractor = (1<<7 + 1<<14 + 1<<21 + 1<<28 + 1<<35 + 1<<42 + 1<<49)
	byte8Subtractor = (1<<7 + 1<<14 + 1<<21 + 1<<28 + 1<<35 + 1<<42 + 1<<49 + 1<<56)
)

func decodeUint(r reader) (v uint, err error) {
	var u64 uint64
	if u64, err = decodeUint64(r); err != nil {
		return
	}

	v = uint(u64)
	return
}

func decodeUint8(r reader) (v uint8, err error) {
	return r.ReadByte()
}

func decodeUint16(r reader) (v uint16, err error) {
	var u64 uint64
	if u64, err = decodeUint64(r); err != nil {
		return
	}

	v = uint16(u64)
	return
}

func decodeUint32(r reader) (v uint32, err error) {
	var u64 uint64
	if u64, err = decodeUint64(r); err != nil {
		return
	}

	v = uint32(u64)
	return
}

func decodeUint64(r reader) (v uint64, err error) {
	var b byte
	if b, err = r.ReadByte(); err != nil {
		return
	}

	if v = uint64(b); v < ceiling {
		return
	}

	// Read next byte
	if b, err = r.ReadByte(); err != nil {
		return
	}

	v += uint64(b) << 7

	if b < ceiling {
		v -= byte1Subtractor
		return
	}

	// Read next byte
	if b, err = r.ReadByte(); err != nil {
		return
	}

	v += uint64(b) << 14
	if b < ceiling {
		v -= byte2Subtractor
		return
	}

	// Read next byte
	if b, err = r.ReadByte(); err != nil {
		return
	}

	v += uint64(b) << 21
	if b < ceiling {
		v -= byte3Subtractor
		return
	}

	// Read next byte
	if b, err = r.ReadByte(); err != nil {
		return
	}

	v += uint64(b) << 28
	if b < ceiling {
		v -= byte4Subtractor
		return
	}

	// Read next byte
	if b, err = r.ReadByte(); err != nil {
		return
	}

	v += uint64(b) << 35
	if b < ceiling {
		v -= byte5Subtractor
		return
	}

	// Read next byte
	if b, err = r.ReadByte(); err != nil {
		return
	}

	v += uint64(b) << 42
	if b < ceiling {
		v -= byte6Subtractor
		return
	}

	// Read next byte
	if b, err = r.ReadByte(); err != nil {
		return
	}

	v += uint64(b) << 49
	if b < ceiling {
		v -= byte7Subtractor
		return
	}

	// Read next byte
	if b, err = r.ReadByte(); err != nil {
		return
	}

	v += uint64(b) << 56
	v -= byte8Subtractor
	return
}

func decodeInt(r reader) (v int, err error) {
	var u64 int64
	if u64, err = decodeInt64(r); err != nil {
		return
	}

	v = int(u64)
	return
}

func decodeInt8(r reader) (v int8, err error) {
	var u8 uint8
	if u8, err = decodeUint8(r); err != nil {
		return
	}

	v = *(*int8)(unsafe.Pointer(&u8))
	return
}

func decodeInt16(r reader) (v int16, err error) {
	var i64 int64
	if i64, err = decodeInt64(r); err != nil {
		return
	}

	v = int16(i64)
	return
}

func decodeInt32(r reader) (v int32, err error) {
	var i64 int64
	if i64, err = decodeInt64(r); err != nil {
		return
	}

	v = int32(i64)
	return
}

func decodeInt64(r reader) (v int64, err error) {
	var u64 uint64
	if u64, err = decodeUint64(r); err != nil {
		return
	}

	v = *(*int64)(unsafe.Pointer(&u64))
	return
}

func decodeFloat32(r reader) (v float32, err error) {
	var u32 uint32
	if u32, err = decodeUint32(r); err != nil {
		return
	}

	v = math.Float32frombits(u32)
	return
}

func decodeFloat64(r reader) (v float64, err error) {
	var u64 uint64
	if u64, err = decodeUint64(r); err != nil {
		return
	}

	v = math.Float64frombits(u64)
	return
}

func decodeBytes(r reader, in *[]byte) (err error) {
	var bsLength int
	if bsLength, err = decodeInt(r); err != nil {
		err = fmt.Errorf("error decoding bytes length: %v", err)
		return
	}

	expandSlice(in, bsLength)

	if bsLength == 0 {
		// We do not have any bytes to read, return
		return
	}

	v := *in
	_, err = io.ReadAtLeast(r, v, bsLength)
	return
}

func decodeString(r reader) (str string, err error) {
	var bs []byte
	if err = decodeBytes(r, &bs); err != nil {
		return
	}

	str = getStringFromBytes(bs)
	return
}

func decodeBool(r reader) (v bool, err error) {
	var u8 uint8
	if u8, err = decodeUint8(r); err != nil {
		return
	}

	v = u8 == 1
	return
}
