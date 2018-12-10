package decoder

import (
	"io"
)

// A Decoder reads and decodes JSON values from an input stream.
type Decoder struct {
	r        io.Reader
	Data     []byte
	Err      error
	isPooled byte
	called   byte
	child    byte
	cursor   int
	Length   int
	keysDone int
}
