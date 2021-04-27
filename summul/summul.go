package summul

func SumMultiples(upper int) int {
	s := 0
	for i := 1; i < upper; i++ {
		if i%3 == 0 {
			s += i
			continue
		}
		if i%5 == 0 {
			s += i
			continue
		}
	}
	return s
}
