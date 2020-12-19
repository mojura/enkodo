package mum

import "errors"

var (
	// ErrEmptyBytes are returned when inbound bytes are empty during decode
	ErrEmptyBytes = errors.New("cannot decode, inbound bytes are empty")
	// ErrInvalidLength is returned when a byteslice has an invalid length for it's desired primitive
	ErrInvalidLength = errors.New("invalid length")
	// ErrIsClosed is returned when an action is attempted on a closed instance
	ErrIsClosed = errors.New("cannot perform action on closed instance")
)

const (
	ceiling = 0x80
)
