package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func Sleep(milliseconds int64) {
	var ts syscall.Timespec
	ts.Sec = milliseconds / 1000
	ts.Nsec = (milliseconds % 1000) * 1000000

	syscall.Syscall(syscall.SYS_NANOSLEEP, uintptr(unsafe.Pointer(&ts)), 0, 0)
}

func main() {
	fmt.Println("waiting...")
	Sleep(10000)
	fmt.Println("done")
}
