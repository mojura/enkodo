package enkodo

import (
	"fmt"
	"reflect"
)

func newStructSchema(rtype reflect.Type) (sp *structSchema, err error) {
	var s structSchema
	if s, err = makeStructSchema(rtype); err != nil {
		return
	}

	sp = &s
	return
}

func makeStructSchema(rtype reflect.Type) (s structSchema, err error) {
	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		compare := reflect.TypeOf((*Encodee)(nil)).Elem()
		if field.Type.Implements(compare) {
			s.encodeFns = append(s.encodeFns, genericEncode)
			s.decodeFns = append(s.decodeFns, genericDecode)
			continue
		}

		switch field.Type.Kind() {
		case reflect.String:
			s.encodeFns = append(s.encodeFns, genericEncodeString)
			s.decodeFns = append(s.decodeFns, genericDecodeString)
		case reflect.Int8:
			s.encodeFns = append(s.encodeFns, genericEncodeInt8)
			s.decodeFns = append(s.decodeFns, genericDecodeInt8)
		case reflect.Int16:
			s.encodeFns = append(s.encodeFns, genericEncodeInt16)
			s.decodeFns = append(s.decodeFns, genericDecodeInt16)
		case reflect.Int32:
			s.encodeFns = append(s.encodeFns, genericEncodeInt32)
			s.decodeFns = append(s.decodeFns, genericDecodeInt32)
		case reflect.Int64:
			s.encodeFns = append(s.encodeFns, genericEncodeInt64)
			s.decodeFns = append(s.decodeFns, genericDecodeInt64)
		case reflect.Int:
			s.encodeFns = append(s.encodeFns, genericEncodeInt)
			s.decodeFns = append(s.decodeFns, genericDecodeUint)

		case reflect.Uint8:
			s.encodeFns = append(s.encodeFns, genericEncodeUint8)
			s.decodeFns = append(s.decodeFns, genericDecodeUint8)
		case reflect.Uint16:
			s.encodeFns = append(s.encodeFns, genericEncodeUint16)
			s.decodeFns = append(s.decodeFns, genericDecodeUint16)
		case reflect.Uint32:
			s.encodeFns = append(s.encodeFns, genericEncodeUint32)
			s.decodeFns = append(s.decodeFns, genericDecodeUint32)
		case reflect.Uint64:
			s.encodeFns = append(s.encodeFns, genericEncodeUint64)
			s.decodeFns = append(s.decodeFns, genericDecodeUint64)
		case reflect.Uint:
			s.encodeFns = append(s.encodeFns, genericEncodeUint)
			s.decodeFns = append(s.decodeFns, genericDecodeUint)

		case reflect.Float32:
			s.encodeFns = append(s.encodeFns, genericEncodeFloat32)
			s.decodeFns = append(s.decodeFns, genericDecodeFloat32)
		case reflect.Float64:
			s.encodeFns = append(s.encodeFns, genericEncodeFloat64)
			s.decodeFns = append(s.decodeFns, genericDecodeFloat64)

		case reflect.Bool:
			s.encodeFns = append(s.encodeFns, genericEncodeBool)
			s.decodeFns = append(s.decodeFns, genericDecodeBool)

		case reflect.Struct:
			var ns structSchema
			// TODO: Create a cache for this
			if ns, err = makeStructSchema(field.Type); err != nil {
				return
			}

			s.encodeFns = append(s.encodeFns, ns.MarshalEnkodo)
			s.decodeFns = append(s.decodeFns, ns.UnmarshalEnkodo)

		case reflect.Slice:
			var ss basicSchema
			if ss, err = makeSliceSchema(field.Type); err != nil {
				return
			}

			s.encodeFns = append(s.encodeFns, ss.MarshalEnkodo)
			s.decodeFns = append(s.decodeFns, ss.UnmarshalEnkodo)

		case reflect.Map:
			var ms basicSchema
			if ms, err = makeMapSchema(field.Type); err != nil {
				return
			}

			s.encodeFns = append(s.encodeFns, ms.MarshalEnkodo)
			s.decodeFns = append(s.decodeFns, ms.UnmarshalEnkodo)

		default:
			err = fmt.Errorf("unsupported type provided, %s is not currently supported", field.Type.String())
			return
		}
	}

	return
}

type structSchema struct {
	encodeFns []encoderFn
	decodeFns []decoderFn
}

func (s *structSchema) MarshalEnkodo(enc *Encoder, val interface{}) (err error) {
	rval := reflect.ValueOf(val)
	for i, fn := range s.encodeFns {
		field := rval.Field(i)
		if err = fn(enc, field.Interface()); err != nil {
			return
		}
	}

	return
}

func (s *structSchema) UnmarshalEnkodo(dec *Decoder, field *reflect.Value) (err error) {
	target := *field
	if target.Kind() == reflect.Ptr {
		target = target.Elem()
	}

	for i, fn := range s.decodeFns {
		field := target.Field(i)
		if field.CanAddr() {
			field = field.Addr()
		}

		if err = fn(dec, &field); err != nil {
			return
		}
	}

	return
}
