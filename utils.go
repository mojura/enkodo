package mum

import "unsafe"

func getStringBytes(str string) []byte {
	return *((*[]byte)(unsafe.Pointer(&str)))
}

func getStringFromBytes(bs []byte) string {
	return *((*string)(unsafe.Pointer(&bs)))
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
