package encoder

import (
	"fmt"
)

// Encode encodes a value to JSON.
//
// If Encode cannot find a way to encode the type to JSON
// it will return an InvalidMarshalError.
func (enc *Encoder) Encode(v interface{}) error {
	// if enc.isPooled == 1 {
	// 	panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	// }
	switch vt := v.(type) {
	// case string:
	// 	return enc.EncodeString(vt)
	// case bool:
	// 	return enc.EncodeBool(vt)
	// case MarshalerJSONArray:
	// 	return enc.EncodeArray(vt)
	// case MarshalerJSONObject:
	// 	return enc.EncodeObject(vt)
	case int:
		return enc.EncodeInt(vt)
	case int64:
		return enc.EncodeInt64(vt)
	case int32:
		return enc.EncodeInt(int(vt))
	case int8:
		return enc.EncodeInt(int(vt))
	case uint64:
		return enc.EncodeUint64(vt)
	case uint32:
		return enc.EncodeInt(int(vt))
	case uint16:
		return enc.EncodeInt(int(vt))
	case uint8:
		return enc.EncodeInt(int(vt))
	case float64:
		return enc.EncodeFloat(vt)
	case float32:
		return enc.EncodeFloat32(vt)
	// case *EmbeddedJSON:
	// 	return enc.EncodeEmbeddedJSON(vt)
	default:
		return InvalidMarshalError(fmt.Sprintf(invalidMarshalErrorMsg, vt))
	}
}
