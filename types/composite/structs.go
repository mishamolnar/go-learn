package composite

import (
	"fmt"
	"time"
)

type Wheel struct {
	Circle
	Strokes int
}

type Circle struct {
	Point
	Radius int
}

type Point struct {
	x, y int
}

func printWheel() {
	wheel := Wheel{Circle{Point{1, 2}, 1}, 2}
	wheel.x = 100
	wheel.y = 100
	wheel.Radius = 1
	fmt.Println(wheel)
}

func StructEntrypoint() {
	printWheel()
}

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	position  string
	Salary    int
	ManagerID int
}

var employees []Employee

func findByID(ID int) *Employee {
	for index := range employees {
		if employees[index].ID == ID {
			return &employees[index]
		}
	}
	return nil
}

func employeeCheck() {
	employees = append(employees, Employee{1, "dilbert", "test", time.Now(), "manager", 123123, 2})
	findByID(1).Salary = 1
	fmt.Printf("return type is %T \n", findByID(1))
	fmt.Println(*findByID(1))
	fmt.Println(employees)
}
