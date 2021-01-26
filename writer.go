package enkodo

import "io"

// NewWriter will initialize a new instance of writer
func NewWriter(in io.Writer) *Writer {
	var w Writer
	w.e = newEncoder(in)
	return &w
}

// Writer manages the writing of enkodo output
type Writer struct {
	e *Encoder
}

// Encode will encode an encodee
func (w *Writer) Encode(v Encodee) (err error) {
	if w.e == nil {
		return ErrIsClosed
	}

	return v.MarshalEnkodo(w.e)
}

// Reset will reset the underlying bytes of the Encoder
func (w *Writer) Reset() {
	if w.e == nil {
		return
	}

	w.e.bs = w.e.bs[:0]
}

// WriteTo will write to an io.Writer
func (w *Writer) WriteTo(dest io.Writer) (n int64, err error) {
	if w.e == nil {
		err = ErrIsClosed
		return
	}

	var written int
	if written, err = dest.Write(w.Bytes()); err != nil {
		return
	}

	n = int64(written)
	w.Reset()
	return
}

// Bytes will expose the underlying bytes
func (w *Writer) Bytes() []byte {
	return w.e.bs
}

// Close will close the writer
func (w *Writer) Close() (err error) {
	if w.e == nil {
		return ErrIsClosed
	}

	w.e.teardown()
	w.e = nil
	return
}
