package main

import (
	"fmt"
	"log/slog"
	"math"
)

func main() {
	var sl []int
	sl = append(sl, 10)
	fmt.Println(sl[0])
}

func appendSome(arr []int) {
	arr[1] = 10
	arr[2] = 20
	arr[3] = 30
}

func changeValue(str *string) {
	*str = "bbb"
}

func isPowerOfFour(n int) bool {
	var kv map[string]int
	kv["sdf"] = 12
	var res = float64(n)
	for res > 3 && math.Floor(res) == res {
		res /= 4
	}
	slog.Default()
	return res == 1
}
