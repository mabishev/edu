package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	type myStruct struct {
		Name string
		Age  int
	}
	var (
		src = myStruct{Name: "John", Age: 2}
		dst = myStruct{}
		buf = new(bytes.Buffer)
	)

	//buf2 := &bytes.Buffer{}                     //тоже самое что buf
	//var buf3 *bytes.Buffer = &bytes.Buffer{}   // тоже самое что buf

	enc := json.NewEncoder(buf)
	dec := json.NewDecoder(buf)

	enc.Encode(dst)
	dec.Decode(&src)
	fmt.Println(src)

}
