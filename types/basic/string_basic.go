package basic

import (
	"bytes"
	"fmt"
)

func EntryPoint() {
	rawString()
}

func rawString() {
	fmt.Println(`asdf
asdf
				\n dfasdf`)
}

func intsToString(values []int) string {
	var buff bytes.Buffer
	buff.WriteRune('[')
	for i, v := range values {
		if i > 0 {
			buff.WriteString(", ")
		}
		fmt.Fprintf(&buff, "%d", v)
	}
	buff.WriteRune('[')
	return buff.String()
}
