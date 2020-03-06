// Code generated by gocc; DO NOT EDIT.

package parser

import(
        _ "inutil/token"
        _ "inutil/util"
        _ "inutil/table"
    )

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Program	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Program : Functions	<<  >>`,
		Id:         "Program",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Functions : Functions Function	<<  >>`,
		Id:         "Functions",
		NTType:     2,
		Index:      2,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Functions : "empty"	<<  >>`,
		Id:         "Functions",
		NTType:     2,
		Index:      3,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `Function : func identifier "(" Args ")" StatementBlock	<<  >>`,
		Id:         "Function",
		NTType:     3,
		Index:      4,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `StatementBlock : "{" Statements "}"	<<  >>`,
		Id:         "StatementBlock",
		NTType:     4,
		Index:      5,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Args : Args "," identifier	<<  >>`,
		Id:         "Args",
		NTType:     5,
		Index:      6,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Args : identifier	<<  >>`,
		Id:         "Args",
		NTType:     5,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Args : "empty"	<<  >>`,
		Id:         "Args",
		NTType:     5,
		Index:      8,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `Statements : Statements Statement	<<  >>`,
		Id:         "Statements",
		NTType:     6,
		Index:      9,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statements : StatementBlock	<<  >>`,
		Id:         "Statements",
		NTType:     6,
		Index:      10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statements : "empty"	<<  >>`,
		Id:         "Statements",
		NTType:     6,
		Index:      11,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `Statement : ControlStatements	<<  >>`,
		Id:         "Statement",
		NTType:     7,
		Index:      12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : AssignStatement ";"	<<  >>`,
		Id:         "Statement",
		NTType:     7,
		Index:      13,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : FunctionCall ";"	<<  >>`,
		Id:         "Statement",
		NTType:     7,
		Index:      14,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AssignStatement : Lhs "=" Rhs	<<  >>`,
		Id:         "AssignStatement",
		NTType:     8,
		Index:      15,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Lhs : identifier	<<  >>`,
		Id:         "Lhs",
		NTType:     9,
		Index:      16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rhs : Expression	<<  >>`,
		Id:         "Rhs",
		NTType:     10,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ControlStatements : IfStatement ElseStatement	<<  >>`,
		Id:         "ControlStatements",
		NTType:     11,
		Index:      18,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ControlStatements : ForStatement	<<  >>`,
		Id:         "ControlStatements",
		NTType:     11,
		Index:      19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `IfStatement : if Expression StatementBlock	<<  >>`,
		Id:         "IfStatement",
		NTType:     12,
		Index:      20,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ElseStatement : else StatementBlock	<<  >>`,
		Id:         "ElseStatement",
		NTType:     13,
		Index:      21,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ElseStatement : "empty"	<<  >>`,
		Id:         "ElseStatement",
		NTType:     13,
		Index:      22,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `ForStatement : for ForExpression ";" ForExpression ";" ForExpression StatementBlock	<<  >>`,
		Id:         "ForStatement",
		NTType:     14,
		Index:      23,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ForExpression : Rhs	<<  >>`,
		Id:         "ForExpression",
		NTType:     15,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ForExpression : AssignStatement	<<  >>`,
		Id:         "ForExpression",
		NTType:     15,
		Index:      25,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ForExpression : "empty"	<<  >>`,
		Id:         "ForExpression",
		NTType:     15,
		Index:      26,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `FunctionCall : identifier "(" Params ")"	<<  >>`,
		Id:         "FunctionCall",
		NTType:     16,
		Index:      27,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Params : Params "," Param	<<  >>`,
		Id:         "Params",
		NTType:     17,
		Index:      28,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Params : Param	<<  >>`,
		Id:         "Params",
		NTType:     17,
		Index:      29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Params : "empty"	<<  >>`,
		Id:         "Params",
		NTType:     17,
		Index:      30,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `Param : Rhs	<<  >>`,
		Id:         "Param",
		NTType:     18,
		Index:      31,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BoolOp : "&&"	<<  >>`,
		Id:         "BoolOp",
		NTType:     19,
		Index:      32,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BoolOp : "||"	<<  >>`,
		Id:         "BoolOp",
		NTType:     19,
		Index:      33,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RelOp : "<="	<<  >>`,
		Id:         "RelOp",
		NTType:     20,
		Index:      34,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RelOp : ">="	<<  >>`,
		Id:         "RelOp",
		NTType:     20,
		Index:      35,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RelOp : "!="	<<  >>`,
		Id:         "RelOp",
		NTType:     20,
		Index:      36,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RelOp : "=="	<<  >>`,
		Id:         "RelOp",
		NTType:     20,
		Index:      37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RelOp : "<"	<<  >>`,
		Id:         "RelOp",
		NTType:     20,
		Index:      38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RelOp : ">"	<<  >>`,
		Id:         "RelOp",
		NTType:     20,
		Index:      39,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ArithOp1 : "/"	<<  >>`,
		Id:         "ArithOp1",
		NTType:     21,
		Index:      40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ArithOp1 : "*"	<<  >>`,
		Id:         "ArithOp1",
		NTType:     21,
		Index:      41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ArithOp2 : "+"	<<  >>`,
		Id:         "ArithOp2",
		NTType:     22,
		Index:      42,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ArithOp2 : "-"	<<  >>`,
		Id:         "ArithOp2",
		NTType:     22,
		Index:      43,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expression : Expression ArithOp2 Term	<<  >>`,
		Id:         "Expression",
		NTType:     23,
		Index:      44,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expression : Expression RelOp Term	<<  >>`,
		Id:         "Expression",
		NTType:     23,
		Index:      45,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expression : Expression BoolOp Term	<<  >>`,
		Id:         "Expression",
		NTType:     23,
		Index:      46,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expression : Term	<<  >>`,
		Id:         "Expression",
		NTType:     23,
		Index:      47,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term : Term ArithOp1 Factor	<<  >>`,
		Id:         "Term",
		NTType:     24,
		Index:      48,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term : Factor	<<  >>`,
		Id:         "Term",
		NTType:     24,
		Index:      49,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : "(" Expression ")"	<<  >>`,
		Id:         "Factor",
		NTType:     25,
		Index:      50,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : identifier	<<  >>`,
		Id:         "Factor",
		NTType:     25,
		Index:      51,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : integerLit	<<  >>`,
		Id:         "Factor",
		NTType:     25,
		Index:      52,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : floatLit	<<  >>`,
		Id:         "Factor",
		NTType:     25,
		Index:      53,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : FunctionCall	<<  >>`,
		Id:         "Factor",
		NTType:     25,
		Index:      54,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
}
