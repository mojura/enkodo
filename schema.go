package enkodo

import (
	"fmt"
	"reflect"
)

var basicSchemas = map[reflect.Kind]*basicSchema{
	reflect.String:  newBasicSchema("", genericEncodeString, genericDecodeString),
	reflect.Bool:    newBasicSchema(true, genericEncodeBool, genericDecodeBool),
	reflect.Int:     newBasicSchema(int(1), genericEncodeInt, genericDecodeInt),
	reflect.Int8:    newBasicSchema(int8(1), genericEncodeInt8, genericDecodeInt8),
	reflect.Int16:   newBasicSchema(int16(1), genericEncodeInt16, genericDecodeInt16),
	reflect.Int32:   newBasicSchema(int32(1), genericEncodeInt32, genericDecodeInt32),
	reflect.Int64:   newBasicSchema(int64(1), genericEncodeInt64, genericDecodeInt64),
	reflect.Uint:    newBasicSchema(uint(1), genericEncodeUint, genericDecodeUint),
	reflect.Uint8:   newBasicSchema(uint8(1), genericEncodeUint8, genericDecodeUint8),
	reflect.Uint16:  newBasicSchema(uint16(1), genericEncodeUint16, genericDecodeUint16),
	reflect.Uint32:  newBasicSchema(uint32(1), genericEncodeUint32, genericDecodeUint32),
	reflect.Uint64:  newBasicSchema(uint64(1), genericEncodeUint64, genericDecodeUint64),
	reflect.Float32: newBasicSchema(float32(1), genericEncodeFloat32, genericDecodeFloat32),
	reflect.Float64: newBasicSchema(float64(1), genericEncodeFloat64, genericDecodeFloat64),
}

// MakeSchema will initialize a new Schema
func MakeSchema(val interface{}) (s Schema, err error) {
	typ := reflect.TypeOf(val)
	return c.GetOrCreate(typ)
}

func makeSchema(rtype reflect.Type) (s Schema, err error) {
	var ok bool
	kind := rtype.Kind()
	if s, ok = basicSchemas[kind]; ok {
		return
	}

	switch kind {
	case reflect.Struct:
		return newStructSchema(rtype)
	case reflect.Slice:
		return newSliceSchema(rtype)
	case reflect.Map:
		return newMapSchema(rtype)

	// TODO: Implement this
	//case reflect.Array:
	//	return

	default:
		err = fmt.Errorf("type of %s is not supported", rtype.Kind())
		return
	}
}

type Schema interface {
	MarshalEnkodo(enc *Encoder, val interface{}) error
	UnmarshalEnkodo(dec *Decoder, field *reflect.Value) error
}
