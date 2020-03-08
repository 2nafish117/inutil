package runtime

import (
	"fmt"
	"inutil/runtime/buffer"
	"inutil/runtime/table"
)

// Env holds all runtime information of the VM
type Env struct {
	SymTable     table.SymbolTable
	Stack        *buffer.ExecStack
	Instructions buffer.InstructionBuffer
}

// E is the global VM environment
var E Env

func init() {
	E.Instructions = buffer.InstructionBuffer{
		{Op: buffer.PushName, Args: "var1"},
		{Op: buffer.PushName, Args: "var2"},
		{Op: buffer.PushConst, Args: float64(5.78)},
		{Op: buffer.PushConst, Args: int64(-6)},
		{Op: buffer.ArithAdd, Args: nil},
		{Op: buffer.ArithAdd, Args: nil},
		{Op: buffer.ArithAdd, Args: nil},
		{Op: buffer.PopName, Args: "var3"},
	}

	E.Stack = buffer.NewExecStack()

	E.SymTable = map[string]table.SymbolTableEntry{
		"var1": {Value: int64(5), Scope: 0},
		"var2": {Value: int64(2), Scope: 0},
	}

}

// Execute allows Env to execute instruction buffer
func Execute(e *Env) {
	e.Instructions.Execute(e.SymTable, e.Stack)
}

// Dump dumps the golbal env in readable form
func (e *Env) Dump() {
	fmt.Print("=====ENV DUMP=====\n")
	e.Instructions.Dump()
	e.Stack.Dump()
	e.SymTable.Dump()
}
