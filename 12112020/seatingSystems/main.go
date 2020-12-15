package main

import (
	"fmt"
	"strings"
)

func main() {

	matrix := make([][]string, 0)
	splitInputs := strings.Split(input, "\n")
	for _, splitInput := range splitInputs {
		splitInput = strings.TrimSpace(splitInput)
		cols := make([]string, 0)
		splitChars := strings.Split(splitInput, "")
		for _, char := range splitChars {
			char = strings.TrimSpace(char)
			cols = append(cols, char)
		}
		matrix = append(matrix, cols)
	}

	partTwo(matrix)

}

func partTwo(matrix [][]string) {
	var newSeating [][]string
part2:
	for {
		hadChanges := false
		newSeating := make([][]string, len(matrix))

		for i := 0; i < len(matrix); i++ {
			newSeating[i] = make([]string, len(matrix[i]))
		inner2:
			for j := 0; j < len(matrix[i]); j++ {
				left, right, up, down, topLeft, topRight, downLeft, downRight := false, false, false, false, false, false, false, false
				curr := matrix[i][j]
				if curr == "." {
					newSeating[i][j] = curr
					continue inner2
				}
				left = walkLeft(matrix, i, j)
				right = walkRight(matrix, i, j)
				up = walkUp(matrix, i, j)
				down = walkDown(matrix, i, j)
				topLeft = walkTopLeft(matrix, i, j)
				downRight = walkDownRight(matrix, i, j)
				topRight = walkTopRight(matrix, i, j)
				downLeft = walkDownLeft(matrix, i, j)
				if curr == "L" {
					if !left && !right && !up && !down && !topLeft && !downLeft && !downRight && !topRight {
						newSeating[i][j] = "#"
						hadChanges = true
					} else {
						newSeating[i][j] = curr
					}
				} else {
					countAdjacentSeats := 0
					if left {
						countAdjacentSeats++
					}
					if right {
						countAdjacentSeats++
					}
					if up {
						countAdjacentSeats++
					}
					if down {
						countAdjacentSeats++
					}
					if topLeft {
						countAdjacentSeats++
					}
					if downLeft {
						countAdjacentSeats++
					}
					if downRight {
						countAdjacentSeats++
					}
					if topRight {
						countAdjacentSeats++
					}
					if countAdjacentSeats >= 5 {
						newSeating[i][j] = "L"
						hadChanges = true
					} else {
						newSeating[i][j] = curr
					}
				}
			}
		}
		if !hadChanges {
			break part2
		}
		matrix = newSeating

	}
	countOccupiedSeats := countOccupiedSeats(matrix)
	fmt.Println(countOccupiedSeats)
	fmt.Println(newSeating)
}

func partOne(matrix [][]string) {
	var newSeating [][]string
part1:
	for {
		hadChanges := false
		newSeating := make([][]string, len(matrix))

		for i := 0; i < len(matrix); i++ {
			newSeating[i] = make([]string, len(matrix[i]))
		inner:
			for j := 0; j < len(matrix[i]); j++ {
				left, right, up, down, topLeft, topRight, downLeft, downRight := "", "", "", "", "", "", "", ""
				curr := matrix[i][j]
				if curr == "." {
					newSeating[i][j] = curr
					continue inner
				}
				if j != 0 {
					left = matrix[i][j-1]
				}
				if j < len(matrix[i])-1 {
					right = matrix[i][j+1]
				}
				if i != 0 {
					up = matrix[i-1][j]
				}
				if i < len(matrix)-1 {
					down = matrix[i+1][j]
				}
				if i != 0 && j != 0 {
					topLeft = matrix[i-1][j-1]
				}
				if i < len(matrix)-1 && j < len(matrix[i])-1 {
					downRight = matrix[i+1][j+1]
				}
				if i != 0 && j < len(matrix[i])-1 {
					topRight = matrix[i-1][j+1]
				}
				if i < len(matrix)-1 && j != 0 {
					downLeft = matrix[i+1][j-1]
				}
				if curr == "L" {
					if left != "#" && right != "#" && up != "#" && down != "#" && topLeft != "#" && downLeft != "#" && downRight != "#" && topRight != "#" {
						newSeating[i][j] = "#"
						hadChanges = true
					} else {
						newSeating[i][j] = curr
					}
				} else {
					countAdjacentSeats := 0
					if left == "#" {
						countAdjacentSeats++
					}
					if right == "#" {
						countAdjacentSeats++
					}
					if up == "#" {
						countAdjacentSeats++
					}
					if down == "#" {
						countAdjacentSeats++
					}
					if topLeft == "#" {
						countAdjacentSeats++
					}
					if downLeft == "#" {
						countAdjacentSeats++
					}
					if downRight == "#" {
						countAdjacentSeats++
					}
					if topRight == "#" {
						countAdjacentSeats++
					}
					if countAdjacentSeats >= 4 {
						newSeating[i][j] = "L"
						hadChanges = true
					} else {
						newSeating[i][j] = curr
					}
				}
			}
		}
		if !hadChanges {
			break part1
		}
		matrix = newSeating
	}
	fmt.Println(newSeating)
}

