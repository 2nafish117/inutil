package buffer

import (
	"fmt"
	"inutil/runtime/table"
	"log"
	"reflect"
)

// Enum of all instruction operations
const (
	Nop int8 = iota

	PushName // pushes value of a name onto stack
	PopName  // pops value of a name from stack into symtable

	PushConst // pushes a constant onto stack
	PopConst  // pops a constant from stack (kinda useless i think)

	ArithAdd // adds the top 2 values on the stack, and places result on stack (consumes the inputs)
	ArithSub // subs the top 2 values on the stack, and places result on stack (consumes the inputs)
	ArithMul // muls the top 2 values on the stack, and places result on stack (consumes the inputs)
	ArithDiv // divs the top 2 values on the stack, and places result on stack (consumes the inputs)

	BoolOr  // boolean or of 2 values on stack, places result on stack (consumes input)
	BoolAnd // boolean and of 2 values on stack, places result on stack (consumes input)
	BoolNot // boolean not of the value on stack, places result on stack (consumes input)

	Goto    // unconditional jump
	GotoGZ  // jump to instr given by Arg if stack top is greater than zero
	GotoLZ  // jump to instr given by Arg if stack top is lesser than zero
	GotoGEZ // jump to instr given by Arg if stack top is greater or equal to zero
	GotoLEZ // jump to instr given by Arg if stack top is lesser or equal to zero
	GotoEZ  // jump to instr given by Arg if stack top is equal to zero
	GotoNEZ // jump to instr given by Arg if stack top is not equal to zero

	CallFunc // TODO docs
)

// OpStrings converts enum to string
var OpStrings = []string{
	"Nop",

	"PushName",
	"PopName",

	"PushConst",
	"PopConst",

	"ArithAdd",
	"ArithSub",
	"ArithMul",
	"ArithDiv",

	"BoolOr",
	"BoolAnd",
}

// Instruction type
type Instruction struct {
	Op   int8
	Args interface{}
}

// @TODO: maybe have a function that returns the max of 2 types, for type conversions

// Refactor all these functions ??? any way to make it so each new type doesnt explode
// maybe restrict the automatic type conversions type, so only int64 + int64 works, and so on
// if float64 + int64 is needed, user has to explicitly typcast ??
func arithAdd(val1 interface{}, val2 interface{}) (res interface{}) {
	switch v1 := reflect.ValueOf(val1); v1.Kind() {
	case reflect.Int64:
		switch v2 := reflect.ValueOf(val2); v2.Kind() {
		case reflect.Int64:
			res = val1.(int64) + val2.(int64)
		case reflect.Float64:
			res = float64(val1.(int64)) + val2.(float64)
		default:
			log.Printf("%v arith not implemented\n", v1.Kind())
		}
	case reflect.Float64:
		switch v2 := reflect.ValueOf(val2); v2.Kind() {
		case reflect.Int64:
			res = val1.(float64) + float64(val2.(int64))
		case reflect.Float64:
			res = val1.(float64) + val2.(float64)
		default:
			log.Printf("%v arith not implemented\n", v1.Kind())
		}
	default:
		log.Printf("%v arith not implemented\n", v1.Kind())
	}

	return
}

func arithSub(val1 interface{}, val2 interface{}) (res interface{}) {
	switch v1 := reflect.ValueOf(val1); v1.Kind() {
	case reflect.Int64:
		switch v2 := reflect.ValueOf(val2); v2.Kind() {
		case reflect.Int64:
			res = val1.(int64) - val2.(int64)
		case reflect.Float64:
			res = float64(val1.(int64)) - val2.(float64)
		default:
			log.Printf("%v arith not implemented\n", v1.Kind())
		}
	case reflect.Float64:
		switch v2 := reflect.ValueOf(val2); v2.Kind() {
		case reflect.Int64:
			res = val1.(float64) - float64(val2.(int64))
		case reflect.Float64:
			res = val1.(float64) - val2.(float64)
		default:
			log.Printf("%v arith not implemented\n", v1.Kind())
		}
	default:
		log.Printf("%v arith not implemented\n", v1.Kind())
	}
	return
}

