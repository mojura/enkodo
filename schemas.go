package enkodo

import "reflect"

func newBasicSchema(val interface{}, efn encoderFn, dfn decoderFn) *basicSchema {
	var b basicSchema
	b.typ = reflect.TypeOf(val)
	b.efn = efn
	b.dfn = dfn
	return &b
}

type basicSchema struct {
	typ reflect.Type
	efn encoderFn
	dfn decoderFn
}

func (s *basicSchema) MarshalEnkodo(enc *Encoder, val interface{}) (err error) {
	return s.efn(enc, val)
}

func (s *basicSchema) UnmarshalEnkodo(dec *Decoder, val *reflect.Value) (err error) {
	return s.dfn(dec, val)
}