func countOccupiedSeats(matrix [][]string) int {
	countOccupiedSeats := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "#" {
				countOccupiedSeats++
			}
		}
	}
	return countOccupiedSeats
}

func walkLeft(matrix [][]string, i, j int) bool {
	if j == 0 {
		return false
	}
	for {
		j--
		if matrix[i][j] == "#" {
			return true
		}
		if matrix[i][j] == "L" {
			return false
		}
		if j <= 0 {
			break
		}
	}
	return false
}
func walkRight(matrix [][]string, i, j int) bool {
	if j >= len(matrix[i])-1 {
		return false
	}
	for {
		j++
		if matrix[i][j] == "#" {
			return true
		}
		if matrix[i][j] == "L" {
			return false
		}
		if j >= len(matrix[i])-1 {
			break
		}
	}
	return false
}
func walkUp(matrix [][]string, i, j int) bool {
	if i == 0 {
		return false
	}
	for {
		i--
		if matrix[i][j] == "#" {
			return true
		}
		if matrix[i][j] == "L" {
			return false
		}
		if i <= 0 {
			break
		}
	}
	return false
}
func walkDown(matrix [][]string, i, j int) bool {
	if i >= len(matrix)-1 {
		return false
	}
	for {
		i++
		if matrix[i][j] == "#" {
			return true
		}
		if matrix[i][j] == "L" {
			return false
		}
		if i >= len(matrix)-1 {
			break
		}
	}
	return false
}
func walkTopLeft(matrix [][]string, i, j int) bool {
	if i == 0 || j == 0 {
		return false
	}
	for {
		i--
		j--
		if matrix[i][j] == "#" {
			return true
		}
		if matrix[i][j] == "L" {
			return false
		}
		if i <= 0 || j <= 0 {
			break
		}
	}
	return false
}
func walkTopRight(matrix [][]string, i, j int) bool {
	if i == 0 || j >= len(matrix[i])-1 {
		return false
	}
	for {
		i--
		j++
		if matrix[i][j] == "#" {
			return true
		}
		if matrix[i][j] == "L" {
			return false
		}
		if i <= 0 || j >= len(matrix[i])-1 {
			break
		}
	}
	return false
}
func walkDownLeft(matrix [][]string, i, j int) bool {
	if i >= len(matrix)-1 || j == 0 {
		return false
	}
	for {
		i++
		j--
		if matrix[i][j] == "#" {
			return true
		}
		if matrix[i][j] == "L" {
			return false
		}
		if i >= len(matrix)-1 || j <= 0 {
			break
		}
	}
	return false
}
func walkDownRight(matrix [][]string, i, j int) bool {
	if i >= len(matrix)-1 || j >= len(matrix[i])-1 {
		return false
	}
	for {
		i++
		j++
		if matrix[i][j] == "#" {
			return true
		}
		if matrix[i][j] == "L" {
			return false
		}
		if i >= len(matrix)-1 || j >= len(matrix[i])-1 {
			break
		}
	}
	return false
}

//var input = `L.LL.LL.LL
//LLLLLLL.LL
//L.L.L..L..
//LLLL.LL.LL
//L.LL.LL.LL
//L.LLLLL.LL
//..L.L.....
//LLLLLLLLLL
//L.LLLLLL.L
//L.LLLLL.LL`

var input = `LLLLLLLL.LLLLLLLLLLLLL.LLLLL.LLLLLLLL.LLLLLLLLLLLLLLLL.LLLL..LLLLLLL..LLLLLL.LLLLLLLL.L.LLLLLLLLLLL
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
