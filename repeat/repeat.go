package iteration

func Repeat(letter string, repeats int) string {
	var sum string

	for i := 0; i < repeats; i++ {
		sum = sum + letter
	}

	return sum
}
