package iteration

func Repeat(c string, r int) (out_string string) {
	for i := 0; i < r; i++ {
		out_string += c
	}
	return out_string
}
