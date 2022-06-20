package iteration

const repeatCount = 5

func Repeat(c string) string {
	var s string
	for i := 0; i < repeatCount; i++ {
		s += c
	}

	return s
}
