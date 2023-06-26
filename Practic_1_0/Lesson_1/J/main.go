package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var a, b, c, d, e, f int
	fmt.Fscan(in, &a, &b, &c, &d, &e, &f)
	if b != 0:
	y = (c * e - a * f) / (b * c - a * d)
	if c != 0:
	x = (f - d * y) / c
	print(x, y)
	else:
	if a != 0:
	x = (f * b - d * e) / (b * c - d * a)
	if d != 0:
	y = (f - c * x) / d
	print(x, y)
}


