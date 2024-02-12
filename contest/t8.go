// Task 8
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
		arr := make([]int, 0, n)
		str, _ = reader.ReadString('\n')
		fields = strings.Fields(str)
		for j := 0; j < n; j++ {
			v, _ := strconv.Atoi(fields[j])
			arr = append(arr, v)
		}
		sol := task(arr)
		// fmt.Println(sol)
		// writer.WriteByte('\n')
		for j, v := range sol {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
		fmt.Println()
	}
	// writer.Flush()
}

func task(arr []int) []int {
	ret := make([]int, len(arr))
	prevK := 0
	prevCnt := 0
	prevPos := -4*len(arr) - 1
	for i := 1; i < len(arr)-1; i++ {
		k := 0
		for {
			k++
			if i-k < 0 || i+k >= len(arr) || arr[i-k] >= arr[i-k+1] || arr[i+k] >= arr[i+k-1] {
				break
			}
			if ret[k-1] == 0 {
				ret[k-1] = 1
			}
		}
		k--
		if k == 0 {
			continue
		}
		if k == prevK && i-prevPos == 2*k {
			prevCnt++
		} else {
			prevCnt = 1
		}
		prevK = k
		prevPos = i
		if ret[k-1] < prevCnt {
			ret[k-1] = prevCnt
		}
	}
	return ret
}
