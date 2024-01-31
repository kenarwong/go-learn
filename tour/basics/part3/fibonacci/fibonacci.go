package fibonacci

import "fmt"

func fibonacci() func() int {
	s := []int{}
	return func() int {
		switch len(s) {
		case 0:
			s = append(s, 0)
		case 1:
			s = append(s, 1)
		default:
			s = append(s, s[len(s)-2]+s[len(s)-1])
		}

		return s[len(s)-1]
	}
}

func RunFibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
