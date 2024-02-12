// Task 4
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fh, err := os.Open("./test.txt")
	if err != nil {
		fh = os.Stdin
	}
	reader := bufio.NewReaderSize(fh, 1024*1024)
	// writer := bufio.NewWriterSize(os.Stdout, 1024*1024*1)
	str, _ := reader.ReadString('\n')
	fields := strings.Fields(str)
	t, _ := strconv.Atoi(fields[0])

	for i := 0; i < t; i++ {
		str, _ = reader.ReadString('\n')
		fields = strings.Fields(str)
		n, _ := strconv.Atoi(fields[0])
		// m, _ := strconv.Atoi(fields[1])
		arr := make([]string, 0, n)
		for j := 0; j < n; j++ {
			str, _ = reader.ReadString('\n')
			fields = strings.Fields(str)
			// v, _ := strconv.Atoi(fields[0])
			arr = append(arr, fields[0])
		}
		sol := task(arr)
		// fmt.Println(sol)
		// writer.WriteByte('\n')
		for _, v := range sol {
			// if i > 0 {
			// fmt.Print(" ")
			// }
			fmt.Println(v)
		}
		// fmt.Println()
	}
	// writer.Flush()
}

func drawPath(buf [][]byte, x, y, dx, dy int, ch byte) (newX, newY int) {
	for {
		nX, nY := x+dx, y+dy
		if nX < 0 || nX >= len(buf) || nY < 0 || nY >= len(buf[nX]) || buf[nX][nY] != '.' {
			break
		}
		x, y = nX, nY
		buf[x][y] = ch
	}
	return x, y
}

func task(s []string) []string {
	n := len(s)
	m := len(s[0])
	buf := make([][]byte, n)
	for i, v := range s {
		buf[i] = []byte(v)
	}

	isA := true
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if buf[i][j] == 'A' || buf[i][j] == 'B' {
				if isA {
					isA = false
					x, y := i, j
					x, y = drawPath(buf, x, y, -1, 0, buf[i][j]+32)
					x, y = drawPath(buf, x, y, 0, -1, buf[i][j]+32)
					x, y = drawPath(buf, x, y, -1, 0, buf[i][j]+32)
				} else {
					x, y := i, j
					x, y = drawPath(buf, x, y, 1, 0, buf[i][j]+32)
					x, y = drawPath(buf, x, y, 0, 1, buf[i][j]+32)
					x, y = drawPath(buf, x, y, 1, 0, buf[i][j]+32)
				}
			}
		}
	}

	ret := make([]string, n)
	for i, v := range buf {
		ret[i] = string(v)
	}
	return ret
}
