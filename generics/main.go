package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(sumInts(map[string]int64{"a": 1, "b": 2}))
	fmt.Println(sumInts(map[string]float64{"a": 1.2, "b": 2.2}))
	typeToMap(make([]*Vacancy, 0))
}

func sumInts[K comparable, V int64 | float64](m map[K]V) V {
	var sum V
	for _, i := range m {
		sum += i
	}
	return sum
}

type fileMetadata struct {
	id         int64
	name, path string
	createdAt  time.Time
	deleted    bool
}

type user struct {
	id   int64
	name string
	age  uint32
}

func (f *fileMetadata) getName() string {
	return f.name
}

func (u *user) getName() string {
	return u.name
}

func printName(n namer) {
	fmt.Println(n.getName())
}

type namer interface {
	getName() string
}

type Company struct {
	ID          int
	Name        string
	Description string
}

func (c *Company) getId() int {
	return c.ID
}

type Vacancy struct {
	ID          int
	CompanyID   int
	Title       string
	Description string
}

func (v *Vacancy) getId() int {
	return v.ID
}

type Getter interface {
	getId() int
}

func typeToMap[T Getter](t []T) map[int]T {
	m := make(map[int]T)
	for _, data := range t {
		m[data.getId()] = data
	}
	return m
}
