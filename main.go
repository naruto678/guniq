package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {

	var reader io.Reader

	if len(os.Args) == 2 {
		fileName := os.Args[1]
		reader, _ = os.Open(fileName)

	} else {
		reader = os.Stdin
	}

	scanner := bufio.NewScanner(reader)
	lineMap := map[string]bool{}
	idx := map[string]int{}
	count := 0
	for scanner.Scan() {
		val := scanner.Text()
		if lineMap[val] {
			continue
		} else {
			lineMap[val] = true
			idx[val] = count
			count++
		}
	}
	result := []string{}
	for key := range lineMap {
		result = append(result, key)
	}
	sort.Slice(result, func(i, j int) bool {
		return idx[result[i]] < idx[result[j]]
	})
	for _, val := range result {
		fmt.Println(val)
	}
}
