package intefaces

import (
	"fmt"
	"os/user"
)

type CountWriter user.User

func (c *CountWriter) Write(p []byte) (n int, err error) {
	c.Name += string(p)
	return len(p), nil
}

func WriterUsage() {
	var wr CountWriter
	fmt.Fprint(&wr, "Something")
	fmt.Fprint(&wr, " Something")
	fmt.Println(wr.Name)
}
