package intefaces

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func AssertTypes() {
	interToType()
}

func interToInterface() {
	var w io.Writer
	w = os.Stdout            //os.Stdout is *os.File
	rw := w.(io.ReadWriter)  //ok, because file satisfies interface
	bs := w.(io.ByteScanner) //panic
	fmt.Println(rw, bs)
}

func interToType() {
	var w io.Writer
	w = os.Stdout               //os.Stdout is *os.File
	fl := w.(*os.File)          //ok, because w is really a file
	bf, ok := w.(*bytes.Buffer) //panic: interface conversion: io.Writer is *os.File, not *bytes.Buffer
	fmt.Println(ok)
	fmt.Println(fl, bf)
}
