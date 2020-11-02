package util

import (
	"fmt"
	"reflect"

	"github.com/swnb/demo-package/internal"
)

func Compare(a ...interface{}) string {
	for i, v := range a {
		if v == nil {
			fmt.Println("nil")
		} else {
			t := reflect.TypeOf(v)
			if reflect.Slice == t.Kind() {
				fmt.Println(i+1, t.Name())
			} else {
				fmt.Println(i+1, t.Kind().String(), t.Name())
			}
		}
	}
	return internal.Foo()
}
