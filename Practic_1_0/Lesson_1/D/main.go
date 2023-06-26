package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b, c float64
	fmt.Fscan(in, &a, &b, &c)
	if a == 0 && b == 0 && c == 0 {
		out.WriteString("MANY SOLUTIONS")
	} else if c < 0 {
		out.WriteString("NO SOLUTION")
	} else if (a+b) == c*c && (2*a+b) == c*c {
		out.WriteString("MANY SOLUTIONS")
	} else {
		if a != 0 && (c*c-b)/a == float64(int((c*c-b)/a)) {
			out.WriteString(strconv.Itoa(int((c*c - b) / a)))
		} else {
			out.WriteString("NO SOLUTION")
		}
	}
}
