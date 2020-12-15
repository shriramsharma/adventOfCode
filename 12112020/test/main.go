package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strings"
)

// flag vars
var renderStates bool

func main() {
	flag.BoolVar(&renderStates, "r", false, "if true, will produce a .png file for every calculated state")
	flag.Parse()

	initialState := newFerry(data)

	// Part One
	calcPartOne(initialState)

	// Part Two
	//calcPartTwo(initialState)
}

type ferry struct {
	cells         map[coords]rune
	width, height int
}

type coords struct {
	x, y int
}

func newFerry(rawFerry string) *ferry {
	dataLines := strings.Split(rawFerry, "\n")
	initialState := &ferry{
		height: len(dataLines),
		width:  len(dataLines[0]), // all lines have the same length
		cells:  make(map[coords]rune, len(dataLines)*len(dataLines[0])),
	}
	for y, dataLine := range dataLines {
		for x, cell := range dataLine {
			initialState.cells[coords{x, y}] = cell
		}
	}
	return initialState
}

var runeWeight = map[rune]int{occupied: 1} // free seats, floor and out-of-bounds are counted as Zero

func (f *ferry) countAllOccupiedSeats() int {
	var count int
	for x := 0; x < f.width; x++ {
		for y := 0; y < f.height; y++ {
			count += runeWeight[f.cells[coords{x, y}]]
		}
	}
	return count
}

/******************************************************************/
/************************* PART ONE STUFF *************************/
/******************************************************************/

func calcPartOne(initialState *ferry) {
	if renderStates {
		renderState(initialState, "part_one(0).png")
	}
	state, changes := initialState.nextStateUsingAdjacents()
	iteration := 1 // we have already done an iteration
	for {
		if renderStates {
			renderState(state, "part_one("+fmt.Sprint(iteration)+").png")
		}
		state, changes = state.nextStateUsingAdjacents()
		if changes == 0 {
			log.Println("Part One solution:", state.countAllOccupiedSeats())
			break
		}
		iteration++
	}
}

func (f *ferry) nextStateUsingAdjacents() (newState *ferry, alterations int) {
	newState = &ferry{
		height: f.height,
		width:  f.width,
		cells:  make(map[coords]rune, f.width*f.height),
	}
	for x := 0; x < f.width; x++ {
		for y := 0; y < f.height; y++ {
			currCoords := coords{x, y}
			switch f.cells[currCoords] {
			case floor:
				newState.cells[currCoords] = floor
			case empty:
				if f.countAdjacentOccupiedSeats(currCoords) == 0 {
					newState.cells[currCoords] = occupied
					alterations++
				} else {
					newState.cells[currCoords] = empty
				}
			case occupied:
				if f.countAdjacentOccupiedSeats(currCoords) >= 4 {
					newState.cells[currCoords] = empty
					alterations++
				} else {
					newState.cells[currCoords] = occupied
				}
			default:
				log.Fatal("unknown cell type at", currCoords, ", ", f.cells[currCoords])
			}
		}
	}
	return newState, alterations
}

func (f *ferry) countAdjacentOccupiedSeats(c coords) int {
	var count int
	count += runeWeight[f.cells[coords{c.x - 1, c.y - 1}]]
	count += runeWeight[f.cells[coords{c.x - 1, c.y + 0}]]
	count += runeWeight[f.cells[coords{c.x - 1, c.y + 1}]]

	count += runeWeight[f.cells[coords{c.x + 0, c.y - 1}]]
	count += runeWeight[f.cells[coords{c.x + 0, c.y + 1}]]

	count += runeWeight[f.cells[coords{c.x + 1, c.y - 1}]]
	count += runeWeight[f.cells[coords{c.x + 1, c.y + 0}]]
	count += runeWeight[f.cells[coords{c.x + 1, c.y + 1}]]
	return count
}

/******************************************************************/
/************************* PART TWO STUFF *************************/
/******************************************************************/

