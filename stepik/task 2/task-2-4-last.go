package main

import "fmt"

type myStruct struct {
	On    bool
	Ammo  int
	Power int
}

func (s *myStruct) Shoot() bool {
	if !s.On {
		return false
	}
	if s.Ammo <= 0 {
		return false
	}
	s.Ammo--
	return true
}

func (r *myStruct) RideBike() bool {
	if !r.On {
		return false
	}
	if r.Power <= 0 {
		return false
	}
	r.Power--
	return true
}

func main() {
	testStruct := &myStruct{true, 4, 3}
	fmt.Println(testStruct.Ammo)

}
