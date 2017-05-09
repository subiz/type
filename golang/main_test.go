package typesystem

import (
	"testing"
)

var ts = NewTypeSystem()
func TearUp(t *testing.T) {

}

func TestEmpty(t *testing.T) {
	TearUp(t)

	var ret bool
	ret = ts.Evaluate("boolean", nil, Empty, nil)
	if ret == false {
		t.Fatal("must be true")
	}

	ret = ts.Evaluate("string", nil, Empty, nil)
	if ret == false {
		t.Fatal("must be true")
	}

	ret = ts.Evaluate("number", nil, Empty, nil)
	if ret == false {
		t.Fatal("must be true")
	}

	ret = ts.Evaluate("set of string", nil, Empty, nil)
	if ret == false {
		t.Fatal("must be true")
	}
}

func TestBool(t *testing.T) {
	TearUp(t)
	var truestr = "true"
	ret := ts.Evaluate("boolean", &truestr, True, nil)
	if ret == false {
		t.Fatal("must be true")
	}

	ret = ts.Evaluate("boolean", &truestr, False, nil)
	if ret == true {
		t.Fatal("must be false")
	}

	ret = ts.Evaluate("boolean", nil, True, nil)
	if ret == true {
		t.Fatal("must be false")
	}
}

func TestString(t *testing.T) {
	TearUp(t)
	var ret bool
	str := "ab"
	obj := "abc"
	ret = ts.Evaluate("string", &obj, StartsWith, str)
	if ret == false {
		t.Fatal("must be true")
	}
	str = "a.c"
	obj = "abc"
	ret = ts.Evaluate("string", &obj, Regex, str)
	if ret == false {
		t.Fatal("must be true")
	}
	obj = "abc"
	ret = ts.Evaluate("string", &obj, In, []string{"abc", "bcd"})
	if ret == false {
		t.Fatal("must be true")
	}

	obj = "abc"
	ret = ts.Evaluate("string", &obj, NotIn, []string{"abc", "bcd"})
	if ret {
		t.Fatal("must be false")
	}
}

func TestStrings(t *testing.T) {
	TearUp(t)
	var ret bool
	ret = ts.Evaluate("set of string", []string{"abc"}, StartsWith, "ab")
	if !ret {
		t.Fatal("must be true")
	}
	ret = ts.Evaluate("set of string", []string{"123", "abc"}, Regex, "a.c")
	if !ret {
		t.Fatal("must be true")
	}
	ret = ts.Evaluate("set of string", []string{"abc", "efg"}, In, []string{"abc", "bcd"})
	if !ret {
		t.Fatal("must be true")
	}
	ret = ts.Evaluate("set of string", []string{"abc"}, NotIn, []string{"abc", "bcd"})
	if ret {
		t.Fatal("must be false")
	}
	ret = ts.Evaluate("set of string", []string{"abc", "eft", "bcd"}, Superset, []string{"abc", "bcd"})
	if !ret {
		t.Fatal("must be true")
	}
	ret = ts.Evaluate("set of string", []string{"abc"}, Subset, []string{"abc", "bcd"})
	if !ret {
		t.Fatal("must be true")
	}
	ret = ts.Evaluate("set of string", []string{"abc", "bcd"}, Eq, []string{"abc", "bcd"})
	if !ret {
		t.Fatal("must be true")
	}
}

func TestNumber(t *testing.T) {
	TearUp(t)
	var ret bool
	str := "-20.0"
	obj := "-20"
	ret = ts.Evaluate("number", &obj, Eq, str)
	if !ret {
		t.Fatal("must be true")
	}
	str = "1"
	obj = "5.4"
	ret = ts.Evaluate("number", &obj, Gt, str)
	if !ret {
		t.Fatal("must be true")
	}
	obj = "5"
	ret = ts.Evaluate("number", &obj, In, []string{"25", "5"})
	if !ret {
		t.Fatal("must be true")
	}

	obj = "4"
	ret = ts.Evaluate("number", &obj, InRange, []string{"1", "20"})
	if !ret {
		t.Fatal("must be true")
	}
}