func calcPartTwo(initialState *ferry) {
	if renderStates {
		renderState(initialState, "part_two(0).png")
	}
	state, changes := initialState.nextStateUsingVisibles()
	iteration := 1 // we have already done an iteration
	for {
		if renderStates {
			renderState(state, "part_two("+fmt.Sprint(iteration)+").png")
		}
		state, changes = state.nextStateUsingVisibles()
		if changes == 0 {
			log.Println("Part Two solution:", state.countAllOccupiedSeats())
			break
		}
		iteration++
	}
}

func (f *ferry) nextStateUsingVisibles() (newState *ferry, alterations int) {
	newState = &ferry{
		height: f.height,
		width:  f.width,
		cells:  make(map[coords]rune, f.width*f.height),
	}
	for x := 0; x < f.width; x++ {
		for y := 0; y < f.height; y++ {
			currCoords := coords{x, y}
			switch f.cells[currCoords] {
			case floor:
				newState.cells[currCoords] = floor
			case empty:
				if f.countVisibleOccupiedSeats(currCoords) == 0 {
					newState.cells[currCoords] = occupied
					alterations++
				} else {
					newState.cells[currCoords] = empty
				}
			case occupied:
				if f.countVisibleOccupiedSeats(currCoords) >= 5 {
					newState.cells[currCoords] = empty
					alterations++
				} else {
					newState.cells[currCoords] = occupied
				}
			default:
				log.Fatal("unknown cell type at", currCoords, ", ", f.cells[currCoords])
			}
		}
	}
	return newState, alterations
}

func (f *ferry) countVisibleOccupiedSeats(c coords) int {
	var count int
	count += runeWeight[f.findClosestSeat(c, -1, -1)]
	count += runeWeight[f.findClosestSeat(c, -1, +0)]
	count += runeWeight[f.findClosestSeat(c, -1, +1)]

	count += runeWeight[f.findClosestSeat(c, +0, -1)]
	count += runeWeight[f.findClosestSeat(c, +0, +1)]

	count += runeWeight[f.findClosestSeat(c, +1, -1)]
	count += runeWeight[f.findClosestSeat(c, +1, +0)]
	count += runeWeight[f.findClosestSeat(c, +1, +1)]
	return count
}

func (f *ferry) findClosestSeat(from coords, xDir int, yDir int) rune {
	if xDir == 0 && yDir == 0 {
		panic("provide a direction")
	}

	// Horizontal search
	if yDir == 0 {
		y := from.y
		for x := from.x + xDir; x >= 0 && x < f.width; x += xDir {
			cell := f.cells[coords{x, y}]
			if cell == floor {
				continue
			}
			return cell
		}
		return outOfBounds
	}

	// Vertical search
	if xDir == 0 {
		x := from.x
		for y := from.y + yDir; y >= 0 && y < f.height; y += yDir {
			cell := f.cells[coords{x, y}]
			if cell == floor {
				continue
			}
			return cell
		}
		return outOfBounds
	}

	// Diagonal search
	x := from.x + xDir
	y := from.y + yDir
	for {
		// bounds check
		if x < 0 || x >= f.width {
			return outOfBounds
		}
		if y < 0 || y >= f.height {
			return outOfBounds
		}

		cell := f.cells[coords{x, y}]
		if cell == floor {
			x += xDir
			y += yDir
			continue
		}
		return cell
	}
}

/****************************************************************/
/************************* RENDER STUFF *************************/
/****************************************************************/

