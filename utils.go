package jinja_go

func byteInList(c byte, l []byte) bool {
	for _, item := range l {
		if c == item {
			return true
		}
	}
	return false
}
