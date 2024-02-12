// Task 2
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
		p, _ := strconv.Atoi(fields[1])
		arr := make([]int, 0, n)
		for j := 0; j < n; j++ {
			str, _ = reader.ReadString('\n')
			fields = strings.Fields(str)
			v, _ := strconv.Atoi(fields[0])
			arr = append(arr, v)
		}
		sol := task(arr, p)
		fmt.Printf("%.2f\n", float64(sol)/100)
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

func task(arr []int, p int) int {
	ret := 0
	for _, v := range arr {
		ret += (v * p) % 100
	}
	return ret
}
