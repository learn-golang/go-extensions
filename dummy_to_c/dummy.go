package main

import "C"

//export dummy
func dummy() *C.char {
	return C.CString("I am dummy function from Go.\n")
}

func main() {}
