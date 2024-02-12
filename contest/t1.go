// Task 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	a, _ := strconv.Atoi(fields[0])
	b, _ := strconv.Atoi(fields[1])

	fmt.Println(a - b)
	// for i := 0; i < t; i++ {
	// 	str, _ = reader.ReadString('\n')
	// 	fields := strings.Fields(str)
	// 	n, _ := strconv.Atoi(fields[0])
	// 	arr := make([]float64, 0, n)
	// 	for j := 1; j <= n; j++ {
	// 		v, _ := strconv.ParseFloat(fields[j], 64)
	// 		arr = append(arr, v)
	// 	}
	// 	sol := task(arr)
	// 	fmt.Println(sol)
	// 	// writer.WriteByte('\n')
	// 	// for i, v := range sol {
	// 	// if i > 0 {
	// 	// fmt.Print(" ")
	// 	// }
	// 	// fmt.Print(v)
	// 	// }
	// 	// fmt.Println()
	// }
	// writer.Flush()
}

func task(arr []float64) int {
	result := make([]float64, len(arr))
	copy(result, arr)
	sort.Float64s(result)
	stack := []float64{}
	sPos := 0
	for _, v := range result {
		if len(stack) > 0 && stack[len(stack)-1] == v {
			stack = stack[:len(stack)-1]
			continue
		}
		for sPos < len(arr) && v != arr[sPos] {
			stack = append(stack, arr[sPos])
			sPos++
		}
		if sPos < len(arr) && v == arr[sPos] {
			sPos++
			continue
		}
		return 0
	}
	return 1
}
