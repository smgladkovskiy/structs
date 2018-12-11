package decoder

func (dec *Decoder) DecodeBool(v *bool) error {
	for ; dec.cursor < dec.Length || dec.read(); dec.cursor++ {
		switch dec.Data[dec.cursor] {
		case ' ', '\n', '\t', '\r', ',':
			continue
		case 't':
			dec.cursor++
			err := dec.assertTrue()
			if err != nil {
				return err
			}
			*v = true
			return nil
		case 'f':
			dec.cursor++
			err := dec.assertFalse()
			if err != nil {
				return err
			}
			*v = false
			return nil
		default:
			dec.Err = dec.makeInvalidUnmarshalErr(v)
			err := dec.skipData()
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

func (dec *Decoder) assertTrue() error {
	i := 0
	for ; dec.cursor < dec.Length || dec.read(); dec.cursor++ {
		switch i {
		case 0:
			if dec.Data[dec.cursor] != 'r' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 1:
			if dec.Data[dec.cursor] != 'u' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 2:
			if dec.Data[dec.cursor] != 'e' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 3:
			switch dec.Data[dec.cursor] {
			case ' ', '\b', '\t', '\n', ',', ']', '}':
				// dec.cursor--
				return nil
			default:
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		}
		i++
	}
	if i == 3 {
		return nil
	}
	return dec.raiseInvalidJSONErr(dec.cursor)
}
func (dec *Decoder) skipData() error {
	for ; dec.cursor < dec.Length || dec.read(); dec.cursor++ {
		switch dec.Data[dec.cursor] {
		case ' ', '\n', '\t', '\r', ',':
			continue
		// is null
		case 'n':
			dec.cursor++
			err := dec.assertNull()
			if err != nil {
				return err
			}
			return nil
		case 't':
			dec.cursor++
			err := dec.assertTrue()
			if err != nil {
				return err
			}
			return nil
		// is false
		case 'f':
			dec.cursor++
			err := dec.assertFalse()
			if err != nil {
				return err
			}
			return nil
		// is an object
		case '{':
			dec.cursor = dec.cursor + 1
			end, err := dec.skipObject()
			dec.cursor = end
			return err
		// is string
		case '"':
			dec.cursor = dec.cursor + 1
			err := dec.skipString()
			return err
		// is array
		case '[':
			dec.cursor = dec.cursor + 1
			end, err := dec.skipArray()
			dec.cursor = end
			return err
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
			end, err := dec.skipNumber()
			dec.cursor = end
			return err
		}
		return dec.raiseInvalidJSONErr(dec.cursor)
	}
	return dec.raiseInvalidJSONErr(dec.cursor)
}

func (dec *Decoder) assertFalse() error {
	i := 0
	for ; dec.cursor < dec.Length || dec.read(); dec.cursor++ {
		switch i {
		case 0:
			if dec.Data[dec.cursor] != 'a' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 1:
			if dec.Data[dec.cursor] != 'l' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 2:
			if dec.Data[dec.cursor] != 's' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 3:
			if dec.Data[dec.cursor] != 'e' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 4:
			switch dec.Data[dec.cursor] {
			case ' ', '\t', '\n', ',', ']', '}':
				// dec.cursor--
				return nil
			default:
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		}
		i++
	}
	if i == 4 {
		return nil
	}
	return dec.raiseInvalidJSONErr(dec.cursor)
}
