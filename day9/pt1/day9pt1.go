package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)
	disk_map := strings.Split(string(dat), "")

	// create sparse representation
	sparse_format_slice := make([]string, 0)
	for i, val := range disk_map {
		val_int, err := strconv.Atoi(val)
		check(err)
		var id string
		if i%2 == 0 {
			id = fmt.Sprint(i / 2)
		} else {
			id = "."
		}
		for j := 0; j < val_int; j++ {
			sparse_format_slice = append(sparse_format_slice, id)
		}
	}

	fmt.Println(sparse_format_slice)

	// creating compact representation
	for i := len(sparse_format_slice) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if sparse_format_slice[i] != "." && sparse_format_slice[j] == "." {
				sparse_format_slice[i], sparse_format_slice[j] = sparse_format_slice[j], sparse_format_slice[i]
			}
		}
	}

	fmt.Println(sparse_format_slice)

	// calculate checksum
	checksum := 0
	for pos, id := range sparse_format_slice {
		if id != "." {
			id_int, err := strconv.Atoi(id)
			check(err)
			checksum += pos * id_int
		}

	}

	fmt.Println(checksum)
}
