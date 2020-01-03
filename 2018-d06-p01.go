package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = `78, 335
74, 309
277, 44
178, 286
239, 252
118, 354
170, 152
75, 317
156, 318
172, 45
138, 162
261, 195
306, 102
282, 67
53, 141
191, 237
352, 180
95, 247
353, 357
201, 327
316, 336
57, 43
119, 288
299, 328
125, 327
187, 186
121, 151
121, 201
43, 67
76, 166
238, 148
326, 221
219, 207
237, 160
345, 244
321, 346
48, 114
304, 80
265, 216
191, 92
54, 75
118, 260
336, 249
81, 103
290, 215
300, 246
293, 59
150, 274
296, 311
264, 286`

type Point struct {
	X int
	Y int
}

func main() {
	lines := strings.Split(input, "\n")
	points := make([]Point,len(lines))

	for i, o := range lines {
		ps := strings.Split(o,",")
		//pointX := strconv.ParseInt(splitPoint[0], 10, 64)
		if ss, err := strconv.ParseInt(ps[0], 10, 64); err != nil {
			fmt.Println("Not Cool!")
		} else {
			fmt.Println("Cool!", " ", ss)
		}
		points[i] = Point{}
	}
}
