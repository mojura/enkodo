package mum

// NewReader will initialize a new instance of writer
func NewReader(buffer []byte) *Reader {
	var r Reader
	r.d = newDecoder(buffer)
	return &r
}

// Reader manages the writing of mum output
type Reader struct {
	d *Decoder
}

// Decode will decode an decodee
func (r *Reader) Decode(v Decodee) (err error) {
	if r.d == nil {
		return ErrIsClosed
	}

	v.UnmarshalMum(r.d)
	return
}

// Close will close the reader
func (r *Reader) Close() (err error) {
	r.d = nil
	return
}
