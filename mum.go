package mum

import ()

// Type represents a primitive type
type Type uint8

const (
	// Nil is the zero-value for the Types block
	Nil Type = iota
	// UInt8 represents uint8
	UInt8
	// UInt16 represents uint16
	UInt16
	// UInt32 represents uint32
	UInt32
	// UInt64 represents uin64
	UInt64
	// Int8 represents int8
	Int8
	// Int16 represents int16
	Int16
	// Int32 represents int32
	Int32
	// Int64 represents int64
	Int64
	// Bytes represents a byteslice
	Bytes
	// String represents a string
	String
)

type n16 [2]byte
type n32 [4]byte
type n64 [8]byte
