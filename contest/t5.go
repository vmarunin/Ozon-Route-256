package main

import (
	"bufio"
	"encoding/json"
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
			// fields = strings.Fields(str)
			// v, _ := strconv.Atoi(fields[0])
			arr = append(arr, str)
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

type TaskFS struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files,omitempty"`
	Folders []TaskFS `json:"folders,omitempty"`
}

func task(arr []string) int {
	data := strings.Join(arr, "")
	var root TaskFS
	json.Unmarshal([]byte(data), &root)
	// fmt.Printf("%#v", root)
	var rf func(node TaskFS) (total, forRM int)
	rf = func(node TaskFS) (total, forRM int) {
		total = len(node.Files)
		for _, f := range node.Folders {
			t, r := rf(f)
			total += t
			forRM += r
		}
		isVirus := false
		for _, f := range node.Files {
			if strings.HasSuffix(f, ".hack") {
				isVirus = true
			}
		}
		if isVirus {
			forRM = total
		}
		return total, forRM
	}
	_, ret := rf(root)
	return ret
}
