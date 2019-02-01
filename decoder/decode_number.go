package decoder

var skipNumberEndCursorIncrement [256]int

func (dec *Decoder) skipNumber() (int, error) {
	end := dec.cursor + 1
	// look for following numbers
	for j := dec.cursor + 1; j < dec.Length || dec.read(); j++ {
		end += skipNumberEndCursorIncrement[dec.Data[j]]

		switch dec.Data[j] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', 'e', 'E', '+', '-', ' ', '\n', '\t', '\r':
			continue
		case ',', '}', ']':
			return end, nil
		default:
			// invalid json we expect numbers, dot (single one), comma, or spaces
			return end, dec.raiseInvalidJSONErr(dec.cursor)
		}
	}

	return end, nil
}
