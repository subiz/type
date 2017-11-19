package typesystem

import (
)

type iType interface {
	Evaluate(object interface{}, opstring string, value interface{}) bool
}

// TypeSystem abc
type TypeSystem struct {
	stringts iType
	numberts iType
	boolts iType
	stringsts iType
}

const (
	//Ab abc
	Ab = "ab"
	//After abc
	After = "after"
	An = "an"
	Before = "before"
	Contains = "con"
	Diff = "diff"
	Empty = "empty"
	EndsWith = "end"
	Eq = "eq"
	False = "false"
	Gt = "gt"
	Gte = "gte"
	In = "in"
	InRange = "inRange"
	Lte = "lte"
	Lt = "lt"
	Nab = "nab"
	Nad = "nad"
	Nan = "nan"
	Ne = "ne"
	NotEmpty = "notEmpty"
	NotIn = "notIn"
	NotStartsWith = "notBegin"
	NotEndsWith = "notEnd"
	NotContains = "notCon"
	NotSuperset = "notsup"
	NotSubset = "notsub"
	NotInRange = "notInRange"
	Regex = "regex"
	Superset = "sup"
	Subset = "sub"
	StartsWith = "begin"
	True = "true"
)

const (
	Strings = "set of strings"
	String = "string"
	Number = "number"
	Boolean = "boolean"
	Date = "date"
)

// NewTypeSystem create new type system
func NewTypeSystem() *TypeSystem {
	return &TypeSystem{
		stringts: NewStringType(),
		numberts: NewNumberType(),
		stringsts: NewStringsType(),
		boolts: NewBoolType(),
	}
}

// Evaluate evalue a equation
func (t *TypeSystem) Evaluate(typename string, object interface{}, op string, value interface{}) bool {
	switch typename {
	case Strinng:
		return t.stringts.Evaluate(object, op, value)
	case Number:
		return t.numberts.Evaluate(object, op, value)
	case Boolean:
		return t.boolts.Evaluate(object, op, value)
	case Strings:
		return t.stringsts.Evaluate(object, op, value)
	case Date:
	default:
		panic("unsupported type " + typename)
	}
	return false
}
