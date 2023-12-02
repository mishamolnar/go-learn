package primitive

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("Primitive intialized")
}

type ByteSize int

var f *os.File

func EntryPoint() {
	openFile()
}

func openFile() {
	f, err := os.OpenFile("/Users/mmoln/GolandProjects/go-learn/primitive/test.txt", os.O_WRONLY, os.ModeAppend) //creates new f, outer f - null
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString("ABCD")
	if err != nil {
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}
	checkFile()
}

func checkFile() {
	fmt.Print("_________")
	fmt.Print(f.Name()) //throws exception because we have := (declaration) and not = (assignment) in line 21
}

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func measurementsInBytes() {
	fmt.Printf("KB in decimal is %v: and in binary is: %b \n", KB, KB)
	fmt.Printf("MB in decimal is %v: and in binary is: %b \n", MB, MB)
	fmt.Printf("GB in decimal is %v: and in binary is: %b \n", GB, GB)
	fmt.Printf("TB in decimal is %v: and in binary is: %b \n", TB, TB)
	fmt.Printf("PB in decimal is %v: and in binary is: %b \n", PB, PB)
	fmt.Printf("EB in decimal is %v: and in binary is: %b \n", EB, EB)
}

func floatInt() {
	var a, b = 190302002222211221, 2.0
	fmt.Printf("Types are: %T, %T", a, b)
	//fmt.Print(a * b) mismatched types, int and float
}

func maxNumbers() {
	var max8 int8 = 1<<7 - 1
	var maxu8 uint8 = 1<<8 - 1
	fmt.Printf("Signed in binary: %b, in decimal %d of type %T \n", max8, max8, max8)
	fmt.Printf("Signed in binary: %b, in decimal %d of type %T \n", maxu8, maxu8, maxu8)
}
