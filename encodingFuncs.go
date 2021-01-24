package enkodo

import (
	"math"
	"unsafe"
)

func encodeUint(bs []byte, v uint) (out []byte) {
	return encodeUint64(bs, uint64(v))
}

func encodeUint8(bs []byte, v uint8) (out []byte) {
	return append(bs, v)
}

func encodeUint16(bs []byte, v uint16) (out []byte) {
	return encodeUint64(bs, uint64(v))
}

func encodeUint32(bs []byte, v uint32) (out []byte) {
	return encodeUint64(bs, uint64(v))
}

func encodeUint64(in []byte, v uint64) (buf []byte) {
	switch {
	case v < 1<<7-1:
		return append(in, byte(v))
	case v < 1<<14-1:
		return append(in, byte(v)|0x80, byte(v>>7))
	case v < 1<<21-1:
		return append(in, byte(v)|0x80, byte(v>>7)|0x80, byte(v>>14))
	case v < 1<<28-1:
		return append(in, byte(v)|0x80, byte(v>>7)|0x80, byte(v>>14)|0x80, byte(v>>21))
	case v < 1<<35-1:
		return append(in, byte(v)|0x80, byte(v>>7)|0x80, byte(v>>14)|0x80, byte(v>>21)|0x80, byte(v>>28))
	case v < 1<<42-1:
		return append(in, byte(v)|0x80, byte(v>>7)|0x80, byte(v>>14)|0x80, byte(v>>21)|0x80, byte(v>>28)|0x80, byte(v>>35))
	case v < 1<<49-1:
		return append(in, byte(v)|0x80, byte(v>>7)|0x80, byte(v>>14)|0x80, byte(v>>21)|0x80, byte(v>>28)|0x80, byte(v>>35)|0x80, byte(v>>42))
	case v < 1<<56-1:
		return append(in, byte(v)|0x80, byte(v>>7)|0x80, byte(v>>14)|0x80, byte(v>>21)|0x80, byte(v>>28)|0x80, byte(v>>35)|0x80, byte(v>>42)|0x80, byte(v>>49))
	default:
		return append(in, byte(v)|0x80, byte(v>>7)|0x80, byte(v>>14)|0x80, byte(v>>21)|0x80, byte(v>>28)|0x80, byte(v>>35)|0x80, byte(v>>42)|0x80, byte(v>>49)|0x80, byte(v>>56))
	}
}

func encodeInt(bs []byte, v int) (out []byte) {
	return encodeInt64(bs, int64(v))
}

func encodeInt8(bs []byte, v int8) (out []byte) {
	return encodeUint8(bs, *(*uint8)(unsafe.Pointer(&v)))
}

func encodeInt16(bs []byte, v int16) (out []byte) {
	return encodeInt64(bs, int64(v))
}

func encodeInt32(bs []byte, v int32) (out []byte) {
	return encodeInt64(bs, int64(v))
}

func encodeInt64(bs []byte, v int64) (out []byte) {
	return encodeUint64(bs, *(*uint64)(unsafe.Pointer(&v)))
}

func encodeFloat32(bs []byte, v float32) (out []byte) {
	return encodeUint32(bs, math.Float32bits(v))
}

func encodeFloat64(bs []byte, v float64) (out []byte) {
	return encodeUint64(bs, math.Float64bits(v))
}

func encodeBytes(bs, v []byte) (out []byte) {
	out = encodeInt(bs, len(v))
	out = append(out, v...)
	return
}

func encodeString(bs []byte, v string) (out []byte) {
	bsp := getStringBytes(&v)
	return encodeBytes(bs, *bsp)
}

func encodeBool(bs []byte, v bool) (out []byte) {
	if v {
		return encodeUint8(bs, 1)
	}

	return encodeUint8(bs, 0)
}
