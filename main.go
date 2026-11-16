package dev_utils

import (
	"fmt"
	"runtime"
)

func Panic(v ...any) {
	panic(fmt.Sprint(v...))
}

func Break() {
	runtime.Breakpoint()
}
