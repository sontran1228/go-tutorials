package iteration

const repeatCount = 5

func Repeat(s string) (repeat string) {
	for i := 0; i < repeatCount; i++ {
		repeat += s
	}
	return
}
