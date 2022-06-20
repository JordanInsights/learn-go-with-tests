package iteration

func Repeat(c string, repeatCount int) string {
	var s string
	for i := 0; i < repeatCount; i++ {
		s += c
	}

	return s
}
