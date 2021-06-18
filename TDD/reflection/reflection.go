// reflection used to deal with unknown data types
// like struct of unknown structures passed as interface{}
package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}
