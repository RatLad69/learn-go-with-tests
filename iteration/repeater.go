package iteration

func Repeater(ch string, x int) string {
	var repeated string
	for i := 0; i < x; i++ {
		repeated += ch
	}
	return repeated
}

func LeftPad(inString string, size int) string {
	if len(inString) >= size {
		return inString
	}
	preString := Repeater(" ", size-len(inString))
	return preString + inString
}
