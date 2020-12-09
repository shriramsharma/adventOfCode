package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	splitInput := strings.Split(input, "\n")

	allIDs := make([]int,0)

	total := 0

	for _, inp := range splitInput {
		total++
		rowHigh := 127
		rowLow :=0
		rowMid := 0
		colHigh := 7
		colLow := 0
		colMid := 0
		inp = strings.TrimSpace(inp)
		for i := 0; i <= 6; i++ {
			rowMid = (rowHigh+rowLow)/2
			if inp[i] == 'F' {
				rowHigh = rowMid
				continue
			}
			rowMid += 1
			rowLow = rowMid
		}
		for i := 7; i < len(inp); i++ {
			colMid = (colLow+colHigh)/2
			if inp[i] == 'L' {
				colHigh = colMid
				continue
			}
			colMid +=1
			colLow = colMid
		}

		seatID := (rowMid*8)+colMid
		allIDs =append(allIDs,seatID)
		fmt.Printf("Row: %d Column: %d Seat ID: %d \n",rowMid,colMid,seatID)

	}


	sort.Ints(allIDs)

	for i := 0; i < len(allIDs); i++ {
		if allIDs[i]+1 != allIDs[i+1] {
			fmt.Printf("This is the ID: %d \n",allIDs[i]+1)
			break
		}
	}

	//fmt.Println(allIDs)



}

//var input = `BBFFBBFRLL
//FFFBBBFRRR`

