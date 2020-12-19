package mum

import "errors"

var (
	// ErrEmptyBytes are returned when inbound bytes are empty during decode
	ErrEmptyBytes = errors.New("cannot decode, inbound bytes are empty")
	// ErrInvalidLength is returned when a byteslice has an invalid length for it's desired primitive
	ErrInvalidLength = errors.New("invalid length")
)

const (
	ceiling = 0x80
)
