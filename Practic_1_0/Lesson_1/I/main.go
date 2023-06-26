package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var a, b, c, d, e int
	fmt.Fscan(in, &a, &b, &c, &d, &e)

	holeHeight, holeLength := GetMinMax(d, e)
	array := []int{a, b, c}
	sort.Ints(array)
	brickHeight, brickLenght := array[0], array[1]

	if holeHeight >= brickHeight && holeLength >= brickLenght {
		out.WriteString("YES")
		return
	}
	out.WriteString("NO")
}

func GetMinMax(x, y int) (int, int) {
	if x > y {
		return y, x
	}
	return x, y
}
