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
	var lengthFirst, heightFirst, lengthSecond, heightSecond, minLength, minHeight int

	fmt.Fscan(in, &lengthFirst, &heightFirst, &lengthSecond, &heightSecond)

	a1, a2 := getMax(lengthFirst, lengthSecond), heightFirst+heightSecond
	b1, b2 := getMax(heightFirst, heightSecond), lengthFirst+lengthSecond
	c1, c2 := getMax(lengthFirst, heightSecond), heightFirst+lengthSecond
	d1, d2 := getMax(lengthSecond, heightFirst), lengthFirst+heightSecond
	temp1, temp2 := getMinSquad(a1, a2, b1, b2)
	temp3, temp4 := getMinSquad(c1, c2, d1, d2)
	minLength, minHeight = getMinSquad(temp1, temp2, temp3, temp4)
	out.WriteString(strconv.Itoa(minLength) + " " + strconv.Itoa(minHeight))
}

func getMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func getMinSquad(x1, x2, y1, y2 int) (int, int) {
	if x1*x2 > y1*y2 {
		return y1, y2
	}
	return x1, x2
}
