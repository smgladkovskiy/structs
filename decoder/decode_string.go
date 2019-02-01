package decoder

import (
	"io"
	"unicode/utf16"
	"unicode/utf8"
	"unsafe"
)

func (dec *Decoder) DecodeString(v *string) error {
	for ; dec.cursor < dec.Length || dec.read(); dec.cursor++ {
		switch dec.Data[dec.cursor] {
		case ' ', '\n', '\t', '\r', ',':
			// is string
			continue
		case '"':
			dec.cursor++
			start, end, err := dec.getString()
			if err != nil {
				return err
			}
			// we do minus one to remove the last quote
			d := dec.Data[start : end-1]
			*v = *(*string)(unsafe.Pointer(&d))
			dec.cursor = end
			return nil
		// is nil
		case 'n':
			dec.cursor++
			err := dec.assertNull()
			if err != nil {
				return err
			}
			return nil
		default:
			dec.Err = dec.makeInvalidUnmarshalErr(v)
			err := dec.skipString()
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

func (dec *Decoder) assertNull() error {
	i := 0
	for ; dec.cursor < dec.Length || dec.read(); dec.cursor++ {
		switch i {
		case 0:
			if dec.Data[dec.cursor] != 'u' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 1:
			if dec.Data[dec.cursor] != 'l' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 2:
			if dec.Data[dec.cursor] != 'l' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 3:
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
	if i == 3 {
		return nil
	}
	return dec.raiseInvalidJSONErr(dec.cursor)
}

func (dec *Decoder) read() bool {
	if dec.r != nil {
		// if we reach the end, double the buffer to ensure there's always more space
		if len(dec.Data) == dec.Length {
			nLen := dec.Length * 2
			Buf := make([]byte, nLen, nLen)
			copy(Buf, dec.Data)
			dec.Data = Buf
		}
		var n int
		var err error
		for n == 0 {
			n, err = dec.r.Read(dec.Data[dec.Length:])
			if err != nil {
				if err != io.EOF {
					dec.Err = err
					return false
				}
				if n == 0 {
					return false
				}
				dec.Length = dec.Length + n
				return true
			}
		}
		dec.Length = dec.Length + n
		return true
	}
	return false
}

func (dec *Decoder) getString() (int, int, error) {
	// extract key
	var keyStart = dec.cursor
	// var str *Builder
	for dec.cursor < dec.Length || dec.read() {
		switch dec.Data[dec.cursor] {
		// string found
		case '"':
			dec.cursor = dec.cursor + 1
			return keyStart, dec.cursor, nil
		// slash found
		case '\\':
			dec.cursor = dec.cursor + 1
			err := dec.parseEscapedString()
			if err != nil {
				return 0, 0, err
			}
		default:
			dec.cursor = dec.cursor + 1
			continue
		}
	}
	return 0, 0, dec.raiseInvalidJSONErr(dec.cursor)
}

func (dec *Decoder) parseEscapedString() error {
	if dec.cursor >= dec.Length && !dec.read() {
		return dec.raiseInvalidJSONErr(dec.cursor)
	}
	switch dec.Data[dec.cursor] {
	case '"':
		dec.Data[dec.cursor] = '"'
	case '\\':
		dec.Data[dec.cursor] = '\\'
	case '/':
		dec.Data[dec.cursor] = '/'
	case 'b':
		dec.Data[dec.cursor] = '\b'
	case 'f':
		dec.Data[dec.cursor] = '\f'
	case 'n':
		dec.Data[dec.cursor] = '\n'
	case 'r':
		dec.Data[dec.cursor] = '\r'
	case 't':
		dec.Data[dec.cursor] = '\t'
	case 'u':
		start := dec.cursor
		dec.cursor++
		str, err := dec.parseUnicode()
		if err != nil {
			return err
		}
		diff := dec.cursor - start
		dec.Data = append(append(dec.Data[:start-1], str...), dec.Data[dec.cursor:]...)
		dec.Length = len(dec.Data)
		dec.cursor += len(str) - diff - 1

		return nil
	default:
		return dec.raiseInvalidJSONErr(dec.cursor)
	}

	dec.Data = append(dec.Data[:dec.cursor-1], dec.Data[dec.cursor:]...)
	dec.Length--

	// Since we've lost a character, our dec.cursor offset is now
	// 1 past the escaped character which is precisely where we
	// want it.

	return nil
}

func (dec *Decoder) parseUnicode() ([]byte, error) {
	// get unicode after u
	r, err := dec.getUnicode()
	if err != nil {
		return nil, err
	}
	// no error start making new string
	str := make([]byte, 16, 16)
	i := 0
	// check if code can be a surrogate utf16
	if utf16.IsSurrogate(r) {
		if dec.cursor >= dec.Length && !dec.read() {
			return nil, dec.raiseInvalidJSONErr(dec.cursor)
		}
		c := dec.Data[dec.cursor]
		if c != '\\' {
			i += utf8.EncodeRune(str, r)
			return str[:i], nil
		}
		dec.cursor++
		if dec.cursor >= dec.Length && !dec.read() {
			return nil, dec.raiseInvalidJSONErr(dec.cursor)
		}
		c = dec.Data[dec.cursor]
		if c != 'u' {
			i += utf8.EncodeRune(str, r)
			str, err = dec.appendEscapeChar(str[:i], c)
			if err != nil {
				dec.Err = err
				return nil, err
			}
			i++
			dec.cursor++
			return str[:i], nil
		}
		dec.cursor++
		r2, err := dec.getUnicode()
		if err != nil {
			return nil, err
		}
		combined := utf16.DecodeRune(r, r2)
		if combined == '\uFFFD' {
			i += utf8.EncodeRune(str, r)
			i += utf8.EncodeRune(str, r2)
		} else {
			i += utf8.EncodeRune(str, combined)
		}
		return str[:i], nil
	}
	i += utf8.EncodeRune(str, r)
	return str[:i], nil
}

func (dec *Decoder) appendEscapeChar(str []byte, c byte) ([]byte, error) {
	switch c {
	case 't':
		str = append(str, '\t')
	case 'n':
		str = append(str, '\n')
	case 'r':
		str = append(str, '\r')
	case 'b':
		str = append(str, '\b')
	case 'f':
		str = append(str, '\f')
	case '\\':
		str = append(str, '\\')
	default:
		return nil, InvalidJSONError("Invalid JSON")
	}
	return str, nil
}

func (dec *Decoder) getUnicode() (rune, error) {
	i := 0
	r := rune(0)
	for ; (dec.cursor < dec.Length || dec.read()) && i < 4; dec.cursor++ {
		c := dec.Data[dec.cursor]
		if c >= '0' && c <= '9' {
			r = r*16 + rune(c-'0')
		} else if c >= 'a' && c <= 'f' {
			r = r*16 + rune(c-'a'+10)
		} else if c >= 'A' && c <= 'F' {
			r = r*16 + rune(c-'A'+10)
		} else {
			return 0, InvalidJSONError("Invalid unicode code point")
		}
		i++
	}
	return r, nil
}

func (dec *Decoder) skipString() error {
	for dec.cursor < dec.Length || dec.read() {
		switch dec.Data[dec.cursor] {
		// found the closing quote
		// let's return
		case '"':
			dec.cursor = dec.cursor + 1
			return nil
		// solidus found start parsing an escaped string
		case '\\':
			dec.cursor = dec.cursor + 1
			err := dec.skipEscapedString()
			if err != nil {
				return err
			}
		default:
			dec.cursor = dec.cursor + 1
			continue
		}
	}
	return dec.raiseInvalidJSONErr(len(dec.Data) - 1)
}

func (dec *Decoder) skipEscapedString() error {
	start := dec.cursor
	for ; dec.cursor < dec.Length || dec.read(); dec.cursor++ {
		if dec.Data[dec.cursor] != '\\' {
			d := dec.Data[dec.cursor]
			dec.cursor = dec.cursor + 1
			nSlash := dec.cursor - start
			switch d {
			case '"':
				// nSlash must be odd
				if nSlash&1 != 1 {
					return dec.raiseInvalidJSONErr(dec.cursor)
				}
				return nil
			case 'u': // is unicode, we skip the following characters and place the cursor one one byte backward to avoid it breaking when returning to skipString
				if err := dec.skipString(); err != nil {
					return err
				}
				dec.cursor--
				return nil
			case 'n', 'r', 't', '/', 'f', 'b':
				return nil
			default:
				// nSlash must be even
				if nSlash&1 == 1 {
					return dec.raiseInvalidJSONErr(dec.cursor)
				}
				return nil
			}
		}
	}
	return dec.raiseInvalidJSONErr(dec.cursor)
}
