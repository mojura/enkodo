package mum

import (
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

func decodeUint(bs []byte) (v uint, remaining []byte, err error) {
	var u64 uint64
	if u64, remaining, err = decodeUint64(bs); err != nil {
		return
	}

	v = uint(u64)
	return
}

func decodeUint8(bs []byte) (v uint8, remaining []byte, err error) {
	v = bs[0]
	remaining = bs[1:]
	return
}

func decodeUint16(bs []byte) (v uint16, remaining []byte, err error) {
	var u64 uint64
	if u64, remaining, err = decodeUint64(bs); err != nil {
		return
	}

	v = uint16(u64)
	return
}

func decodeUint32(bs []byte) (v uint32, remaining []byte, err error) {
	var u64 uint64
	if u64, remaining, err = decodeUint64(bs); err != nil {
		return
	}

	v = uint32(u64)
	return
}

func decodeUint64(bs []byte) (v uint64, remaining []byte, err error) {
	if len(bs) == 0 {
		err = ErrEmptyBytes
		return
	}

	if v = uint64(bs[0]); v < ceiling {
		remaining = bs[1:]
		return
	}

	second := bs[1]
	v += uint64(second) << 7

	if second < ceiling {
		v -= byte1Subtractor
		remaining = bs[2:]
		return
	}

	third := bs[2]
	v += uint64(third) << 14
	if third < ceiling {
		v -= byte2Subtractor
		remaining = bs[3:]
		return
	}

	forth := bs[3]
	v += uint64(forth) << 21
	if forth < ceiling {
		v -= byte3Subtractor
		remaining = bs[4:]
		return
	}

	fifth := bs[4]
	v += uint64(fifth) << 28
	if fifth < ceiling {
		v -= byte4Subtractor
		remaining = bs[5:]
		return
	}

	sixth := bs[5]
	v += uint64(sixth) << 35
	if sixth < ceiling {
		v -= byte5Subtractor
		remaining = bs[6:]
		return
	}

	seventh := bs[6]
	v += uint64(seventh) << 42
	if seventh < ceiling {
		v -= byte6Subtractor
		remaining = bs[7:]
		return
	}

	eighth := bs[7]
	v += uint64(eighth) << 49
	if eighth < ceiling {
		v -= byte7Subtractor
		remaining = bs[8:]
		return
	}

	v += uint64(bs[8]) << 56
	v -= byte8Subtractor
	remaining = bs[9:]
	return
}

func decodeInt(bs []byte) (v int, remaining []byte, err error) {
	var u64 int64
	if u64, remaining, err = decodeInt64(bs); err != nil {
		return
	}

	v = int(u64)
	return
}

func decodeInt8(bs []byte) (v int8, remaining []byte, err error) {
	var u8 uint8
	if u8, remaining, err = decodeUint8(bs); err != nil {
		return
	}

	v = *(*int8)(unsafe.Pointer(&u8))
	return
}

func decodeInt16(bs []byte) (v int16, remaining []byte, err error) {
	var i64 int64
	if i64, remaining, err = decodeInt64(bs); err != nil {
		return
	}

	v = int16(i64)
	return
}

func decodeInt32(bs []byte) (v int32, remaining []byte, err error) {
	var i64 int64
	if i64, remaining, err = decodeInt64(bs); err != nil {
		return
	}

	v = int32(i64)
	return
}

func decodeInt64(bs []byte) (v int64, remaining []byte, err error) {
	var u64 uint64
	if u64, remaining, err = decodeUint64(bs); err != nil {
		return
	}

	v = *(*int64)(unsafe.Pointer(&u64))
	return
}

func decodeFloat32(bs []byte) (v float32, remaining []byte, err error) {
	var u32 uint32
	if u32, remaining, err = decodeUint32(bs); err != nil {
		return
	}

	v = math.Float32frombits(u32)
	return
}

func decodeFloat64(bs []byte) (v float64, remaining []byte, err error) {
	var u64 uint64
	if u64, remaining, err = decodeUint64(bs); err != nil {
		return
	}

	v = math.Float64frombits(u64)
	return
}

func decodeBytes(bs []byte) (v []byte, remaining []byte, err error) {
	var len int
	if len, remaining, err = decodeInt(bs); err != nil {
		return
	}

	v = remaining[:len]
	remaining = remaining[len:]
	return
}

func decodeString(bs []byte) (v string, remaining []byte, err error) {
	var data []byte
	if data, remaining, err = decodeBytes(bs); err != nil {
		return
	}

	v = string(data)
	return
}

func decodeBool(bs []byte) (v bool, remaining []byte, err error) {
	var u8 uint8
	if u8, remaining, err = decodeUint8(bs); err != nil {
		return
	}

	v = u8 == 1
	return
}
