package gomod

import (
	f "fmt"
	gdp "github.com/mishamolnar/golang-dependencies"
	golang_dependency_one_b "github.com/mishamolnar/golang-dependency-one-b"
)

func UseDependencies() {
	f.Println(gdp.G)
	f.Println(gdp.Level())
	f.Println(golang_dependency_one_b.Produce())
}
