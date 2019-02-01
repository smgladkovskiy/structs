package decoder

import (
	"fmt"
)

const invalidJSONCharErrorMsg = "Invalid JSON, wrong char '%c' found at position %d"
const invalidUnmarshalErrorMsg = "Cannot unmarshal JSON to type '%T'"

// InvalidUnmarshalError is a type representing an error returned when
// Decoding cannot unmarshal JSON to the receiver type for various reasons.
type InvalidUnmarshalError string

// InvalidJSONError is a type representing an error returned when
// Decoding encounters invalid JSON.
type InvalidJSONError string

func (err InvalidJSONError) Error() string {
	return string(err)
}

func (err InvalidUnmarshalError) Error() string {
	return string(err)
}

func (dec *Decoder) makeInvalidUnmarshalErr(v interface{}) error {
	return InvalidUnmarshalError(
		fmt.Sprintf(
			invalidUnmarshalErrorMsg,
			v,
		),
	)
}

const invalidMarshalErrorMsg = "Invalid type %T provided to Marshal"

func (dec *Decoder) raiseInvalidJSONErr(pos int) error {
	var c byte
	if len(dec.Data) > pos {
		c = dec.Data[pos]
	}
	dec.Err = InvalidJSONError(
		fmt.Sprintf(
			invalidJSONCharErrorMsg,
			c,
			pos,
		),
	)
	return dec.Err
}
