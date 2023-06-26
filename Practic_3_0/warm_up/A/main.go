package main

import (
	"fmt"
)

var transform = [256]byte{
	'-': '-',
	'0': '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A': 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'a': 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
}

func slugify(s string) string {
	buf := make([]byte, len(s))
	n := 0
	skip := true
	for _, b := range []byte(s) {
		b = transform[b]
		if b == 0 {
			skip = true
			continue
		}
		if skip && n != 0 {
			buf[n] = '-'
			n++
		}
		skip = false
		buf[n] = b
		n++
	}
	return string(buf[:n])
	// buf = buf[:n]
	// return *(*string)(unsafe.Pointer(&buf))
}
func main() {
	const phrase = "s!!!!"
	slug := slugify(phrase)
	fmt.Println(slug)
	// a-100x-investment-2019
}
