package typesystem

import (
	"errors"
	"bitbucket.org/subiz/gocommon"
)

type ITypeSystem interface {
	Evaluate(typename string, object interface{}, op string, value interface{}) bool
}

type iType interface {
	Evaluate(object interface{}, opstring string, value interface{}) bool
}

type TypeSystem struct {
	stringts iType
	numberts iType
	boolts iType
	stringsts iType
}

const (
	Ab = "ab"
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

func NewTypeSystem() ITypeSystem {
	return &TypeSystem{
		stringts: NewStringType(),
		numberts: NewNumberType(),
		stringsts: NewStringsType(),
		boolts: NewBoolType(),
	}
}

func (t *TypeSystem) Evaluate(typename string, object interface{}, op string, value interface{}) bool {
	switch typename {
	case "string":
		return t.stringts.Evaluate(object, op, value)
	case "number":
		return t.numberts.Evaluate(object, op, value)
	case "boolean":
		return t.boolts.Evaluate(object, op, value)
	case "set of string":
		return t.stringsts.Evaluate(object, op, value)
	case "date":
	default:
		common.Panic(errors.New("invalid type"), "unsupported type " + typename)
	}
	return false
}
