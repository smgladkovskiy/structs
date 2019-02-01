package encoder

const hex = "0123456789abcdef"

func StringToBytes(str string) []byte {
	var buf []byte
	l := len(str)
	buf = append(buf, '"')
	for i := 0; i < l; i++ {
		c := str[i]
		if c >= 0x20 && c != '\\' && c != '"' {
			buf = append(buf, c)
			continue
		}
		switch c {
		case '\\', '"':
			buf = append(buf, '\\', c)
		case '\n':
			buf = append(buf, '\\', 'n')
		case '\f':
			buf = append(buf, '\\', 'f')
		case '\b':
			buf = append(buf, '\\', 'b')
		case '\r':
			buf = append(buf, '\\', 'r')
		case '\t':
			buf = append(buf, '\\', 't')
		default:
			buf = append(buf, `\u00`...)
			buf = append(buf, hex[c>>4], hex[c&0xF])
		}
		continue
	}
	buf = append(buf, '"')
	return buf
}
