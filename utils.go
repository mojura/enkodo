package enkodo

import (
	"bytes"
	"fmt"
	"io"
	"unsafe"
)

const notEnoughBytesLayout = "not enough bytes available to decode <%T>, needed %d and has an available %d"

func getStringBytes(str *string) *[]byte {
	return ((*[]byte)(unsafe.Pointer(str)))
}

func getStringFromBytes(bs []byte) string {
	return *((*string)(unsafe.Pointer(&bs)))
}

// Marshal will encode a value
func Marshal(v Encodee) (bs []byte, err error) {
	return MarshalAppend(v, nil)
}

// MarshalAppend will encode a value to a provided slice
func MarshalAppend(v Encodee, buffer []byte) (bs []byte, err error) {
	enc := newEncoder(nil)
	enc.bs = buffer
	if err = enc.Encode(v); err != nil {
		return
	}

	bs = enc.bs
	return
}

// Unmarshal will decode a value
func Unmarshal(bs []byte, v Decodee) (err error) {
	dec := newDecoder(bytes.NewReader(bs))
	return dec.Decode(v)
}

func newNotEnoughBytesError(target interface{}, needed, remaining int) (err error) {
	err = fmt.Errorf(notEnoughBytesLayout, target, needed, remaining)
	return
}

type reader interface {
	io.Reader
	io.ByteReader
}

func expandSlice(bs *[]byte, sz int) {
	if *bs != nil && cap(*bs) >= sz {
		*bs = (*bs)[:sz]
		return
	}

	*bs = make([]byte, sz)
}
