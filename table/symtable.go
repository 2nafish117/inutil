package table

// type Type int32
// type Value interface{}
// type Scope int32

// const (
// 	TYPE_S64 Type = 1 + iota
// 	Type_U64
// 	TYPE_S32
// 	TYPE_U32
// 	TYPE_S16
// 	TYPE_U16
// 	TYPE_S8
// 	TYPE_U8

// 	TYPE_F64
// 	TYPE_F32

// 	TYPE_STRING
// )

type SymbolTableEntry struct {
	Value interface{}
	Scope int32
}

type symbolTable map[string]SymbolTableEntry
