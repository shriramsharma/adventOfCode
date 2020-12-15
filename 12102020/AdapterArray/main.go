package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {

	splitInput := strings.Split(input, "\n")
	//m := map[int]int{
	//	1: 1, 2: 0, 3: 1,
	//}
	nums := make([]int, 0)
	numsMap := make(map[int]bool)
	numsMap[0] = true
	for _, s := range splitInput {
		s = strings.TrimSpace(s)
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, n)
		numsMap[n] = true
	}
	sort.Ints(nums)

	paths := make([]int, nums[len(nums)-1]+1)
	paths[0] = 1
	for i := 0; i < len(nums); i++ {
		numWays := 0
		for j := 1; j <= 3; j++ {
			if numsMap[nums[i]-j] {
				numWays += paths[nums[i]-j]
			}
		}
		paths[nums[i]] = numWays
	}

	fmt.Println(paths[len(paths)-1])

	//comb := make([]int, 0)
	//final = make([][]int, 0)
	//cache := make(map[int][][]int)
	//adapterAllPossibleWaysPerf(numsMap, 0, comb, nums[len(nums)-1], cache)

	//fmt.Println(len(final))

}

//func adapterAllPossibleWays(numsMap map[int]bool, curr int, comb []int, end int, cache map[int][]int) []int {
//
//	if !numsMap[curr] && curr != 0 {
//		return comb
//	}
//
//	comb = append(comb, curr)
//	for i := curr; i < curr+3; i++ {
//		adapterAllPossibleWays(numsMap, i+1, comb, end, cache)
//	}
//	if curr == end {
//		final = append(final, comb)
//	}
//	return comb
//}

//func adapterAllPossibleWaysPerf(numsMap map[int]bool, curr int, comb []int, end int, cache map[int][][]int) [][]int {
//
//	if !numsMap[curr] && curr != 0 {
//		return [][]int{{}}
//	}
//
//	if v, ok := cache[curr]; ok {
//		return v
//	}
//	cache[curr] = make([][]int, 0)
//	comb = append(comb, curr)
//	for i := curr; i < curr+3; i++ {
//		c := make([]int, 0)
//		d := adapterAllPossibleWaysPerf(numsMap, i+1, c, end, cache)
//		cache[curr] = d
//	}
//	//comb = append(comb, cache[curr]...)
//	cache[curr] = append(cache[curr], comb)
//	return cache[curr]
//}

//var input = `16
//10
//15
//5
//1
//11
//7
//19
//6
//12
//4`

//var input = `28
//33
//18
//42
//31
//14
//46
//20
//48
//47
//24
//23
//49
//45
//19
//38
//39
//11
//1
//32
//25
//35
//8
//17
//7
//9
//4
//2
//34
//10
//3`

var input = `149
87
67
45
76
29
107
88
4
11
118
160
20
115
130
91
144
152
33
94
53
148
138
47
104
121
112
116
99
105
34
14
44
137
52
2
65
141
140
86
84
81
124
62
15
68
147
27
106
28
69
163
97
111
162
17
159
122
156
127
46
35
128
123
48
38
129
161
3
24
60
58
155
22
55
75
16
8
78
134
30
61
72
54
41
1
59
101
10
85
139
9
98
21
108
117
131
66
23
77
7
100
51`
