package iteration

func Repeat(ch string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += ch
	}
	return repeated
}
