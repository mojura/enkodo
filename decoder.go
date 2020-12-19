package mum

func newDecoder(bs []byte) *Decoder {
	var d Decoder
	if bs != nil {
		d.bs = bs
	}

	return &d
}

// Decoder helps to Marshal data
type Decoder struct {
	bs []byte
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
	v, d.bs, err = decodeUint8(d.bs)
	return
}

// Uint16 decodes a uint16 type
func (d *Decoder) Uint16() (v uint16, err error) {
	v, d.bs, err = decodeUint16(d.bs)
	return
}

// Uint32 decodes a uint32 type
func (d *Decoder) Uint32() (v uint32, err error) {
	v, d.bs, err = decodeUint32(d.bs)
	return
}

// Uint64 decodes a uint64 type
func (d *Decoder) Uint64() (v uint64, err error) {
	v, d.bs, err = decodeUint64(d.bs)
	return
}

// Int decodes an int type
func (d *Decoder) Int() (v int, err error) {
	v, d.bs, err = decodeInt(d.bs)
	return
}

// Int8 decodes an int8 type
func (d *Decoder) Int8() (v int8, err error) {
	v, d.bs, err = decodeInt8(d.bs)
	return
}

// Int16 decodes an int16 type
func (d *Decoder) Int16() (v int16, err error) {
	v, d.bs, err = decodeInt16(d.bs)
	return
}

// Int32 decodes an int32 type
func (d *Decoder) Int32() (v int32, err error) {
	v, d.bs, err = decodeInt32(d.bs)
	return
}

// Int64 decodes an int64 type
func (d *Decoder) Int64() (v int64, err error) {
	v, d.bs, err = decodeInt64(d.bs)
	return
}

// Float32 decodes a float64 type
func (d *Decoder) Float32() (v float32, err error) {
	v, d.bs, err = decodeFloat32(d.bs)
	return
}

// Float64 decodes a float64 type
func (d *Decoder) Float64() (v float64, err error) {
	v, d.bs, err = decodeFloat64(d.bs)
	return
}

// Bool will return a decoded boolean value
func (d *Decoder) Bool() (v bool, err error) {
	v, d.bs, err = decodeBool(d.bs)
	return
}

// Bytes will return decoded bytes
func (d *Decoder) Bytes() (v []byte, err error) {
	v, d.bs, err = decodeBytes(d.bs)
	return
}

// String will return a decoded string
func (d *Decoder) String() (v string, err error) {
	v, d.bs, err = decodeString(d.bs)
	return
}

// Decode will decode a decodee
func (d *Decoder) Decode(v Decodee) (err error) {
	return v.UnmarshalMum(d)
}

// Decodee is a data structure to be dedoded
type Decodee interface {
	UnmarshalMum(*Decoder) error
}
