package enkodo

import (
	"fmt"
	"reflect"
)

var (
	iface interface{}

	genericInterface = reflect.TypeOf(iface)
	genericSlice     = reflect.TypeOf([]interface{}{})
	genericMap       = reflect.TypeOf(map[interface{}]interface{}{})
)

func genericEncode(enc *Encoder, val interface{}) (err error) {
	return enc.Encode(val.(Encodee))
}

func genericEncodeInterface(enc *Encoder, val interface{}) (err error) {
	typ := reflect.TypeOf(val)
	if err = enc.Uint(uint(typ.Kind())); err != nil {
		return
	}

	var s Schema
	if s, err = c.GetOrCreate(typ); err != nil {
		return
	}

	return s.MarshalEnkodo(enc, val)
}

func genericEncodeString(enc *Encoder, val interface{}) (err error) {
	return enc.String(val.(string))
}

func genericEncodeBool(enc *Encoder, val interface{}) (err error) {
	return enc.Bool(val.(bool))
}

func genericEncodeInt8(enc *Encoder, val interface{}) (err error) {
	return enc.Int8(val.(int8))
}

func genericEncodeInt16(enc *Encoder, val interface{}) (err error) {
	return enc.Int16(val.(int16))
}

func genericEncodeInt32(enc *Encoder, val interface{}) (err error) {
	return enc.Int32(val.(int32))
}

func genericEncodeInt64(enc *Encoder, val interface{}) (err error) {
	return enc.Int64(val.(int64))
}

func genericEncodeInt(enc *Encoder, val interface{}) (err error) {
	return enc.Int(val.(int))
}

func genericEncodeUint8(enc *Encoder, val interface{}) (err error) {
	return enc.Uint8(val.(uint8))
}

func genericEncodeUint16(enc *Encoder, val interface{}) (err error) {
	return enc.Uint16(val.(uint16))
}

func genericEncodeUint32(enc *Encoder, val interface{}) (err error) {
	return enc.Uint32(val.(uint32))
}

func genericEncodeUint64(enc *Encoder, val interface{}) (err error) {
	return enc.Uint64(val.(uint64))
}

func genericEncodeUint(enc *Encoder, val interface{}) (err error) {
	return enc.Uint(val.(uint))
}

func genericEncodeFloat32(enc *Encoder, val interface{}) (err error) {
	return enc.Float32(val.(float32))
}

func genericEncodeFloat64(enc *Encoder, val interface{}) (err error) {
	return enc.Float64(val.(float64))
}

func genericDecode(dec *Decoder, field *reflect.Value) (err error) {
	return field.
		Interface().(Decodee).
		UnmarshalEnkodo(dec)
}

func genericDecodeInterface(dec *Decoder, val *reflect.Value) (err error) {
	var kindU uint
	if kindU, err = dec.Uint(); err != nil {
		return
	}

	var (
		s  Schema
		ok bool
	)

	kind := reflect.Kind(kindU)
	if s, ok = basicSchemas[kind]; ok {
		return s.UnmarshalEnkodo(dec, val)
	}

	switch kind {
	case reflect.Slice:
		s, err = newSliceSchema(genericSlice)
	case reflect.Map:
		s, err = newMapSchema(genericMap)

	// Note: This needs to be a slow struct type (which acts like a map)
	//case reflect.Struct:
	//	s, err = newStructSchema(val.Type())

	default:
		return fmt.Errorf("cannot decode kind of <%s>", kind.String())
	}

	if err != nil {
		return
	}

	return s.UnmarshalEnkodo(dec, val)
}

func genericDecodeString(dec *Decoder, field *reflect.Value) (err error) {
	var v string
	if v, err = dec.String(); err != nil {
		return
	}

	indirect := reflect.Indirect(*field)
	indirect.Set(reflect.ValueOf(v))
	return
}

func genericDecodeBool(dec *Decoder, field *reflect.Value) (err error) {
	var v bool
	if v, err = dec.Bool(); err != nil {
		return
	}

	indirect := reflect.Indirect(*field)
	indirect.Set(reflect.ValueOf(v))
	return
}

