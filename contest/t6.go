// Task 6
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
		// p, _ := strconv.Atoi(fields[1])
		arr := make([]string, 0, n)
		for j := 0; j < n; j++ {
			str, _ = reader.ReadString('\n')
			fields = strings.Fields(str)
			// v, _ := strconv.Atoi(fields[0])
			arr = append(arr, fields[0])
		}
		sol := task(arr)
		fmt.Println(sol)
		// writer.WriteByte('\n')
		// for _, v := range sol {
		// if i > 0 {
		// fmt.Print(" ")
		// }
		// fmt.Println(v)
		// }
		// fmt.Println()
	}
	// writer.Flush()
}

func worstMark(arr []string, exclRow, exclCol int) (m byte, r, c int) {
	var worst byte = '6'
	wR, wC := -1, -1
	for i, s := range arr {
		if i == exclRow {
			continue
		}
		for j := 0; j < len(s); j++ {
			if j == exclCol {
				continue
			}
			if s[j] < worst {
				worst = s[j]
				wR = i
				wC = j
			}
		}
	}
	return worst, wR, wC
}

func task(arr []string) string {
	_, r, c := worstMark(arr, -1, -1)

	// Row
	_, _, c2 := worstMark(arr, r, -1)
	m2, _, _ := worstMark(arr, r, c2)
	// Col
	_, r3, _ := worstMark(arr, -1, c)
	m3, _, _ := worstMark(arr, r3, c)

	if m2 > m3 {
		return fmt.Sprint(r+1, c2+1)
	} else {
		return fmt.Sprint(r3+1, c+1)
	}
}
