package types

import (
	"errors"
	"math"
)

type MyInt int32

var (
	ErrorNotAMyInt = errors.New("not a MyInt")
	ErrorDivideByZero = errors.New("not divide by zero")
)

func (i MyInt) Divide(n int) (MyInt, error){
	if MyInt(n) > 0 {
		cal := i / MyInt(n)
		if inBoundInt32(cal){
			return cal, nil
		}
		return 0, ErrorNotAMyInt
	} else {
		return 0, ErrorDivideByZero
	}
}

func (i MyInt) Add(n int) (MyInt, error){
	cal := i + MyInt(n)
	if inBoundInt32(cal){
		return 0, ErrorNotAMyInt
	}
	return cal, nil
}

func (i MyInt) Sub(n int) (MyInt, error){
	cal := i - MyInt(n)
	if inBoundInt32(cal){
		return 0, ErrorNotAMyInt
	}
	return cal, nil
}

func (i MyInt) Multiply(n int) (MyInt, error){
	cal := int64(i) * int64(n)
	if inBoundInt32(MyInt(cal)){
		return 0, ErrorNotAMyInt
	}
	return MyInt(cal), nil
}

func inBoundInt32(num MyInt) bool{
	if num <= math.MaxInt32 && num >= math.MinInt32 {
		return true
	}
	return false
}