// Thanks to https://yourbasic.org/golang/create-image/ for the example
func renderState(state *ferry, filename string) {
	w, h := state.width*8, state.height*8

	upLeft := image.Point{0, 0}
	lowRight := image.Point{w, h}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x := 0; x < state.width; x++ {
		for y := 0; y < state.height; y++ {
			cell := state.cells[coords{x, y}]
			drawCell(cell, x, y, img)
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		panic("could not create file: " + filename)
	}
	png.Encode(f, img)
}

func drawCell(cell rune, x, y int, img *image.RGBA) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			xPixel, yPixel := x*8+i, y*8+j
			switch cell {
			case floor:
				img.Set(xPixel, yPixel, lightColor)
			case empty:
				if freeSeatSprite[i+j*8] == '#' {
					img.Set(xPixel, yPixel, darkColor)
				} else {
					img.Set(xPixel, yPixel, lightColor)
				}
			case occupied:
				if occupiedSeatSprite[i+j*8] == '#' {
					img.Set(xPixel, yPixel, darkColor)
				} else {
					img.Set(xPixel, yPixel, lightColor)
				}
			default:
				log.Println("draw cell args:", cell, x, y)
				panic("not a valid cell type")
			}
		}
	}
}

var lightColor = color.RGBA{0xEE, 0xEE, 0xEE, 0xFF}
var darkColor = color.RGBA{0x33, 0x33, 0x33, 0xFF}

// possible optimization: turn both into [8]byte, where bit = 1 -> dark pixel
const freeSeatSprite = `          ####   #    #  #    #  #    #  #    #  ######         `
const occupiedSeatSprite = `          ####   ######  ######  ######  ######  ######         `

const outOfBounds = rune(0)
const floor = '.'
const empty = 'L'
const occupied = '#'

const exampleData = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

