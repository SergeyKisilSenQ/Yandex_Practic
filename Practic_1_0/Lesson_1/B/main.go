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
	var x, y, z int
	fmt.Fscan(in, &x)
	fmt.Fscan(in, &y)
	fmt.Fscan(in, &z)
	if x+y > z && z+y > x && x+z > y {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
