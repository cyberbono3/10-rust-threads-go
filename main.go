package main

/*
#cgo LDFLAGS: -L./lib -lembed
#include "./lib/process.h"
*/
import "C"

func main() {
	C.process()
}