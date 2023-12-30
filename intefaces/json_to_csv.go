package intefaces

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"regexp"
	"strings"
)

func ReadAndMap() {
	file, _ := os.Open("/Users/mmoln/Downloads/Musical_Instruments_5.json")
	out, _ := os.Create("/Users/mmoln/Downloads/reviews_with_score.csv")
	dec := json.NewDecoder(file)
	enc := json.NewEncoder(out)
	chars := []string{"]", "^", "\\\\", "[", "\"", "(", ")", "-", ":", "\\", "\n"}
	regexStr := strings.Join(chars, "")
	re := regexp.MustCompile("[" + regexStr + "]+")

	for dec.More() {
		var r Review

		err := dec.Decode(&r)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(r)
		r.Summary = re.ReplaceAllString(r.Summary, "")
		r.ReviewText = re.ReplaceAllString(r.ReviewText, "")
		r.Asin = uuid.NewString()

		err = enc.Encode(r)
		if err != nil {
			log.Fatal(err)
		}
	}
}

type Review struct {
	Asin       string  `json:"asin"`
	Summary    string  `json:"summary"`
	ReviewText string  `json:"reviewText"`
	Overall    float32 `json:"overall"`
}