const data = `LLLLLLLL.LLLLLLLLLLLLL.LLLLL.LLLLLLLL.LLLLLLLLLLLLLLLL.LLLL..LLLLLLL..LLLLLL.LLLLLLLL.L.LLLLLLLLLLL
LLLLL.LL.LLLLLLL.LLLLL.LLLLL.LLLLLLL..LLLLLLLLLLLLLLLL.LLLLLLLLLL.LL.L.LLLLL.LLL.LLLL..LLLLLLLLLLL.
LLLLLLLL.LLLLLLLLL.LLL.LLLL..LLLLLLLL.LLLLLLLLL.LLLLLLLLLLLL.LL.LLLL.LLLLL.L.LLL.LLLLLLLLLLL.LLLL.L
L.LLLLL..LLLLLLL..LLL.LLLLLL.LLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLL.LLLLLLLLLLLLL
LLLLLLLLLLLLLLL...LLLLLLLLL..L.LLLLLLLLLLLLLLLL.LLL.LL.LLLLL.LLLLLLLLLLLLLLL.LLLLLLLL.LLLLLL.LLLLLL
LLLLLLLL.LLLLLLL.LLLLL.LLLLL.LLLLLLLLL.LLLLLLLL.LLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLL.LLLLLLLL..LLLLL
LL.LLLLL.LLLLLLL.LLLLL.LLLLL.LLLLLLLLLL.LLLL.LL.LLLLLL.LLL.LLL.LLL.LLLLL.LLL.LLLLLLLL..L.LLLLLLLL.L
LLLLLLLL.LLLL.LLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLL..LLL.LL.LLLLLLL.LLLLLLLL..LLLLLLLLLLLL
LLLLLLLL.LLLLLLL.LLLLL.LLLLL.LL.LLLLL.LLLLLLLLL.LLLLLL.LLL.L.LLLLLLLLLLLLLLL.LLLLLLLL.LLLLLL.LLLLL.
.L.LL..L.L...L.LL....L...LL......LL..L.L...L.L.LLLLLL..L...LLL..LL..L..............L.L..L.......LL.
LLLLLLLL..LLLLLL.LLL.L.L.LLL.LLL.L.LLLLL.LLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLL.LLLLL.LLLLLL.LLLLLL
LLLLLLLL.LLL.LLL.LLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLL..LLLLL.LLLLL.L.LLLLLL.LLLLLLLLL.LLLLLL..LLLLL
LLLLLLLL.LLL.LLL.L.LLL.LLLLL.LLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLLLL.LLLL.LLLLLLLLLLLLLLLL.LLLLLLLLLLLLL
.LLLLLLL.LLLLLLL.LLLLL.LLLLL.LLLLLLLL.LLLLLLLLLLLLLLLL.L.LLLLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLLLLLL
LLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLL.LLLLLLLLL.LLLLLL.LLLLL.LLLLLLLLLLLLLLL.LLLLLLLLLLLLLLL.LLLLLL
LLL.LLLL.LLLLLLL.LLLLLLLLLLL.LLLLLL.L.LLLLLLLLL..LLLLL.LLLLL.LLLL.LL..LL.LLL.LLLLLLLLLLLLLLLLLL.LL.
........LL...L..L....L.LLLLL..L..LL........L.L.L.LL.L.LL..L.LLL.L....LL........L..LLLL.LLL....L....
LLLLLLLL.LLLLLLL.LLLLLLLLLLL.LL.LLLLL.LLLLLLLLL.LLLLLLLLLLLL.LLLLL.L.LLLLLLL.LLLLLLLLLLLLLLL.LLLLLL
LLLLLLLL.LL.LLLLLLLLLLLLLLLL.LLLLLLLL.L.LLL.LLL.LLLLLLLLLLLL.LLLLLLL.LLLLL.L.LLLLLLLL.LLLLLL.LLLLLL
L.LLLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLLLL..LLLLLL.L.LLL.LL.LLLLL.LLLLLLL.LLLLLLL.LLLLL.LLLLLLLLLLLLLLLL
LLLLLLLLL..L.LLLLLLLLLLLLLLLLLLLLLLLL.L.LLLLLLL.LLLLLL.LLLLLLLLLLLLL.LLLLLL..LLLLLLLL.LLLLLL.LLLLL.
LLLLLLLLLLLLLL.L.LLLLLLLLLLL.LLLLLLL..LLLLLLL.L.LLLLLLLLLLLL.LLLLLL..LLLLLLLL.LLLLLLL.LLLLLL.LLLLLL
LLLLLLLL.L.L.LLL.LLLLL.LLLLL.LLLLLLLL.LLLLLL.LL.LLLLLL.LLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLL
.........L.LL...LLL.....LL..LL.....L.L.L.......L..L......LLLL.L...L.L...L.....LL.........L.LL.L..LL
L.LLLLLL.LLLLLLL..LLLL.LLLLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLL.LLLLLL..LLLLLL.LLLLLLLL.LLLL.LLLLLLLL
LLLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLLLLL.LLLLL.LLL.LLLL.L.LLLLLLLLLLLLL.LLLLLLL.LLLLLLL.LLLL.LLLLLLLLL
LLLLLLLL.LLLLLLL..LLLL.LLLLLLLLLLL.LL.LLLLLLLLL.LLLLLL.LLLLL.LLLLLLLLLLLLLL..LLLLLLLL.LLLLLL.LLLLLL
LLLLLLLL.LLLLLL..LLLLLLLLLLL.LLLLLLLLLLLLLLLLLL.L.LLLLL.LLLL.LLLLLLL.L.LLLLL..LLLLLLL.LLLL.L...LLLL
L.LLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLLL.LLL.LLLLLLL.L.LLLLLLLLLL.LLLLLLL..L.LLLL.LLLLLLL.LLLLLLL.LLLLLL
LLLLL.LL.LLLLLLL..L.LL.LLLLL.LL.LLLLL.LLL.LLLLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLL.LLLLLLLL.LLLLLL.LLLLLL
LL.LLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLLLLLLLLLLL.LLLL..LLLLLLL.LLLLLLLLLLLLL.LL.LLLLLLLLLLLLL
..L.......L...........LL...L.L....L.........LL.L.LLL.L.........L..LL....L.L....LL.L......LL....L.L.
LLLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLLLLL.LL.LLLL.L.LL..LLLLLL.LL.LLLLLL.LLLLLL..LLLLLL.L.LLLLLL.LLLLLL
LL.LL.LL.LLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLL.L.LLLLLL.L.LLL.LLLLLLL.LLLLLLL.LLLLLLLL.LLLLLL.LLLLLL
LLLLLLLL.L.LLLLL.LLLLL.LLLLLLLLLL.LLL..LLLL.LLL.LLLLLL.LLLLL.LLLLLLL.L.LLLLL.LLLLLLLL.LLLLLLLLLLLLL
LLLLLLLLLLLLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLL.LLLLL.LLLLLLL.LLLLLLL.LLLLLLLLLL.LLLL.LLLLLL
LL.LLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLLLL.LLLLL.LLLLLL.LLLLLL..LLLLLLL.LLLLLLLL.LLLLLLLLLL.LL
L........L.L.L...LL.LLL...L........LL..L...L..LL..LL.L.L....L..L..L.L....L.L.....LLL.L.LL...L...L.L
LLLLLLLL.L.LLLLL.LLLLLLLLLLLLLLLLLLLL.LLLLLLLL..LL.LLL.LLLLL.LLLLLLL.LLLLLLL.LLLLLLLL.LLLL.L.LLLLLL
LLLLLLLL.LLLLLLL.LLLLL.LLLLL.LLLLLL.LLLLLLLL.LL.LLLLLLLLLLLLLLLLLLLL.LLL.LLL.LLLLL..LLLLLLLLLLLLL.L
LLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLL.L.LLLLL.LLLLLLLLLLLLLL.LLLL.LL.LLLLLLLLLLLLLLLL.LLLLLLLLLLLLL
.LLLLLLL.LLLLLLLLLLLLLLLLLLL.LLL.LLLLLLLLLLLLLL.LLLLLLLLLLLL.LLLLLLL.LLLLLLL..LLLLLLL.LLLLLLLL.LLLL
LLLLLLLLLLLL..LL.LLLLLLLL.LL.LLLLLLLL..LLLLLLLL.LLLLLL.LLLLLLLLL.LLL.LLLLLLL.LLLLLLLL.LLLLLL.LLLLLL
.L..L......L.L.LL..L...LL.L...L....L.L.....LLLL...L.L...LL.....L......L..L.L.L.....L.L......L....LL
LLLLLLLLLLLLLLLLL.LLLL.LLL.L.LLLLLLLL..LL.LLLLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLL.LLLLLLLL.LLLLLL.LLLLLL
LLLLLLLLLLLLLLLLLLLLLLLLL.LL.LLLLLLLLLLLLL.LLLL.LLLLLL.LLLL..LLLLLLLLLLLLLLL.LLLLLLLLLLLLLLL.LLLLLL
.LLLLLLL.L.LLLLLL.LLLLLLLLLL.LLLLLLLLLLLLLLLLLL.LLLLLL.LLLLL.LLLLLLLLLLLLLLL.LLLLLL.L.LLLLLLLLLLLLL
LLLLLLLLLLL.LLLLLLLLLL.LLLLL.LLLLLLLL..LLLLLLLL.LLLLLL.LLLLL.LLLLLLL.LLLLLLL.LLLLLLLLLLLLLLL.LLLLLL
LLLLLLLL.LLLLLL.LLLLLL.LLLLLLLL.LLLLL.LL.LLLLLL.LLLLLL.LLLLL.LLLLLLL.L.LLLLLLLLLLLLLL.LLLLLL.LLLLLL
LLLLLLLL.LLLLLLLLLLLLL..LLLLLLLLLLLLLLLLLLLLL.L.LLLLLLLLLLLL.LLLLLLLLLLL.LLLLLLLLLLLL.LLLLLLLL.LLLL
LLLLLLLLLLLLL.LL.LLLLLLLLLLL.LLLLL.LL.LLLLL.LLL.LLLLLL.LLLLLLLLLLL.LLLLLLLLLLLL.LLLLL.LLLLLLLLLLLL.
LLLL.LLL.LLLLLLL.LLLLL.LLL.LLLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLL.LLLLLLL.LLLLLL..LLLLLLLL..LLLLLLLLLLLL
LLLLLLLL.LLLL.LL.LLLLL.LLLLLLLLLLLLLL.LLLLLLLLL..LLLLL.L.LLL.LLLLLLL.LLLLLL.LLLLLLLLLL.LLLLL.LLLLLL
.....LL.L..L..L..LL.............LL...L..L.LLLLL.L.L...L.....L...L.....L....LL.L..L.....LLL.........
LLLLLLLLLLLL.LLL.LLLLLLL.LLLL.LLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLLL.LLLLLLL.LL.LL.LL.LLLL.L.LLLLLL
LLLLLLLLLL.LLL.L.LLLLLLLLLLL.LLLLLLLL.LLL..LLLL.LLLLLLLLLLLL.LLLLLLL.LLLLLLL.LLLL.LLL.LLLLLLLL.LLLL
LLL.LLLL.LLL.LLL.LLLLL.LLLL..LLLLL.LL.LL.LLLLLL.LLLLLL.LLLLL.LLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLLLLLL
LLLLLLLL.LLLLLL..L.LLLLLLL.L.LLLLLL.L.LLLLLLLLL.LLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLL.LLLLLL
LLLL.LLLLLLLLLLL.LLLLL.LLLLL.LLLLLLLL.LLLLLL.LL.LLLLLL.LLLLL.L.LLLLLLLLLLLLLLLLL.LLLL.LLLLLL.LLLLLL
LLLLLLLL.LLLLLLLLLLLLL.LLLLL.LLLLLLLL.LLLLLLL.L.LLLLLL.LLLLL.LLLLLLLLLLLLLLLLLLLLLL.L.LLLLLL.LLLLLL
LLLLLLLL.LLLLLLL.LLLLLLLL.LLLL.LLLLLLLLLLLLLLLL.LLLLLLLLLL.L.LL.LLLL.LLLLLLLLLLLLLLLL.LLLLLL.LLL.LL
LLLLLLLL.LLLLLLL.LLLLLLL.LLLLLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLL.LL.LLLLL.LLLLLLLLLLLLL
LLLLLLLL.LLLLLLLLLLLLL.LLLLL.LLLLLLLL.LLL.LLLLL.LLLLLLLLLLLL.LLLLLLLLLLLLLLL.LLL.LLLLLLLLLLL.LLLLLL
..LL..LL..LL......L....L....L....L....L..L...........L.L...L.L.L...LL.L..L..........LLLL.L..L.L....
LLLLLLLL.LLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLL.L..LLL.LLL.LLLLLLLL..LLLLL.LLLLLL
LLLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLLLLLLLLLLLLLLL.LLLLLLLLLLLL.LLLLLLL.LLLLLLL.LLLLLLL..LLLLLL..LLLLL
LLLLLLLL.LLLLLLLLLLLLLLLLLLL.LLL.LLLL.LL.LLLLLL.LLL.LL.LLLLLL.LLLLLL.LLLLLLL.LLLLLLLL.LLLLLLLLLLL.L
LLLL.LLL.LLLLLLL.LLLLLL.LLL..L.LLLLLL.LL.LLLLLLLLL.LLLLLLLLL.L.LLLLLLLL.LLLL.LL.L.LLL.LLLLLL.LLLLLL
LLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLL.LLLLLLLL.LLL.LLLL.L.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLL
L..LLL.L....L...LL..LLLLL..L...LL.L..L.L...............L.LL.L.....LLL.LL...L..L.LL..L..L...L.......
LLLLLLLL.LL.LLLL.LLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLL.L.LLLLLL.LLLLLLLLL.LLL
LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLLLLLL
LLLLLLLL.LLLLLLL.LLLLL.LLLLLLLLLLLLL..LLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLL.L.L.LLLLLLL.LLLL....LLLLL
LLLLLLLLLLLLLLLL.LLLLL.L.LLLLL.LLLLL..LLLLLLLLL..LL.LL.LLLLL.LLLLL.LL.LLLLLLLLLLLLLLLLLLLLL..LLLL.L
LLLLLLLL.LLLLLLL.LLLL..LLLLL.LLLLLLLLLLLLLLLLLL.LLLLLL.LLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLLLL.LLL.LLLLL
L.......L....L....L.LL..L.LL...L.L.LL........L......L.......LL.L.L...L...L...LL.L.L.L..............
LLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLL.LLLLLLLLL.LLLLLLLLLLL..LLLL.LL.LLLLLLL.L.LLLL.L.L.LLLL.LLLLLL
LLLLLLLLLL.LLLLL.LLLLL.L.LLL..LLLL.LLLL.LLLLL.LLLLLLLL.L.LLL.LLLLLLL.LLLLLLLLLLLL.LLL.LLLLLLLLLLLLL
LLLLLLLL.LLLLLLLL.LLLL.LLLLLLLLLLLL.L.LLLLLLLLL.LLLLLL.L.LLLLLLLLLLL.LLLLLLL..LLLLLLL.LLLLLLLLLLLLL
LLLLLLLL.L.LLLL..LLLLL.LLLLL.LLLLLLLLLLLLLLLLL..LLLLLL.LLLLL.LL.LLLL.LLLLLLL.LLLLLLLLLLLLLLL.LLLLLL
LLLLLLLL..LLLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLL.LLL.LLLLLL.LLLLL.LLLLLLL.LLLLLLL.LLLLLLLL.LLLL.L..LLLLL
LLLLLLLL.LLLLLLL.LLLLL.LLLLLLLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLL.LLLLLLL.L.LLLLLLLLL.LLLL.LLLLLL.LLLLLL
LLLLLLLL.LLLLLLL.LLLLL.LLLLLLLLLLLLLL.L.LLLLLLLLLL.L.L..LLLL.LLLLLLL.LLLLLLL.LLLLL.LLLLLLLLL.LLLLLL
LLLLLLLL.LLLLLLL.LLLLL.LLLLLLLLLLLLL..LLLLLLLLL.LLLLLL.LLLLL.LLLLLLL.LLLLLLL.LLLLLLLLLLLLLLL.LLLLLL
..LL.....L....L...L.LL........LL.L........L....L....L.L....LLL.L....LLLLLL.L.L.L..LLL.L.L.L..LL..LL
LLLL.LLL.LLLLLL.LLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLL.LLLLLLLL.LLLLL.LLLLLLL
LLLLLLLL.LLLLLLL.LLLLL.LLLL..LLLLLLLL.LLLLLL.LL.L.LLLL.LLLLL.LLL.LLL.LLLLLLL.LLLLL.LLLLLLLLL.LLLLL.
LLLLLLLLLLLLLLLLLLLL.L.LLLLLLLLLLLLLL.LLLL.LLLL.LLLLLL.LLLLL..LLLLLL.LLLLLLL.LLLLL.LL.LLLLLL.LLLLLL
LLLLLLLL.LLLLLLLLLLLLL.LLLLL.LLLLLLLL.LLLLLLLLL.LLL.LL.L.LLLLLLLL.LLLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLL
LLLLLLLL.LLLLLLL.LLLLL.LLLLLLLLLLLLLL.LLLL.LLLL.LLLLLL.LLLLL.LLLLLLL.LLLLLLLLLL.LLLLL.LLLLLL.LLLLLL
LLL.LLLLLLLLLLLL.LLLLL.LLL.LLLLLLLLLLLLLLLLLLLLLLL.LLL.LLLLL.LLLLLLL.LLLLLLL.LLLL.LLL.LLLLLL.LLLLLL`
