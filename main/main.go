package main

import (
	"go-learn/intefaces"
	"log/slog"
	"math"
)

func main() {
	intefaces.Server()
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
