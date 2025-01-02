package edu

import (
	"fmt"
)

type Flags uint

const (
	FlagUp Flags = 1 << iota // последовательность возрастающих степеней двойки(побитовый сдвиг)
	FlagBroadcast
	Flagloopback
	FlagPointToPoint
	FlagMulticast
	Flag32
	Flag64
)
const (
	_   = 1 << (10 * iota) // если дать имя константы, то результат будет равен 1
	KiB                    // 1024
	MiB                    // 1048576
	GiB                    // 1073741824
	TiB                    // 1099511627776 (превышает 1 << 32)
	PiB                    // 1125899906842624
	EiB                    // 1152921504606846976
	ZiB                    // 1180591620717411303424 (превышает 1 << 64) слишком велики, чтобы хранить их в любой переменной целочисленного типа, но они являются корректными константами
	YiB                    // 1208925819614629174706176 слишком велики, чтобы хранить их в любой переменной целочисленного типа, но они являются корректными константами
)

// побитовый сдвиг
func BitShift() {
	fmt.Println(FlagUp, FlagBroadcast, Flagloopback, FlagPointToPoint, FlagMulticast, Flag32, Flag64)
	fmt.Println(KiB, MiB, GiB, TiB)

	arr := make([]int, 10)
	for i := range arr {
		arr[i] = 512 >> i
	}
	fmt.Println(arr) //[1 2 4 8 16 32 64 128 256 512]
}
