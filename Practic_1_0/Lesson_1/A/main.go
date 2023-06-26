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
	var troom, tcond int
	var mode string
	fmt.Fscanln(in, &troom, &tcond)
	fmt.Fscan(in, &mode)

	switch mode {
	case "freeze":
		if troom > tcond {
			fmt.Println(tcond)
		} else {
			fmt.Println(troom)
		}
	case "heat":
		if troom < tcond {
			fmt.Println(tcond)
		} else {
			fmt.Println(troom)
		}
	case "auto":
		fmt.Println(tcond)
	case "fan":
		fmt.Println(troom)
	default:
		fmt.Println("incorrect data")
	}
}
