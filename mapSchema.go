package enkodo

import "reflect"

func newMapSchema(rtype reflect.Type) (sp *basicSchema, err error) {
	var s basicSchema
	if s, err = makeMapSchema(rtype); err != nil {
		return
	}

	sp = &s
	return
}

func makeMapSchema(rtype reflect.Type) (s basicSchema, err error) {
	elem := rtype.Elem()
	if elem.Kind() == reflect.Interface {
		s.efn = genericEncodeMap
		s.dfn = genericDecodeMap
		return
	}

	var keyS, valS Schema
	if keyS, err = c.GetOrCreate(rtype.Key()); err != nil {
		return
	}

	if valS, err = c.GetOrCreate(elem); err != nil {
		return
	}

	s.efn = makeGenericMapEncoderFn(keyS, valS)
	s.dfn = makeGenericMapDecoderFn(keyS, valS)
	return
}
