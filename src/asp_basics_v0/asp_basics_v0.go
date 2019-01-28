package main

import (
	"fmt"
	"time"
)

	
func main() {
	start := time.Now()
	n := 1000000000
	for i := 0; i < n; i++ {
		type hammings struct {
			barcode string
			query string
			hamdist int
		}
		s := hammings{barcode: "AGCT", query: "AGGT", hamdist: 1}
		fmt.Printf("Barcode: %s\nQuery: %s\nHamming Distance: %d\n", s.barcode, s.query, s.hamdist)
	}
	
	elapsed := time.Since(start)
	fmt.Printf("It took %s to run %d iterations.\n", elapsed, n)
}
