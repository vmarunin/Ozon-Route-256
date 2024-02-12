// Task 7
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
	writer := bufio.NewWriterSize(os.Stdout, 1024*1024*1)
	str, _ := reader.ReadString('\n')
	fields := strings.Fields(str)
	t, _ := strconv.Atoi(fields[0])

	m := map[string]bool{}

	for i := 0; i < t; i++ {
		str, _ = reader.ReadString('\n')
		fields = strings.Fields(str)
		m[fields[0]] = true
		buf := []byte(fields[0])
		for j := 1; j < len(buf); j++ {
			buf[j-1], buf[j] = buf[j], buf[j-1]
			m[string(buf)] = true
			buf[j-1], buf[j] = buf[j], buf[j-1]
		}
	}
	str, _ = reader.ReadString('\n')
	fields = strings.Fields(str)
	n, _ := strconv.Atoi(fields[0])
	for i := 0; i < n; i++ {
		str, _ = reader.ReadString('\n')
		fields = strings.Fields(str)
		if m[fields[0]] {
			fmt.Fprintln(writer, 1)
		} else {
			fmt.Fprintln(writer, 0)

		}
	}
	writer.Flush()
}
