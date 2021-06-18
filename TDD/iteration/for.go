package iteration

import (
	"errors"
	"fmt"
)

func Repeat(args ...interface{}) (ret string, err error) {
	str, iter, err := RepeatArgsProcessor(args)
	if err != nil {
		return
	}
	for i := 0; i < iter; i++ {
		ret += str
	}
	return ret, nil
}

func RepeatArgsProcessor(args []interface{}) (str string, iterations int, err error) {
	if len(args) > 2 || len(args) < 1 {
		err = errors.New("check number of arguments")
		return
	}
	for i, v := range args {
		switch i {
		case 0:
			str = fmt.Sprintf("%v", v)
		case 1:
			param := v.(int)
			iterations = param
		}
	}
	// fmt.Println(iterations)
	if iterations == 0 {
		iterations = 5
	}
	return str, iterations, nil
}
