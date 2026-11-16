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

func BreakC(conditions ...bool) {
	for i, c := range conditions {
		_ = i
		if c {
			Break();
			break
		}
	}
}

func Assert(ok bool, v ...any) {
	if !ok {
		Panic(v...)
	}
}

func Unreachable(v ...any) {
	if len(v) == 0 {
		Panic("unreachble code executed")
	}
	Panic(v...)
}
