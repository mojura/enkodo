package enkodo

import (
	"reflect"
	"sync"
)

var c = cache{m: make(map[reflect.Type]Schema)}

type cache struct {
	mux sync.RWMutex
	m   map[reflect.Type]Schema
}

func (c *cache) Get(t reflect.Type) (s Schema, ok bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	s, ok = c.m[t]
	return
}

func (c *cache) Create(t reflect.Type) (s Schema, err error) {
	if s, err = makeSchema(t); err != nil {
		return
	}

	c.mux.Lock()
	defer c.mux.Unlock()
	c.m[t] = s
	return
}

func (c *cache) GetOrCreate(t reflect.Type) (s Schema, err error) {
	var ok bool
	if s, ok = c.Get(t); ok {
		return
	}

	return c.Create(t)
}