func genericDecodeInt8(dec *Decoder, field *reflect.Value) (err error) {
	var v int8
	if v, err = dec.Int8(); err != nil {
		return
	}

	indirect := reflect.Indirect(*field)
	indirect.Set(reflect.ValueOf(v))
	return
}

func genericDecodeInt16(dec *Decoder, field *reflect.Value) (err error) {
	var v int16
	if v, err = dec.Int16(); err != nil {
		return
	}

	indirect := reflect.Indirect(*field)
	indirect.Set(reflect.ValueOf(v))
	return
}

func genericDecodeInt32(dec *Decoder, field *reflect.Value) (err error) {
	var v int32
	if v, err = dec.Int32(); err != nil {
		return
	}

	indirect := reflect.Indirect(*field)
	indirect.Set(reflect.ValueOf(v))
	return
}

func genericDecodeInt64(dec *Decoder, field *reflect.Value) (err error) {
	var v int64
	if v, err = dec.Int64(); err != nil {
		return
	}

	indirect := reflect.Indirect(*field)
	indirect.Set(reflect.ValueOf(v))
	return
}

func genericDecodeInt(dec *Decoder, field *reflect.Value) (err error) {
	var v int
	if v, err = dec.Int(); err != nil {
		return
	}

	indirect := reflect.Indirect(*field)
	indirect.Set(reflect.ValueOf(v))
	return
}

func genericDecodeUint8(dec *Decoder, field *reflect.Value) (err error) {
	var v uint8
	if v, err = dec.Uint8(); err != nil {
		return
	}

	indirect := reflect.Indirect(*field)
	indirect.Set(reflect.ValueOf(v))
	return
}

func genericDecodeUint16(dec *Decoder, field *reflect.Value) (err error) {
	var v uint16
	if v, err = dec.Uint16(); err != nil {
		return
	}

	indirect := reflect.Indirect(*field)
	indirect.Set(reflect.ValueOf(v))
	return
}

func genericDecodeUint32(dec *Decoder, field *reflect.Value) (err error) {
	var v uint32
	if v, err = dec.Uint32(); err != nil {
		return
	}

	indirect := reflect.Indirect(*field)
	indirect.Set(reflect.ValueOf(v))
	return
}

func genericDecodeUint64(dec *Decoder, field *reflect.Value) (err error) {
	var v uint64
	if v, err = dec.Uint64(); err != nil {
		return
	}

	indirect := reflect.Indirect(*field)
	indirect.Set(reflect.ValueOf(v))
	return
}

func genericDecodeUint(dec *Decoder, field *reflect.Value) (err error) {
	var v uint
	if v, err = dec.Uint(); err != nil {
		return
	}

	indirect := reflect.Indirect(*field)
	indirect.Set(reflect.ValueOf(v))
	return
}

func genericDecodeFloat32(dec *Decoder, field *reflect.Value) (err error) {
	var v float32
	if v, err = dec.Float32(); err != nil {
		return
	}

	indirect := reflect.Indirect(*field)
	indirect.Set(reflect.ValueOf(v))
	return
}

func genericDecodeFloat64(dec *Decoder, field *reflect.Value) (err error) {
	var v float64
	if v, err = dec.Float64(); err != nil {
		return
	}

	indirect := reflect.Indirect(*field)
	indirect.Set(reflect.ValueOf(v))
	return
}

func genericEncodeSlice(enc *Encoder, val interface{}) (err error) {
	s := val.([]interface{})
	if err = enc.Int(len(s)); err != nil {
		return
	}

	for _, val := range s {
		if err = genericEncodeInterface(enc, val); err != nil {
			return
		}
	}

	return
}

func genericDecodeSlice(dec *Decoder, val *reflect.Value) (err error) {
	var limit int
	if limit, err = dec.Int(); err != nil {
		return
	}

	fmt.Println("Generic decode slice!", val)
	s := make([]interface{}, 0, limit)
	for i := 0; i < limit; i++ {
		var val interface{}
		rval := reflect.ValueOf(&val)
		if err = genericDecodeInterface(dec, &rval); err != nil {
			return
		}

		switch n := rval.Interface().(type) {
		case *string:
			fmt.Printf("Str: <%s> / %v\n", *n, val)
		}

		s = append(s, val)
	}

	fmt.Println("S?", s)

	reflect.Indirect(*val).
		Set(reflect.ValueOf(s))
	return
}

