package decoder

func (dec *Decoder) skipObject() (int, error) {
	var objectsOpen = 1
	var objectsClosed = 0
	for j := dec.cursor; j < dec.Length || dec.read(); j++ {
		switch dec.Data[j] {
		case '}':
			objectsClosed++
			// everything is closed return
			if objectsOpen == objectsClosed {
				// add char to object data
				return j + 1, nil
			}
		case '{':
			objectsOpen++
		case '"':
			j++
			var isInEscapeSeq bool
			var isFirstQuote = true
			for ; j < dec.Length || dec.read(); j++ {
				if dec.Data[j] != '"' {
					continue
				}
				if dec.Data[j-1] != '\\' || (!isInEscapeSeq && !isFirstQuote) {
					break
				} else {
					isInEscapeSeq = false
				}
				if isFirstQuote {
					isFirstQuote = false
				}
				// loop backward and count how many anti slash found
				// to see if string is effectively escaped
				ct := 0
				for i := j - 1; i > 0; i-- {
					if dec.Data[i] != '\\' {
						break
					}
					ct++
				}
				// is pair number of slashes, quote is not escaped
				if ct&1 == 0 {
					break
				}
				isInEscapeSeq = true
			}
		default:
			continue
		}
	}
	return 0, dec.raiseInvalidJSONErr(dec.cursor)
}
