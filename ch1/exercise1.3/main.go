package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func measureTime(f func()) {
	start := time.Now()
	f()
	timeSeconds := time.Since(start).Nanoseconds()
	functionName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	fmt.Println(functionName, "time in seconds:", timeSeconds)
}

func main() {
	var functions = []func(){}
	functions = append(functions, echo1)
	functions = append(functions, echo2)
	functions = append(functions, echo3)
	for _, function := range functions {
		measureTime(function)
	}
}
