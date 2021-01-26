package enkodo

import "io"

// NewReader will initialize a new instance of writer
func NewReader(in io.Reader) *Reader {
	var r Reader
	r.d = newDecoder(in)
	return &r
}

// Reader manages the writing of enkodo output
type Reader struct {
	d *Decoder
}

// Decode will decode an decodee
func (r *Reader) Decode(v Decodee) (err error) {
	if r.d == nil {
		return ErrIsClosed
	}

	return v.UnmarshalEnkodo(r.d)
}

// Close will close the reader
func (r *Reader) Close() (err error) {
	r.d = nil
	return
}
