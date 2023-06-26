package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var firstApart, countFloor, secondApart, secondIn, secondFloor, countFloorIn int
	var firstFloor, firstIn = -1, -1
	fmt.Fscan(in, &firstApart, &countFloor, &secondApart, &secondIn, &secondFloor)

	if secondFloor == 1 && secondIn == 1 {
		if firstApart < secondApart {
			firstFloor, firstIn = 1, 1
		} else if countFloor == 1 {
			firstFloor, firstIn = 1, 0
		} else if firstApart < countFloor {
			firstFloor, firstIn = 0, 1
		} else {
			firstFloor, firstIn = 0, 0
		}
		out.WriteString(strconv.Itoa(firstIn) + " " + strconv.Itoa(firstFloor))
		return
	}

	countFloorInChek := float64(secondApart) / float64(secondFloor+(countFloor*(secondIn-1)))
	if countFloorInChek < 1.0 {
		out.WriteString(strconv.Itoa(firstIn) + " " + strconv.Itoa(firstFloor))
		return
	}

	countFloorMin := int(math.Round(float64(secondApart) / float64(secondFloor+(countFloor*(secondIn-1)))))
	countFloorMax := int(math.Round(float64(secondApart) / float64(secondFloor-1+(countFloor*(secondIn-1)))))

	if secondApart > countFloor*countFloorMin+((secondIn-1)*countFloor*countFloorMin) {
		out.WriteString(strconv.Itoa(firstIn) + " " + strconv.Itoa(firstFloor))
		return
	} else if firstApart == secondApart {
		out.WriteString(strconv.Itoa(secondIn) + " " + strconv.Itoa(secondFloor))
		return
	}
	var ok bool
	for i := countFloorMin; 0 <= countFloorMax-i; i++ {
		if ok && countFloorInChek == float64(secondApart/secondFloor+(countFloor*(secondIn-1))) {
			newFirstIn, newFirstApart := GetIn(firstApart, countFloorIn, countFloor)
			firstFloor = firstApart / (countFloorIn * countFloor)
			if newFirstApart/countFloorMin != int(math.Round(float64(newFirstApart)/float64(countFloorMax))) {
				firstFloor = 0
				out.WriteString(strconv.Itoa(newFirstIn) + " " + strconv.Itoa(firstFloor))
				return
			}
			firstIn = 0
			out.WriteString(strconv.Itoa(firstIn) + " " + strconv.Itoa(firstFloor))
			return
		}
		if ok && float64(secondFloor) == math.Round(float64(secondApart)/float64(i))-float64((secondIn-1)*countFloor) {
			newFirstIn, newFirstApart := GetIn(firstApart, countFloorIn, countFloor)
			if math.Round(float64(newFirstApart)/float64(countFloorMin)) == math.Round(float64(newFirstApart)/float64(countFloorMax)) {
				firstFloor = 0
				out.WriteString(strconv.Itoa(newFirstIn) + " " + strconv.Itoa(firstFloor))
				return
			}
		}
		if float64(secondFloor) == math.Round(float64(secondApart)/float64(i))-float64((secondIn-1)*countFloor) {
			countFloorIn = countFloorMin
			ok = true
		}
	}
	countFloorIn = countFloorMax

	firstIn, firstApart = GetIn(firstApart, countFloorIn, countFloor)
	if countFloorInChek == float64(secondApart/secondFloor+(countFloor*(secondIn-1))) {
		if firstApart/countFloorMin != int(math.Round(float64(firstApart)/float64(countFloorMax))) {
			firstFloor = 0
			out.WriteString(strconv.Itoa(firstIn) + " " + strconv.Itoa(firstFloor))
			return
		}
		firstIn = 0
	}

	firstFloor = firstApart / countFloorIn
	if float64(firstApart)/float64(countFloorIn) == float64(firstApart/countFloorIn) {
		out.WriteString(strconv.Itoa(firstIn) + " " + strconv.Itoa(firstFloor))
		return
	}
	out.WriteString(strconv.Itoa(firstIn) + " " + strconv.Itoa(firstFloor+1))

}

func GetIn(firstApart, countFloorIn, countFloor int) (int, int) {
	firstIn := 1
	for firstApart-countFloorIn*countFloor > 0 {
		firstApart = firstApart - countFloorIn*countFloor
		firstIn++
	}
	return firstIn, firstApart
}
