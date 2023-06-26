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
	var material, pieceMass, piece, detailMass, totalCount int
	fmt.Fscan(in, &material, &pieceMass, &detailMass)
	out.WriteString(strconv.Itoa(productionProcess(material, pieceMass, piece, detailMass, totalCount)))

}

func productionProcess(material, pieceMass, piece, detailMass, totalCount int) int {
	if material < pieceMass {
		return 0
	}
	if pieceMass < detailMass {
		return piece
	}
	piece = material / pieceMass
	remainsMassPiece := material % pieceMass
	count := 0
	for i := piece; i > 0; i-- {
		count += pieceMass / detailMass
	}
	remainsMassDet := (pieceMass * piece) % (detailMass * count)
	remainsTotal := remainsMassPiece + remainsMassDet
	totalCount += count
	if remainsTotal/pieceMass >= 1 {
		material = remainsTotal
		return productionProcess(material, pieceMass, piece, detailMass, totalCount)
	}
	return totalCount
}
