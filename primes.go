package numbermath

// import (
//     "fmt"
// )

type singleSieve struct {
	a []bool
}

var s = singleSieve{a: make([]bool, 0)}

// Prime number generator
func PrimeGenerator(under int) chan int {
	m := (under - 1) / 2
	if len(s.a) < m {
		// fmt.Println("Recalculate sieve.")
		sieveEraFast(under)
	}

	c := make(chan int)
	p := 0
	go func() {
		defer close(c)
		c <- 2
		for i := 0; i < m; i++ {
			p = (i+1)*2 + 1
			if s.a[i] {
				c <- p
			}
		}
	}()
	return c
}

func sieveEraFast(under int) {
	m := (under - 1) / 2
	s.a = make([]bool, m)
	for i := 0; i < m; i++ {
		s.a[i] = true
	}
	i := 0
	p := 3
	// fmt.Println(2)
	count := 1
	for p*p <= under {
		if s.a[i] {
			// fmt.Println(p)
			count++

			j := (p*p - 3) / 2
			for j < m {
				s.a[j] = false
				j = j + p
			}
		}
		i++
		p += 2
	}
	for i < m {
		if s.a[i] {
			// fmt.Println(p)
			count++
		}
		i++
		p += 2
	}
	// fmt.Printf("%d primes found\n", count)
	return
}
