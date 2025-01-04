package main

import "fmt"

//	type SyncBuffer struct {
//		lock   sync.Locker
//		buffer bytes.Buffer
//	}
type vFile struct {
	fd      int
	name    string
	dirInfo string
	nepipe  int
}

func NewFile(fd int, name string) *vFile {
	if fd < 0 {
		return nil
	}
	return &vFile{fd, name, nil, 0}
}
func main() {
	t := NewFile(1, "Test")
	fmt.Println(t)
}
