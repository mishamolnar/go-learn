package scope

import (
	"log"
	"os"
)

var cwd string

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd() failed: %v", err)
	}
	log.Printf("Woring dir is %v", cwd) // Woring dir is /Users/mmoln/GolandProjects/go-learn/main
}

func EntryPoint() {
	log.Println(cwd) //nothing
}
