// Task 3
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
		// n, _ := strconv.Atoi(fields[0])
		// p, _ := strconv.Atoi(fields[1])
		// arr := make([]int, 0, n)
		// for j := 0; j < n; j++ {
		// str, _ = reader.ReadString('\n')
		// fields = strings.Fields(str)
		// v, _ := strconv.Atoi(fields[0])
		// arr = append(arr, v)
		// }
		sol := task(fields[0])
		fmt.Println(sol)
		// writer.WriteByte('\n')
		// for i, v := range sol {
		// if i > 0 {
		// fmt.Print(" ")
		// }
		// fmt.Print(v)
		// }
		// fmt.Println()
	}
	// writer.Flush()
}

func task(s string) string {
	state := 0
	fsm := []map[byte]int{
		// 0 - wait
		{
			'M': 2,
			'R': 1,
			'C': 1,
			'D': 1,
		},
		// 1 - error
		{
			'M': 1,
			'R': 1,
			'C': 1,
			'D': 1,
		},
		// 2 - run
		{
			'M': 1,
			'R': 3,
			'C': 4,
			'D': 0,
		},
		// 3 - reloaded
		{
			'M': 1,
			'R': 1,
			'C': 4,
			'D': 1,
		},
		// 4 - cancelled
		{
			'M': 2,
			'R': 1,
			'C': 1,
			'D': 1,
		},
	}
	for i := 0; i < len(s); i++ {
		state = fsm[state][s[i]]
	}
	if state != 0 {
		return "NO"
	}
	return "YES"
}
