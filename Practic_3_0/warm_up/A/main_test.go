package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkSlugify(b *testing.B) {
	s := randomWord(1000)
	name := fmt.Sprintf("%s", s)
	b.Run(name, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			slugify(s)
		}
	})
}

func randomWord(n int) string {
	const letters = "AQWE RTY UIOPSDF GHJKLZX CVBNMab cdefghijkываываы ывауаф ыфафуафакп990403005349953400534 lmnopq rstu vwxyz!№;%:?*(){}[]/`,.|'\"_  "
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}
