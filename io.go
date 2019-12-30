package readerer

import "io"

// ReaderFromReaderAt implements io.Reader interface.
type ReaderFromReaderAt struct {
	r      io.ReaderAt
	offset int64
}

// FromReaderAt convert io.ReaderAt to io.Reader. The offset is the start position
// of the ReaderFromReaderAt
func FromReaderAt(ra io.ReaderAt, offset int64) *ReaderFromReaderAt {
	return &ReaderFromReaderAt{r: ra, offset: offset}
}

// Read is the implementation of io.Reader.
func (r *ReaderFromReaderAt) Read(p []byte) (n int, err error) {
	return r.r.ReadAt(p, r.offset)
}
