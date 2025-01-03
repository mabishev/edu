package edu

import (
	"bytes"
	"fmt"
)

func BufferBytes() {
	var buffer bytes.Buffer
	for i := 0; i < 10; i++ {
		buffer.WriteString(fmt.Sprint(i))
	}
	buffer.WriteString(fmt.Sprint("салем"))

	//result := buffer.String()
	fmt.Println(buffer.String())
}
