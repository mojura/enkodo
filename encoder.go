package mum

func newEncoder(bs []byte) *Encoder {
	var e Encoder
	if bs != nil {
		e.bs = bs
	}

	return &e
}

// Encoder helps to Marshal data
type Encoder struct {
	bs []byte
}

// Uint encodes a uint type
func (e *Encoder) Uint(v uint) {
	e.bs = encodeUint(e.bs, v)
}

// Uint8 encodes a uint8 type
func (e *Encoder) Uint8(v uint8) {
	e.bs = encodeUint8(e.bs, v)
}

// Uint16 encodes a uint16 type
func (e *Encoder) Uint16(v uint16) {
	e.bs = encodeUint16(e.bs, v)
}

// Uint32 encodes a uint32 type
func (e *Encoder) Uint32(v uint32) {
	e.bs = encodeUint32(e.bs, v)
}

// Uint64 encodes a uint64 type
func (e *Encoder) Uint64(v uint64) {
	e.bs = encodeUint64(e.bs, v)
}

// Int encodes an int type
func (e *Encoder) Int(v int) {
	e.bs = encodeInt(e.bs, v)
}

// Int8 encodes an int8 type
func (e *Encoder) Int8(v int8) {
	e.bs = encodeInt8(e.bs, v)
}

// Int16 encodes an int16 type
func (e *Encoder) Int16(v int16) {
	e.bs = encodeInt16(e.bs, v)
}

// Int32 encodes an int32 type
func (e *Encoder) Int32(v int32) {
	e.bs = encodeInt32(e.bs, v)
}

// Int64 encodes an int64 type
func (e *Encoder) Int64(v int64) {
	e.bs = encodeInt64(e.bs, v)
}

// Float32 encodes an float32 type
func (e *Encoder) Float32(v float32) {
	e.bs = encodeFloat32(e.bs, v)
}

// Float64 encodes an float64 type
func (e *Encoder) Float64(v float64) {
	e.bs = encodeFloat64(e.bs, v)
}

// Bytes will encode a byteslice to the writer
func (e *Encoder) Bytes(v []byte) {
	e.bs = encodeBytes(e.bs, v)
}

// String will encode a string to the writer
func (e *Encoder) String(v string) {
	e.bs = encodeString(e.bs, v)
}

// Bool will encode a boolean value to the writer
func (e *Encoder) Bool(v bool) {
	e.bs = encodeBool(e.bs, v)
}

// Encode will encode an encodee
func (e *Encoder) Encode(v Encodee) {
	v.MarshalMum(e)
}

// Encodee is a data structure to be encoded
type Encodee interface {
	MarshalMum(*Encoder) error
}
