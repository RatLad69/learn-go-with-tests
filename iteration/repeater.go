package iteration

func Repeater(ch string, x int) string {
	var repeated string
	for i := 0; i < x; i++ {
		repeated += ch
	}
	return repeated
}
