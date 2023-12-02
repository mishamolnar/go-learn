package functions

import (
	"fmt"
	"math/rand"
	"net/http"
)

func EntrypointFunc() {
}

func zero(x int, _ int) int {
	return x
}

func bareReturn() (x, y int) { //initializes x, y to default values in function invocation
	if rand.Int() < 100 {
		return //return 0, 0
	}
	x = rand.Int()
	y = rand.Int()
	return //returns x, y
}

func fetchUrl(url string) (int, error) {
	res, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("Fetch  failed for url %s, error %v", url, err)
	}
	return res.StatusCode, nil
}

func functionValues() {
	f := fetchUrl
	fmt.Printf("%T", f)
}
