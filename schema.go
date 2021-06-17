package enkodo

import (
	"fmt"
	"reflect"
)

var basicSchemas = map[reflect.Kind]Schema{
	reflect.String:  &basicSchema{efn: genericEncodeString, dfn: genericDecodeString},
	reflect.Bool:    &basicSchema{efn: genericEncodeBool, dfn: genericDecodeBool},
	reflect.Int:     &basicSchema{efn: genericEncodeInt, dfn: genericDecodeInt},
	reflect.Int8:    &basicSchema{efn: genericEncodeInt8, dfn: genericDecodeInt8},
	reflect.Int16:   &basicSchema{efn: genericEncodeInt16, dfn: genericDecodeInt16},
	reflect.Int32:   &basicSchema{efn: genericEncodeInt32, dfn: genericDecodeInt32},
	reflect.Int64:   &basicSchema{efn: genericEncodeInt64, dfn: genericDecodeInt64},
	reflect.Uint:    &basicSchema{efn: genericEncodeUint, dfn: genericDecodeUint},
	reflect.Uint8:   &basicSchema{efn: genericEncodeUint8, dfn: genericDecodeUint8},
	reflect.Uint16:  &basicSchema{efn: genericEncodeUint16, dfn: genericDecodeUint16},
	reflect.Uint32:  &basicSchema{efn: genericEncodeUint32, dfn: genericDecodeUint32},
	reflect.Uint64:  &basicSchema{efn: genericEncodeUint64, dfn: genericDecodeUint64},
	reflect.Float32: &basicSchema{efn: genericEncodeFloat32, dfn: genericDecodeFloat32},
	reflect.Float64: &basicSchema{efn: genericEncodeFloat64, dfn: genericDecodeFloat64},
}

// MakeSchema will initialize a new Schema
func MakeSchema(val interface{}) (s Schema, err error) {
	typ := reflect.TypeOf(val)
	return c.GetOrCreate(typ)
}

type Schema interface {
	MarshalEnkodo(enc *Encoder, val interface{}) error
	UnmarshalEnkodo(dec *Decoder, field *reflect.Value) error
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

type basicSchema struct {
	efn encoderFn
	dfn decoderFn
}

func (s *basicSchema) MarshalEnkodo(enc *Encoder, val interface{}) (err error) {
	return s.efn(enc, val)
}

func (s *basicSchema) UnmarshalEnkodo(dec *Decoder, val *reflect.Value) (err error) {
	return s.dfn(dec, val)
}
