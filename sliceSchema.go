package enkodo

import (
	"reflect"
)

func newSliceSchema(rtype reflect.Type) (sp *basicSchema, err error) {
	var s basicSchema
	if s, err = makeSliceSchema(rtype); err != nil {
		return
	}

	sp = &s
	return
}

func makeSliceSchema(rtype reflect.Type) (s basicSchema, err error) {
	typ := rtype.Elem()
	if typ.Kind() == reflect.Interface {
		s.efn = genericEncodeSlice
		s.dfn = genericDecodeSlice
		return
	}

	var valS Schema
	if valS, err = c.GetOrCreate(typ); err != nil {
		return
	}

	s.efn = makeGenericSliceEncoderFn(valS)
	s.dfn = makeGenericSliceDecoderFn(valS)
	return
}
