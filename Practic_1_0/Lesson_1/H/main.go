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
	var timeOnFirstWay, timeOnSecondWay, countTrainOnFirstWay, countTrainOnSecondWay int
	fmt.Fscan(in, &timeOnFirstWay, &timeOnSecondWay, &countTrainOnFirstWay, &countTrainOnSecondWay)

	minTimeOnFirstWay := timeOnFirstWay*(countTrainOnFirstWay-1) + countTrainOnFirstWay
	maxTimeOnFirstWay := timeOnFirstWay*(countTrainOnFirstWay-1) + countTrainOnFirstWay + (2 * timeOnFirstWay)
	minTimeOnSecondWay := timeOnSecondWay*(countTrainOnSecondWay-1) + countTrainOnSecondWay
	maxTimeOnSecondWay := timeOnSecondWay*(countTrainOnSecondWay-1) + countTrainOnSecondWay + (2 * timeOnSecondWay)
	minTime := GetMax(minTimeOnFirstWay, minTimeOnSecondWay)
	maxTime := GetMin(maxTimeOnFirstWay, maxTimeOnSecondWay)

	if minTime > maxTime || minTime < 0 || maxTime < 0 {
		out.WriteString("-1")
		return
	}

	out.WriteString(strconv.Itoa(minTime) + " " + strconv.Itoa(maxTime))

}

func GetMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
