package main

import (
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	funcs := make([]interface{}, 3, 3)               // I use interface{} to allow any kind of func
	funcs[0] = func(a int) int { return a + 1 }      // good
	funcs[1] = func(a string) int { return len(a) }  // good
	funcs[2] = func(a string) string { return ":(" } // bad
	for _, fi := range funcs {
		f := reflect.ValueOf(fi)
		functype := f.Type()
		good := false
		spew.Dump(functype.NumIn())
		for i := 0; i < functype.NumIn(); i++ {
			if "int" == functype.In(i).String() {
				good = true // yes, there is an int among inputs
				break
			}
		}
		for i := 0; i < functype.NumOut(); i++ {
			if "int" == functype.Out(i).String() {
				good = true // yes, there is an int among outputs
				break
			}
		}
		if good {
			fmt.Println(f)
		}
	}
}