func arithMul(val1 interface{}, val2 interface{}) (res interface{}) {
	switch v1 := reflect.ValueOf(val1); v1.Kind() {
	case reflect.Int64:
		switch v2 := reflect.ValueOf(val2); v2.Kind() {
		case reflect.Int64:
			res = val1.(int64) * val2.(int64)
		case reflect.Float64:
			res = float64(val1.(int64)) * val2.(float64)
		default:
			log.Printf("%v arith not implemented\n", v1.Kind())
		}
	case reflect.Float64:
		switch v2 := reflect.ValueOf(val2); v2.Kind() {
		case reflect.Int64:
			res = val1.(float64) * float64(val2.(int64))
		case reflect.Float64:
			res = val1.(float64) * val2.(float64)
		default:
			log.Printf("%v arith not implemented\n", v1.Kind())
		}
	default:
		log.Printf("%v arith not implemented\n", v1.Kind())
	}

	return
}

func arithDiv(val1 interface{}, val2 interface{}) (res interface{}) {
	switch v1 := reflect.ValueOf(val1); v1.Kind() {
	case reflect.Int64:
		switch v2 := reflect.ValueOf(val2); v2.Kind() {
		case reflect.Int64:
			res = val1.(int64) / val2.(int64)
		case reflect.Float64:
			res = float64(val1.(int64)) / val2.(float64)
		default:
			log.Printf("%v arith not implemented\n", v1.Kind())
		}
	case reflect.Float64:
		switch v2 := reflect.ValueOf(val2); v2.Kind() {
		case reflect.Int64:
			res = val1.(float64) / float64(val2.(int64))
		case reflect.Float64:
			res = val1.(float64) / val2.(float64)
		default:
			log.Printf("%v arith not implemented\n", v1.Kind())
		}
	default:
		log.Printf("%v arith not implemented\n", v1.Kind())
	}

	return
}

// Execute executes an instruction
func (i *Instruction) Execute(symTable table.SymbolTable, stack *ExecStack) {
	switch i.Op {
	case Nop:
		{
			// Do nothing
		}

	case PushName:
		{
			name := i.Args.(string)
			value := symTable[name].Value
			stack.Push(value)
		}
	case PopName:
		{
			name := i.Args.(string)
			value, err := stack.Pop()
			scope := symTable[name].Scope
			if err != nil {
				log.Fatal(err.Error())
			}
			symTable[name] = table.SymbolTableEntry{Value: value, Scope: scope}
		}

	case PushConst:
		{
			value := i.Args
			stack.Push(value)
		}
	case PopConst:
		{
			_, err := stack.Pop()
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	case ArithAdd:
		{
			val1, err1 := stack.Pop()
			if err1 != nil {
				log.Fatal(err1.Error())
			}
			val2, err2 := stack.Pop()
			if err2 != nil {
				log.Fatal(err2.Error())
			}

			res := arithAdd(val1, val2)
			stack.Push(res)
		}
	case ArithSub:
		{
			val1, err1 := stack.Pop()
			if err1 != nil {
				log.Fatal(err1.Error())
			}
			val2, err2 := stack.Pop()
			if err2 != nil {
				log.Fatal(err2.Error())
			}

			res := arithSub(val1, val2)
			stack.Push(res)
		}
	case ArithMul:
		{
			val1, err1 := stack.Pop()
			if err1 != nil {
				log.Fatal(err1.Error())
			}
			val2, err2 := stack.Pop()
			if err2 != nil {
				log.Fatal(err2.Error())
			}

			res := arithMul(val1, val2)
			stack.Push(res)
		}
	case ArithDiv:
		{
			val1, err1 := stack.Pop()
			if err1 != nil {
				log.Fatal(err1.Error())
			}
			val2, err2 := stack.Pop()
			if err2 != nil {
				log.Fatal(err2.Error())
			}

			res := arithDiv(val1, val2)
			stack.Push(res)
		}
	case BoolOr:
		{
		}
	case BoolAnd:
		{
		}
	case CallFunc:
		{
		}
	}
}

// InstructionBuffer type for buffer of instructions
type InstructionBuffer []Instruction

// Dump dumps readable instruction buffer
func (ib *InstructionBuffer) Dump() {
	fmt.Printf("=====INSTRUCTIONS=====\n")
	for i, instr := range *ib {
		if instr.Args != nil {
			fmt.Printf("%5d%15s%30v\n", i, OpStrings[instr.Op], instr.Args)
		} else {
			fmt.Printf("%5d%15s%30s\n", i, OpStrings[instr.Op], "<nil>")
		}
	}
	fmt.Println()
}

// Execute executes an instruction buffer
func (ib *InstructionBuffer) Execute(symTable table.SymbolTable, stack *ExecStack) {
	for _, instr := range *ib {
		instr.Execute(symTable, stack)
	}
}
