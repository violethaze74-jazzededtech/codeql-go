// Code generated by https://github.com/gagliardetto/codebox. DO NOT EDIT.

package main

import "errors"

func TaintStepTest_ErrorsAs_B0I0O0(sourceCQL interface{}) interface{} {
	fromError656 := sourceCQL.(error)
	var intoInterface414 interface{}
	errors.As(fromError656, intoInterface414)
	return intoInterface414
}

func TaintStepTest_ErrorsNew_B0I0O0(sourceCQL interface{}) interface{} {
	fromString518 := sourceCQL.(string)
	intoError650 := errors.New(fromString518)
	return intoError650
}

func TaintStepTest_ErrorsUnwrap_B0I0O0(sourceCQL interface{}) interface{} {
	fromError784 := sourceCQL.(error)
	intoError957 := errors.Unwrap(fromError784)
	return intoError957
}

func RunAllTaints_Errors() {
	{
		source := newSource(0)
		out := TaintStepTest_ErrorsAs_B0I0O0(source)
		sink(0, out)
	}
	{
		source := newSource(1)
		out := TaintStepTest_ErrorsNew_B0I0O0(source)
		sink(1, out)
	}
	{
		source := newSource(2)
		out := TaintStepTest_ErrorsUnwrap_B0I0O0(source)
		sink(2, out)
	}
}
