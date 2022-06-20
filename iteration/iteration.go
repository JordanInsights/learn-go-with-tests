package iteration

func Repeat(c string) string {
	var s string
	for i := 0; i < 5; i++ {
		s = s + c
	}

	return s
}