var input = `BFFFBFFRLR
FBBBFFBRRL
FFBFBFFRRR
FBBFBFFRRR
BBBFFFFLLR
FFBFBBFRRR
BFFBBFFLRL
BBBFBBFLRR
BFFBFBFLRR
BBBFFBFLLL
BFBFBBBRRR
FBBBBBFLRR
BBFFBBFRRL
BFBBBBBRLR
FBBFFFFRLR
BBFFBFFRRL
FBBFFFBRRR
FFBBBBFRRL
BFBBBFFRRR
BBBFBFBRRR
FBFFBFBRLR
BFFBBFBLLL
BBFFFFFRRL
FBBBBFBRLR
FBBBBBFRLR
BFFBBBFLRR
FBFBFBFLLR
FFBFFFBRRR
BBBFFBBLRR
BBBFFBBLRL
BBFBFFBLRR
BBBFFFFRRL
FBFFFBFLRR
BBFBBBFRLL
BBBFFBFLRR
FBFFFFBLRR
FBBBBBFLRL
BBFFFBBRRL
BBFFBFBLRL
BBBFBBFRLR
FFBBFFFRRL
BBFBBBFLRL
BFBBFFBRRR
BFBFFBFLRR
BFBBBFFRLR
FFFBBBBLRL
BBFBFFBRRL
FBFBBFBLRL
FBBBBBBRLL
FBBFBBFLLR
FBBBFFFLLR
BFFFFFBRLR
BBFFFBBLRL
FBBBFFBLRL
FFFBBBBRRR
FBBBFBFLRR
BFFBFFBLLL
BFBBBBBLLR
BBFFFFBLRR
FFBFFFBRRL
BFFFBFFRRL
FFBBFBFLLL
FBFFFFFLRR
FBFFBFFLRR
FBBBFFFRRR
FBFFBFBLRR
FFBBBFFRLR
BBFFFFFRLR
BFFFBFBRRR
FBFBBBFRLL
FFBBFFBRRL
FFBBBBBRLL
FBBFFBBRRR
FBBBFFBLLR
FBBFBFFLRR
FBFFFBFRLR
BBFFBFFLLL
BFBFFFBLRL
BFFBBFFLLL
FBFBBBFRRR
BFFBFFFLLL
BBFBBFFRLR
BFFBBBBRLL
FBBFFBBRRL
FFBFFFFRLR
BFBFBFBLRR
FFBBBBBRRR
FFBBFBBLLR
FBBFFBFLRR
FBFBBBBLLL
BFBBBBFRRR
FFBFBBFLLR
BFBFBBFRLL
BFBFFFBRRL
FFBBFBFRRR
FBFFBBBRLR
BBBFFBFRRR
FFBFFFFRRL
BFBFFBBLRR
FBBBBFBLLL
FFFBFFFRRR
BBFBBFBRRL
FFFBFFBRRR
FBBBFBFLLR
FFBBFFFLLR
BFFBBFFRRR
FFBBFBBLRR
FFBBBBBLLR
BBFBFBFRRL
BBFBBFBLLL
BBFBFBFRLR
BBFBFFBRRR
FFBBFBBLRL
FBFFFFFRLR
FBFBFBBLRR
BFFFFBFRLL
BBBFFBBRRR
FFBFBFFRRL
BFFFBBFLLL
FBBBBBBLRL
BFBBFBFRRR
BFFFFFBLLR
BFBBBFFLRL
BBBFFFBRRL
BBFBFFFLRR
FBFFFFFRRR
BBFFFBBLRR
FBFBBFBRLR
FBFBBBBRLL
FBFFFFFRRL
FBFBBFFRLR
BFBBFBBRLL
BFFBBFFRRL
FBFFFBBRLR
FFBBBBFLRL
BFBBBBBRRL
BBFBBBFRRL
FFBBBFFLRR
FBBBBFBRLL
BFFFBFFLRL
BFBFBFBLLL
BBFFBFFRLR
FBBBBFBRRR
BFBBFFBLRR
FBFFBBBLRR
FBFFFFFLLL
BFFBFBFRRL
FBBFFBFRRL
BBFBFBBRRL
BBFBBFBRRR
BFBFBBFLRL
BBFBFFBLLR
BBFFBFBLRR
FBBFFFBLRL
BFBFFFFRRL
BFBBFBFRRL
BFBBFFBLLL
BFFFFBFLLL
BBFBFFFLLR
BFFFBFBLLL
BBFBFBFRRR
BFFBBBFLLL
FBFFFFBRLR
BFFBBFBLRR
BBFBBBFRLR
BBFBFBBRLR
BFBBBBBRLL
BFFFFFFLRR
FFBBBBFLLR
BBBFFBFLLR
BFBFBBFRRL
BBFBBBFRRR
FFBBBBFLRR
FBFBBFBLLL
FFBBBFFRLL
FBFBFFFRLR
BFBBBFBLLL
BFBFBFBRRL
BBBFBBBLRR
BBFBBFFLLR
BFBFFFBLLL
BFFBFFFLLR
FBFBFFBLLR
BBFFFBBLLL
BBFFBFBLLR
FFBFBBFLLL
BFFFFBFRLR
FBFFBFFLRL
BFBFFBFRLL
BBFFBBBRLR
FBFBFFFRRL
FBFBBFFRRR
FBFBFBBLLL
BFBFBBBRRL
BFBFFBFRRL
BBBFBFFLLL
FFBFBBBRLR
FBFBBBFLRR
FBBFBBBRRR
BFBBBBFRRL
FFFBFBBLLL
FFFBBBFLLR
FFFBBBFRRR
FFBBFFFLRR
FFBFBFBLRL
FBBFBBFLLL
FFFBBFBRLL
FBFFBFBRLL
FFBBFBBRLR
BBFBFFFLLL
BFBFFBFLRL
FFFBFBFLLL
BBBFBFBRLL
FFBFBBBRRL
FBBBBFBLLR
BBBFBFBRRL
BFBBBFBLLR
FFFBBFFLRR
BFBBFFBRLL
FFBFBBFRRL
BBFFFFFRRR
BBBFFBBRLR
BFFFBBBRRR
BBBFFFFRLL
BBFBBBBLRL
FFBFBFFLRL
FFBFBFBRRL
BFFBFFBRRR
FBFFFBBRLL
FBBFBBBRLR
BFFFFFBLRL
FFBBFFBLRL
BBBFFBBRRL
FFBFBFFLLL
FBFFFBFRLL
BFBFFBFRRR
BFBFBBFLRR
FFBBFFBRRR
BFBFBBFRLR
FFFBBBBRLL
BBFBBBFLLR
FBFFBBBLLL
FFBFBFFLLR
BBFFFBBRLL
BFBFBFFRRR
BFBFBFBRLL
BFFFFBBLRL
FFBFFBFLRL
BFBBFFBRLR
FBBBBBFLLR
FBFBFFFLLL
BBFBFFFLRL
BBBFBFFLLR
FFFBFBFLLR
BFFBFBBRLL
FBBFBBBLLL
BFFFBFBRRL
FFFBBFBLRR
FBFFBFFRLL
BBFFBBBRRL
BFBBFBBLRR
FBFFBFBRRR
FBFFBBFRRL
BBFFFBFRLL
FFBBBFBLRR
BFFFBBFRRR
BBBFFBFLRL
FFBFFBFLLR
FBFFBFFLLL
FFFBBBFLLL
BFBFBBBLRL
BFBBFBFLRR
BFBFFFBRLL
FBFBFFBRRR
BFFFBBFRLL
FBBFFBFRRR
BBBFFBBLLL
FBFFFFBRLL
FFBBBFBLLR
BBFFBFFRRR
FFFBBBBLRR
BBFFFFBRRL
BBFBFBBLRL
FBBBFBFLRL
FBFFBBFRLL
FFBFBBBLLL
BFBFBFFRLL
FBBBFBBRLL
FBFBBFBRRL
BBFFBFFLRR
BBFFFFBRLR
FFFBBBBRLR
BFFFBBFLRL
BFBFBFFLRR
FBBBBFFRRR
FBFBBBBLLR
FBBFFBBRLR
BFBFBBBLLR
BBBFFBBLLR
BFBFFFFLRR
BFFFBBBRLR
BBFFFFBLLR
FBFBBFFLRL
BFFBFFBLLR
BFFFBFBLRR
BFBBBBFLRL
BBBFFBFRLL
BFBBFBFLLL
BBBFFFBLRL
FFBFBBFRLR
FBBBFBBRRL
FBFFBBFLRL
FBBBBFFLRR
FFFBBBFLRR
FBBFBFFLLL
BBFFFFFLLR
BBBFFFBLLL
BFFBBBBRRR
BFFBBBBLLL
FBFBBBBRRL
FFFBBFBLLR
BBFBFBBLLR
BBFBFFBRLR
FFFBFBBRLR
FBFFFBFLLL
FBFBBBBRRR
FBBFBBFRRL
FBBBFBFRRL
BBFBBFFLLL
BFBBFFFLLR
FBBFFFFLLR
BFBBBFFLRR
BBFBBFFRRR
BBFBFFBLLL
FFFBBFFLRL
FBBBBFBLRL
FBBBBBBLLR
BFBBFFFRRR
FBBFBFFLRL
BFBBBBBLLL
FFBBFFBLLL
BBBFFBBRLL
FFBBBBFRRR
FBBFFFBRLL
BBFBFBFRLL
BFBFFBBLLR
BFFFFFFRRR
BFFFFFFRRL
FBFBFBBRRL
FBBFBFBLRL
BBBFBBFLLL
FBFFBFFRRR
BFFFFBBLLR
FBFFBBBLLR
FBBBFFBLLL
FFBBBFBLLL
BFFFBFBLRL
FBBFBBFLRL
BBFFFBBRLR
BFBFFFFLLL
BBFFBBBRRR
FBFBFFBRLL
BFFBFBFLLL
FFBFFFBRLL
FFBFFFBRLR
BBBFFFFLLL
FFBFFBFRRL
BFFBBBFRLR
BBFFFFFLRR
FBBBFFBRLR
BFBFFBFRLR
BFFBFFFLRL
BFFBFFFRLR
FBBFFFFLLL
FBFFFBFLLR
FBFBFFBRLR
BFFFFBBRRR
BBBFBFFRRL
FFFBBFFLLL
FFBBFFFRRR
FBBFFBFLRL
BFFBFBBLLR
FFBBBFBRRR
FFFBFBFRRL
FFBBBBBLLL
BBFFBBBLLR
FFFBBBFRRL
BFFFBBFLRR
BFFBFFFRRR
BFBBBFFRLL
FBBFFBBLRL
FBBFBBFLRR
FBBFFBBLRR
FFBFFBBLRL
FBBFFBFRLL
BFFBBBFRRL
FBFBFFFLLR
FFBFBBBLLR
FFFBFBFRLL
BBBFBBBLRL
FBBFFFFLRR
FBBBFFFRLL
BBFBBBBRLL
FBFBFFFRLL
FBFFFFFLRL
BFBFFFFRRR
BBFFBBFLRL
FBBFFBFRLR
FFBBFFFRLL
BFBBFFFLRL
BBFBFFFRRR
BBFBFFFRLR
FBFBBBFRLR
BFBBFFBRRL
BBBFFBFRRL
BFFFFFFRLL
FFFBFFBRLR
BBFFFBFLRR
FBBFFBFLLL
FBFFBFBLLR
FBFBBFBLRR
FBFBFBFRRR
FFBFBFBRLL
BFBFFFBRLR
BFBBFFBLLR
FBBBFBFRLL
FBBFBFBLLR
FFFBBFFRRL
FBFBFFBLRR
FBBFBFBRLL
FBFFFFBLLR
FBBBFFFLRL
BBBFBBFRRL
FBBFFFFRLL
BFBBBBFLLL
BFFBBBFLLR
FBFBBBFLLL
FBBFFFFRRL
BFFFBFFRRR
BBBFFFBLLR
FFBBFBFRLR
FBBFBBBRLL
BBFFFFFLRL
FBBFFBFLLR
BFBFFBBRRR
BFBBFFFRRL
FFBBBBFRLL
BBFBFFBRLL
BFFFFFBLLL
FFBBBFBLRL
FFFBBFFRRR
BBFBFBFLRR
FFFBFFBLRL
BFFFBBBLRL
BFBBFBBLLR
FFBBBFFRRL
BBFFFBFLLR
BFFFFBBLRR
FFBFFBFRRR
FFBFBFFRLL
FBFBBBBLRR
FFFBBFBRRR
BFBFBFBRRR
FFFBFBFLRR
BFFBFFFRLL
BFFFFFBRRL
BFFBFBFLLR
BBFFBBFRLR
BFFFFFFLLR
BBFFFBBRRR
BBFBFBBRLL
FFFBBBFRLR
FBFBBBFRRL
FFFBFFBRRL
BFBFFFFLLR
BFBBFBBRRL
FFBFBFFRLR
FFFBBBBLLL
FFBFBFBRRR
FBBFBBBLRR
BFFBFFBLRL
BFFFBFFLLR
BFFFFBFLLR
BBFBFBBLRR
BFFBFFFRRL
BFFBBFBRLL
BBFFBFBRRL
BBBFFFFRRR
BFFBFFBRLL
FBFBBFFRLL
FBBFFFFLRL
BBBFBFFLRR
FFBBBBBLRL
FBBFFFBRLR
FFBBBBBRLR
FBBBFFBLRR
FBBBFBBLLR
FFBBBFFLLR
BFFFFFFLLL
FFFBFBBRRL
FFBBBBFRLR
FBFFFFBRRR
FFBFFBBLLL
FBBFFBBLLR
FFFBFBFRRR
BFFBBFBLRL
BFBFFFBLLR
BBFFFFBRRR
FBBFBFBLRR
FBBFFFBRRL
FFBBBFFRRR
FBFFFFBLRL
FBBFFFBLLL
FBFFFBBLRL
FBFBFBFLLL
BFFBFFFLRR
BFBFFBFLLL
FFBFFBBLRR
FBFBFFBRRL
FBFBBBBRLR
BBFFBFBRRR
BBBFBBFLRL
BFBBFBBLLL
FFBFBFBLLR
FBBBBFFLRL
FFBBFBBRRR
FBFFBBBRLL
FFBBFFFLLL
BFFFFFBRLL
BBFBBBBLLR
FBBBFFBRRR
BFFBBFBRLR
FBBBBFBRRL
BBFFBBFLRR
FFBFBFBLRR
FFFBFBBRRR
FFBFFFFRLL
FBBFFFFRRR
BFBFBFFLLR
BBFFBFFLLR
FFBFBBBRRR
BFBBBFFRRL
BBFFBFFRLL
FBFBBFFLLL
BFFBFBBLRL
FFBFFBFLRR
BBFBFFFRRL
BBFBFBFLLL
FBBBFFFLRR
FBBFFBBLLL
BBFFBBFRLL
FFBBFBBRRL
BBFBBFFLRL
BFBFFBFLLR
BBBFFFBRLL
FFBBFBFLRL
FBFBFBBLRL
FFBFFFFLLR
BBFBBBBRRR
FBBFFBBRLL
FFBBBFBRLL
FFBFFBBLLR
BBFBBBBLLL
BFBBBFFLLR
BBFBBFBRLR
FBBBBBFLLL
FBBBBFFRLR
BFBBFFFLRR
FFBFBFBLLL
BFFFFBFRRR
FBFBFBFRLL
FFBFBBFLRL
BFFFFBBLLL
BBFBFFFRLL
FBFFFBBRRR
FBFBBBBLRL
BBFFBBBLLL
BFFFBBBLLL
BFBBBFBLRL
BFFBFBBLRR
BBFFBFBRLR
BBFFFBBLLR
BFBFFFFRLR
BBFBFFBLRL
BFFFBFFLLL
BFBFFFBLRR
FFBBFFFLRL
FBFBFBBRLR
FBFFBBBRRL
BBBFBFBRLR
BFFFBFBRLR
BBFFBBBRLL
FBBBBBBRRR
FFFBFBBLRL
BFBBFFFLLL
FBBFBFBRLR
FFBFFBBRRR
BFFBBBBLLR
FBFBFBBLLR
BFBFBBBRLR
FBFBBBFLRL
BBFBBBFLRR
FBBBBFFRRL
FBBBFBBRLR
FFFBFFBLRR
BFFFFBFLRL
BBBFFBFRLR
FFBFBFBRLR
BBBFBFFLRL
BBFFFFFLLL
FBFFBBFLRR
FFBBFBFLRR
FFBBBBBRRL
BBFFBFBLLL
FFBFFFBLLR
BFFFFFFLRL
FFBBFFBRLL
FFBBFBFRLL
FBBBFFBRLL
BBFFFFBRLL
FFBFFBBRRL
BBFFBBBLRR
BBBFBFBLRR
FFBBBBBLRR
FBBFBBBRRL
BFBBFBFRLL
FBFFBFFRLR
FFFBBBBLLR
FFFBBFBRRL
FFBFBBFLRR
FFBFFFFLRR
BFBFBBFRRR
BFFFFFFRLR
FFBFFBBRLL
BFBFBBFLLL
BBFBFBFLLR
BFBFFFFRLL
BFBBFBFLLR
FFBBFBFRRL
FFBBBFFLLL
FBBBFBBLLL
BBBFBBFLLR
BFFBBFFRLR
FBBFBBBLLR
FFFBBBFLRL
BFFFFBBRLR
BFBBFBFRLR
FBBBFFFRLR
BBBFBFBLRL
BFBFBBBLLL
BBFFBBFLLR
BBFBBFBLRR
FFBFFBFLLL
FBBFBBFRRR
FBBFBBFRLR
BFBFBFFRRL
BFFFFBBRRL
FFBBFFBRLR
FFFBFBFLRL
FBFFFFFRLL
FBFFBBFLLL
FBFFFFBLLL
BBBFFFFRLR
FFFBBBFRLL
FBBFBFFRLL
FFFBBFFRLR
FBBFBFFLLR
BFFFBFBRLL
BFBFFBBLRL
FBBFBFBRRR
BFFFBFBLLR
FBFBFBBRLL
FBBFBFBLLL
FFFBBFBRLR
FBBBFBFLLL
BBBFBBFRRR
BFBFBBFLLR
FFBFFFFRRR
BFBBFBFLRL
BFFBBFBLLR
BBFBFBBLLL
BFBBFFBLRL
BFBFFBBLLL
FBBBBBBRLR
FBFFBFBLLL
BFFBBBFRLL
FBBBBBBRRL
FBBBFBBLRL
BBFFFFBLLL
FFFBBFFRLL
FBFFFBBLLR
BFFBBFFRLL
FFBFFBBRLR
BFFFBBFRLR
FFBBFFBLLR
FFBBBBFLLL
BBBFBFBLLL
FBFBBFBRLL
BBBFBBFRLL
BFFFBBFLLR
FBBBBBFRRR
FBFBFFBLLL
BBBFFFFLRL
FBBBBFFLLR
BFFFFBFRRL
BFFBBBFLRL
BFBFFFFLRL
FFFBBFBLRL
BBFBBBFLLL
FBFFBBBRRR
BBFFFBFRRR
BBFFFFFRLL
BBBFBBBLLL
FFBBFFBLRR
BBFBFBBRRR
BFFBFFBRLR
BFBBFFFRLR
FBFBFFBLRL
FBFFBFBLRL
BFFFBFFRLL
FFBFBBBLRL
FBFFBFFLLR
FFBFFBFRLR
FBBFFFBLLR
BBBFFFBRLR
BFBFBBBLRR
FBFBBFFLLR
FBFBFBFRRL
BFFBBBBLRR
BBFFBBBLRL
FFBFFFBLRR
BFFBBFBRRL
FBBBFBBRRR
FFBBBFBRLR
BBFFFFBLRL
FBFBFFFRRR
BFFBFBFRRR
BFBFBFFLLL
BFFBFBFRLL
FBFFFBBRRL
BFFFBFFLRR
FFBFBFFLRR
FBFFBBFRRR
BFFBBFFLRR
FFBBFFFRLR
BFBFBFFRLR
BFBBBFBRRR
BBFBFBFLRL
FFFBFFBLLL
FBBFBBBLRL
FBFBFBBRRR
BBFFBFBRLL
FFFBFFBLLR
FBBFBBFRLL
FFBBFBBRLL
FBFBFFFLRL
BBFBBFBLLR
FBBBBFFRLL
FFBBBFFLRL
FBFFBBBLRL
FBFBFFFLRR
BBFFBBFRRR
BBFFFBFLRL
FBFFFBFRRR
BFBFBFBRLR
FBFFBBFRLR
BFBBFFFRLL
FFBFFFFLLL
FFBBBFBRRL
BFFFBBBRLL
BFFFBBBRRL
FBFBFBFLRR
FBFBBBFLLR
BBBFBFFRLL
BFFBFBBRRL
BFFBBBFRRR
BFFBBBBLRL
FBBFBFFRRL
FFBBFBFLLR
FBBBFBFRRR
BBFFFBFRRL
BFBBBFFLLL
BBFFBFFLRL
FFFBFFBRLL
FBFFFBFRRL
FBFFFBBLRR
FBFBBFBLLR
FBFFBFFRRL
BFBBBFBRLR
BFFFBBFRRL
FBFFBFBRRL
FFBFBBBLRR
BFBFBBBRLL
FBBBFBBLRR
FBFFFBBLLL
BFBFBFBLLR
FBBBBFFLLL
FFBFFBFRLL
BBFFBBFLLL
FFFBBBBRRL
BFBFFBBRRL
BFFBFBBRRR
BFFBFBFLRL
BBFBBBBLRR
FBFBFBFLRL
BFFFBBBLRR
BFFBFBBRLR
BBFFFBFRLR
BBBFFFBRRR
BFBFFBBRLR
FBFBFBFRLR
FFBFBBBRLL
FBFFBBFLLR
BFBFFFBRRR
BFFBFBBLLL
BFBBBBFRLL
BBFBBFBRLL
BFBBFBBRRR
BFFBBFBRRR
BFBBBBBRRR
FFFBFBFRLR
BFFBBFFLLR
BBFBBBBRLR
FFFBFBBRLL
BFFBFFBRRL
BBBFFFFLRR
FBBBFFFRRL
BBFBBBBRRL
BFFFFBFLRR
BBFBBFFRLL
FBBFBFFRLR
FBFBBFFRRL
BFFFFFBRRR
BFFBBBBRLR
BFBFFBBRLL
FBFBBFBRRR
FFBFBBFRLL
FBFFFFBRRL
BFBFBFFLRL
BBFBBFFLRR
BFBBBFBRLL
FBFBBFFLRR
BFFFFBBRLL
BBBFBBBLLR
FBBFBFBRRL
BBFFFBFLLL
BFBBBFBRRL
BFFBFBFRLR
FFFBFBBLLR
FBBBFBFRLR
FBFFFFFLLR
BBBFFFBLRR
BBFBBFBLRL
BFBBBBBLRR
BFFBBBBRRL
FBBFFFBLRR
FBBBBBBLLL
BFBBBFBLRR
BFBBBBFLRR
BFBBBBBLRL
BFBBBBFLLR
FFFBBFBLLL
BBFBBFFRRL
FBBBBFBLRR
BBBFBFFRRR
BFBFBFBLRL
FBBBBBBLRR
BFFBFFBLRR
FFBBFBBLLL
BBBFBFBLLR
FBBBFFFLLL
BBBFBFFRLR
BFBBFBBRLR
FFBFFFBLRL
FFBFFFBLLL
FBBBBBFRRL
FFFBBFFLLR
FBFFFBFLRL
BFBBBBFRLR
BFBBFBBLRL
BFFFFFBLRR
FBBBBBFRLL
FFBFFFFLRL
FFFBFBBLRR`