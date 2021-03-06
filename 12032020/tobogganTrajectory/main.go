package main

import (
	"fmt"
	"strings"
)

func main() {

	splitLines := strings.Split(input, "\n")

	offset := 0
	treesEncountered := 0
	for i, line := range splitLines {
		if i % 2 != 0 {
			continue
		}
		line = strings.TrimSpace(line)
		if line[offset % len(line)] == '#' {
			treesEncountered++
		}
		offset +=1
	}
	fmt.Printf("Tress Encountered: %d\n", treesEncountered)

}

var input = `..##.........##.........##.........##.........##.........##.......
#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
.#....#..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
.#...##..#..#...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
..#.##.......#.##.......#.##.......#.##.......#.##.......#.##.....
.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
.#........#.#........#.#........#.#........#.#........#.#........#
#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...
#...##....##...##....##...##....##...##....##...##....##...##....#
.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#`

//var input = `.#....#..##.#..####....#.......
//......#..#....#....###......#.#
//#..#.....#..............##.#.#.
//#.#...#....#...#......##..#..#.
//...#..#.##..#..#........###.#.#
//...#.#..........#.........##...
//...#.#....#.#....#..#......#...
//..##.#.....#.......#.###..#..##
//..#.......#.......#....##......
//....##........##.##...#.###.##.
//#.......#.......##..#......#...
//..##.............##.#......#...
//...#.####....#.....#...##......
//.............##.#......#.......
//..#...#....#......#....#.......
//..#....#..#............#.......
//##...#..#........##..#......#..
//##........##........#.#.......#
//#.......#........#.#..#....###.
//.....#..#.#..........##...#....
//..##...#......#.#...#..#...#...
//##.##...#......#....#....#...#.
//#.......#..#.#..#....#..###.#.#
//#.............#.#....#..#.#....
//...#.......###.#.##.##.#...#..#
//.##.......##..##...#..###......
//.......#.#.#.#..####..#..#..#..
//...##......#.#.##.###....#.###.
//###......###......#.#####..#...
//..#........##..#..##.##..#...#.
//.....##..#...#..#.##.....#.#...
//#......#.##....#..##.#....#.#..
//##.#.##..#................#....
//......#.#....#......##.....#...
//..#...##..#..#...#..#.#..#.....
//........#.#.#.##...#.#.....#.#.
//#.#......#.....##..#...#.......
//..#.#......#...........###.....
//......##....#....##..#..#.#.#.#
//##....#.###...#......#..#...#..
//#.#.##....###...####.......#..#
//##...........#.....#........#.#
//.##.#..#.....#......#.......#..
//##..##..###....#.........##....
//..#..#..#.##...#.#...#........#
//#.##.###...#.......#...........
//.........#.................#...
//#.##...#.....#..##........#....
//....#.#...##...#...........#...
//.#.....#.#..#...##..##.....#...
//.#.....####....#..##..#........
//...#..#......##.#..##.#.#.#..#.
//.##.#.....#.....#...#.......##.
//......#..#..#......#...####....
//.......#......##..#..##.....#..
//......#.#..#...#..#.#..........
//....##.........#...............
//.#....#..##.....#....##.##.....
//#.#.....#...#....####....#.#...
//#.....#....#.#...#.............
//...#..#.....#....##..#..#......
//...#.#............#...........#
//###.#..#.#......#.....##.....#.
//####....#....###.....#..#.#####
//.###..#...#.#..#......##.#.#.#.
//.....#.##.#....#..##....#..#..#
//...#....#...##.....#......#.#..
//....#...#....#...#....#.....#.#
//.#.....#.....#.#..#......#..#..
//..#..##....##.##....#.....##...
//#..##...#.##...#..#.#.#.....#..
//...#..##.#..#....#.#....######.
//..........#..#.....#....#...##.
//#.#####.#.###..#.....#.........
//#....#......#..#.#.##.##..###..
//..#...###.#.#....##.##...##....
//.......#....#..#...##......#...
//...#.#...#..#.....#..##.#......
//###..##...........#............
//..#....#.##....#.#..##...#.....
//##....#...#....#.....#.#..##...
//..............#.##.#..#..##.###
//......#..#..#..#.#....###...##.
//.#...#..#.#.#....#..........#..
//..#.#.....#..#...........#.##..
//...#.#......#......#..#..#.#...
//...#....#.#.#.....#...#.##..#..
//.#.#..#...#........##.......#..
//##..........#..#...#....###.#..
//#.....###......#..#.#.#....#.#.
//..###.......#.#...............#
//#....#.....#.#......#..##.##...
//#.##.#.#....#..#.#...#.#...#..#
//#....#..#...........#.......#..
//...#.####.....#.........###.##.
//......#..#.....#..........#..#.
//#...#.#..####...#...#.#.##...##
//.##.........#......#.#.#.......
//.......##...##.##....###...##..
//...#..#....#..#.#.#.....#.#....
//#....#.#.#.......##..###..##...
//......#............#.#...#..#..
//#.#.....#......#...#.......###.
//...#.#................#...#....
//.....#......#.#..#...##.#.#...#
//#....#.#..#..#..##..#.##..#....
//#.................#..#....#....
//..#....#.......####....###.....
//.#..#.#.#...###..#...#..###....
//..#..#.#......#.###..........#.
//.....#......#.......##....##.#.
//.#....#........#.#.##.#........
//#.#..##..#..#.#...####....##...
//...#....#.#..#...#..........#..
//.#.....#.##....#...##..........
//....##....#.....#.....#...#.###
//.#...##.#.#..##..#...#.#..#..#.
//..#.......#.##.....#.#........#
//...#...#.....##..#.#.#....#....
//...#.....#.......##.........#.#
//.##.....#..#.#...#.#...#.#...#.
//...........#...#.###..#...#..#.
//#.##........#..###.##...####...
//.#.....#.#...##...#..#..#...##.
//..#....#..#...#.....#.....##...
//..###..#.#.....##........#.##..
//.#.#..##........#.##....#..#.##
//.####.#..##..#.#..#....##....#.
//.##....##...#.#........#.......
//....#..#..#...#..#..#..#.#.....
//...#......................#....
//#.....#.#....#..#..#.#..#....#.
//##.....#.....##..........###...
//.#..#..............#...##.#.#.#
//...#...#.#.............#.#..#.#
//..#.....#.......#......#.#.....
//.###.#..#..#..#.#..#.....#.....
//.....##..##...##.......#....###
//.#........###...##..#....##....
//#....#.#......##..#....#.##..#.
//#....#.#...#........##...###...
//.#.....#...#.###....#.##.#.####
//###......#....#...#....##..#..#
//##....#..###......#...#.#.#....
//..........#......##.#..#.......
//.#..#......###.........##...#..
//....#......#....#.........#.#.#
//##.#.#...#.#...#...#..#......#.
//....#.##.........#..#.....##.#.
//........#...#..#.#.#.#.....##..
//..#......#.#.#..#.....##.......
//..............#....#....##.#..#
//....#.#.....#....#.#.###.#....#
//..#..........#..#......#.##..#.
//...#...#.#.............#..#....
//#.......#..#..##.........##..#.
//..##..#............#.....#.....
//....#.#..##...#.#..#.........#.
//........#.......#.##....#....#.
//...#.....#.#.....#.#....#......
//..#......##.#.............#.#.#
//#.#.............#.#.....#......
//.##....##.#.....#....#...##....
//.#.#....##....#.....##.........
//...#.....#.....#.....#..#.###..
//.......#....#...##.#...#...#..#
//..#.#.......#...###.#...#....#.
//.....###..##....###.#.##.......
//....#..................##.#.##.
//.#.......###.##...#.#.....#....
//....#....##...##.....#.#...#..#
//#..#.....#......##...#....#....
//#..##.........#.....#...#......
//...#..##.......##......#..#####
//.#..###.###.#.##........#......
//.#...#....#....#.#....#...##...
//##........#....#.........##..#.
//.#.....##............#.#.......
//....#....#...........###.....##
//.#......#.#.#..#....#.#.....##.
//......#.##.#..##....#.#.#..#...
//#....#......#...#..####........
//......#..#..##..#.......#.#..#.
//##....##.###.#...#.##.#..#.###.
//.#.........#...##...#.#......#.
//..#.#...........####.#....##.##
//.....#.#..##.#...###...#..#.#..
//...#..#..##.#...#.....#.##...##
//..##......##..........#..###...
//.#......##.....#.##....#.#.##.#
//...#.......##..##.....#....#...
//.##...#...#....#..#............
//#..#....#...........#..........
//......#...#.#.......#...#.##..#
//..#.###..#.....#.....#..#.....#
//....#....#..........##....#..#.
//.......##.#.#.#......#....#...#
//####..#.###........#..#......#.
//#........##.#.#.#.............#
//.#......#......#..#.##.....#...
//.....##.##......##.#.....#.#.#.
//.......##.#.....##.......#.#.#.
//.#.#..#.#..#.##...#.#....#.#..#
//.#..##....#..#...##.......#..#.
//.#.#..#.......#................
//#........#.#.#......#.#.#.#....
//##......#...#..##.#...##.##....
//##.#..#...........##...#..###..
//......#.####...#........#.#.#..
//...#........##..###.#.#...#...#
//.#.....##..#......##......###..
//..#.#...#......#..#..##.#.....#
//#....#..#.#..........#...#.....
//.#......#.##..###..#.#....#..##
//.......#.......#..#..#......#..
//..##.....##.#..#..#.....##.....
//........#.##...#.#.#..#..#..#..
//...#.######.........#.....#..##
//.#.#............#....#.........
//#...#....###.#......#.#........
//#.........#....#...##..........
//....#...........#.###.#...###..
//.........#........#.#.#..#...#.
//.#.......#.#.....#.#.....#.##..
//.....#.......#.....#.#.#.......
//#.##..#..##.......#...#......#.
//.###.....##...##.#...##.##.#.#.
//...#......##..##............#.#
//...#......................#..##
//#..#..#................#...#...
//#..#....#.#.#...##.......#..#..
//....#.#..###.##...#..#.###..#..
//..#...#....####.#............#.
//......#....#.#...#.#.#.........
//#...#........#.....##..###.#..#
//#....#...#...##...#..#....##...
//#..#...#..#.......#.#..##.#..#.
//#.#..........#...........##....
//.#...###...#......#.......#.#.#
//.........#.........#...#...##..
//##.#..###......##..#.....#..#..
//....##...............#.....#...
//.....#.....###.#.....#.#.......
//....#..#......###..#..##..#....
//......................#.....#..
//..#..#...##....##....#........#
//..#....#...#...#.......#.....#.
//...##.#.#.##......#.#.#.#.####.
//.###...#..#......#.#..#........
//#..#...##.#..#..##..##....#...#
//...#...#..#..#..#........#..##.
//.##....#.##.#....#...#.#.#....#
//#..#....#..#....#.......##..#.#
//...#.#....####...........#...#.
//#...#####...#.#..#......#...#.#
//.##....#.#.#..#......#..##.....
//..........#..#.#.#.....##......
//.....#....#..................#.
//.........#...#...#....#..###...
//.#.#.#....#....................
//......##............##.###..#..
//#.#...#........####.##..#.#.##.
//#........#.#.#.#..#.##.....#...
//......####..#.##.......#....#..
//.........#...#...#.....#.......
//..##.....#...#...#.....##.....#
//....#...##....#.....#..#..##.##
//..#.........##...##..###..#....
//#....#.#.........##.###.#...##.
//.##...#....#..#..#.#....##.....
//##..#..#..#...........#.##....#
//....#..........#...#..#.....#..
//........###..#..#.#.#.....##...
//#...#...#..###............###..
//..#.....#.#.#..#..#.#..#......#
//..#...##..#....#...#......#....
//#....#........##.....#..##....#
//#.....#.#.#..#.......##.#.#.##.
//..##...#...#.....#..........#..
//##.....#....#......#..........#
//......#..#..........#.#..####..
//......#...#............##...##.
//..#.......##.......#...###.###.
//.#..#.#.#...#..##.#......#.#...
//.##.....##.#.#...#.##.........#
//#.#.######...........#.#####.#.
//........#.##...##....##.#.##.#.
//....#......#.....#.....###...##
//#..............#.#....#.#....#.
//....#..###.#.........##.#.#....
//..#.#.#..##....####..........#.
//...#..#.......#................
//...#....#..............#....#..
//.....#...#...#....#.#.#..#...#.
//......##.............###.##..##
//.#...#.#..#......#..#.##.......
//##.....#.....#.##...#....#.....
//..#..#.#.#.#.#..........#..###.
//##..........#........#....#.#..
//.....#...#........#.#..###....#
//.###.#........#.##......#.#...#
//#...##....#....#....##.#.#.....
//.....#.#............#..........
//..#.##....................#....
//.....#..#..#.#..#.##.......#...
//.....###......#......##......##
//#.....#.#.......##.......#...#.
//.#.#...#......#..###...#.....#.
//#.#..#...#..##.....#...#.#..#..
//.....#.#..........#..#.........
//.###..##..##.....#...#...#..##.
//#...#.#....#.......##..#.......
//###...#.#.#..#.......#......#..
//....##........#..........##....
//............#....#...........#.
//#..#.#....##..#.#..#......##...
//.###....##...#....##..........#
//.###........#........###.....#.
//...#...#.#......#...#....#.....
//.###.......#.........#.........
//....##.#......#...###......##.#
//.###...#..##.....##.......#....
//.#.#...#..#.##....#........#...`
