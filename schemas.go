package enkodo

import "reflect"

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
