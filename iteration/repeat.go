package iteration

// const repeatCount = 5

func Repeat(character string, repeatNumber int) string {
	var repeated string
	for i := 0; i < repeatNumber; i++ {
		repeated += character
	}
	return repeated
}
