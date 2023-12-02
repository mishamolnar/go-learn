package bufio

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// scans from standart input or file, outputs number of occurances of each line
func Unique() {
	counts := make(map[string]int)
	files := os.Args[1:]
	for _, filename := range files {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
		} else {
			scan(f, counts)
		}
	}
	if len(os.Args) == 1 {
		scan(os.Stdin, counts)
	}
	for key, val := range counts {
		fmt.Printf("%v occurs %v times \n", strings.Trim(key, " "), val)
	}
}

func scan(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			fmt.Printf("\nDuplicate line occured in the file %v \n", f.Name())
		}
	}
}