func genericEncodeMap(enc *Encoder, val interface{}) (err error) {
	fmt.Println("Oh?", val)
	rval := reflect.ValueOf(val)
	if err = enc.Int(rval.Len()); err != nil {
		return
	}

	keys := rval.MapKeys()
	for _, key := range keys {
		if err = genericEncodeInterface(enc, key.Interface()); err != nil {
			return
		}

		entry := rval.MapIndex(key)
		if err = genericEncodeInterface(enc, entry.Interface()); err != nil {
			return
		}
	}

	return
}

func genericDecodeMap(dec *Decoder, ref *reflect.Value) (err error) {
	var limit int
	if limit, err = dec.Int(); err != nil {
		return
	}

	*ref = reflect.MakeMapWithSize(ref.Elem().Type(), limit)
	fmt.Println("Target?", ref.Interface())
	for i := 0; i < limit; i++ {
		var key interface{}
		rkey := reflect.ValueOf(key)
		if err = genericDecodeInterface(dec, &rkey); err != nil {
			return
		}

		var val interface{}
		rval := reflect.ValueOf(val)
		if err = genericDecodeInterface(dec, &rval); err != nil {
			return
		}

		ref.SetMapIndex(rkey.Elem(), rval.Elem())
	}

	return
}

func makeGenericSliceEncoderFn(s Schema) (efn encoderFn) {
	return func(enc *Encoder, val interface{}) (err error) {
		rval := reflect.ValueOf(val)
		limit := rval.Len()
		if err = enc.Int(limit); err != nil {
			return
		}

		for i := 0; i < limit; i++ {
			item := rval.Index(i)
			if err = s.MarshalEnkodo(enc, item.Interface()); err != nil {
				return
			}
		}

		return
	}
}

func makeGenericSliceDecoderFn(s Schema) (dfn decoderFn) {
	return func(dec *Decoder, target *reflect.Value) (err error) {
		var limit int
		if limit, err = dec.Int(); err != nil {
			return
		}

		if target.IsNil() {
			sliceRef := reflect.MakeSlice(target.Type(), 0, limit)
			target.Set(sliceRef)
		}

		slice := target.Elem()
		typ := slice.Type().Elem()

		for i := 0; i < limit; i++ {
			fieldVal := reflect.New(typ)
			if err = s.UnmarshalEnkodo(dec, &fieldVal); err != nil {
				return
			}

			slice = reflect.Append(slice, reflect.Indirect(fieldVal))
		}

		reflect.Indirect(*target).Set(slice)
		return
	}
}

func makeGenericMapEncoderFn(keyS, valS Schema) (efn encoderFn) {
	return func(enc *Encoder, val interface{}) (err error) {
		rval := reflect.Indirect(reflect.ValueOf(val))
		keys := rval.MapKeys()
		if err = enc.Int(len(keys)); err != nil {
			return
		}

		for i := 0; i < len(keys); i++ {
			key := keys[i].Interface()
			val := rval.MapIndex(keys[i]).Interface()
			if err = keyS.MarshalEnkodo(enc, key); err != nil {
				return
			}

			if err = valS.MarshalEnkodo(enc, val); err != nil {
				return
			}
		}

		return
	}
}

func makeGenericMapDecoderFn(keyS, valS Schema) (dfn decoderFn) {
	return func(dec *Decoder, ref *reflect.Value) (err error) {
		var limit int
		if limit, err = dec.Int(); err != nil {
			return
		}

		target := ref.Elem()
		typ := target.Type()
		if target.IsNil() {
			target = reflect.MakeMap(typ)
			reflect.Indirect(*ref).Set(target)
		}

		keyTyp := typ.Key()
		valTyp := typ.Elem()

		for i := 0; i < limit; i++ {
			key := reflect.New(keyTyp)
			val := reflect.New(valTyp)

			if err = keyS.UnmarshalEnkodo(dec, &key); err != nil {
				return
			}

			if err = valS.UnmarshalEnkodo(dec, &val); err != nil {
				return
			}

			key = reflect.Indirect(key)
			val = reflect.Indirect(val)
			target.SetMapIndex(key, val)
		}

		return
	}
}
