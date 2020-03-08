package table

import (
	"fmt"
)

// SymbolTableEntry is a single value scope pair
type SymbolTableEntry struct {
	Value interface{}
	Scope int32
}

// SymbolTable is the name value pair of all variables in the env
type SymbolTable map[string]SymbolTableEntry

// Dump does memory dump of whats in the table
func (st *SymbolTable) Dump() {
	fmt.Println("=====SYMBOL TABLE=====")
	for key, val := range *st {
		fmt.Printf("%5s%5v\n", key, val)
	}
	fmt.Println()
}
