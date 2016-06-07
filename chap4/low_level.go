package main

// #include <stdio.h>
// #include <string.h>
//  int string_length (char* str) {
//    return strlen(str);
//  }
import "C"
import "fmt"

func main() {
	v := C.CString("Don't Forget My Memory Is Not Visible To Go!")
	x := C.string_length(v)
	fmt.Println("A C function has determined your string is", x, "characters in length")
}
