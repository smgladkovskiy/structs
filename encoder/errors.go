package encoder

const invalidMarshalErrorMsg = "Invalid type %T provided to Marshal"

// InvalidUsagePooledEncoderError is a type representing an error returned
// when decoding is called on a still pooled Encoder
type InvalidUsagePooledEncoderError string

func (err InvalidUsagePooledEncoderError) Error() string {
	return string(err)
}

type InvalidMarshalError string

func (err InvalidMarshalError) Error() string {
	return string(err)
}
