package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var inTel, onTel1, onTel2, onTel3 string

	//var answ bool
	fmt.Fscan(in, &inTel)
	fmt.Fscan(in, &onTel1)
	fmt.Fscan(in, &onTel2)
	fmt.Fscan(in, &onTel3)

	inTel = string(format(inTel))
	onTel1 = string(format(onTel1))
	onTel2 = string(format(onTel2))
	onTel3 = string(format(onTel3))
	if inTel == onTel1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	if inTel == onTel2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	if inTel == onTel3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func format(in string) (new []rune) {
	for _, r := range in {
		if unicode.IsDigit(r) {
			new = append(new, r)
		}
	}
	if len(new) == 11 && (new[0] == '7' || new[0] == '8') {
		new = new[1:]

	}
	if len(new) == 7 {
		new = func(s []rune) []rune {
			res := make([]rune, len(s)+3)
			copy(res[3:], s)
			res[0] = '4'
			res[1] = '9'
			res[2] = '5'
			return res
		}(new)
	}

	return
}
