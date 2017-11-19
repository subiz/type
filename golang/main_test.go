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
	ret = ts.Evaluate(Boolean, nil, Empty, nil)
	if ret == false {
		t.Fatal("must be true")
	}

	ret = ts.Evaluate(String, nil, Empty, nil)
	if ret == false {
		t.Fatal("must be true")
	}

	ret = ts.Evaluate(Number, nil, Empty, nil)
	if ret == false {
		t.Fatal("must be true")
	}

	ret = ts.Evaluate(Strings, nil, Empty, nil)
	if ret == false {
		t.Fatal("must be true")
	}
}

func TestBool(t *testing.T) {
	TearUp(t)
	var truestr = "true"
	ret := ts.Evaluate(Boolean, truestr, True, nil)
	if ret == false {
		t.Fatal("must be true")
	}

	ret = ts.Evaluate(Boolean, truestr, False, nil)
	if ret == true {
		t.Fatal("must be false")
	}

	ret = ts.Evaluate(Boolean, nil, True, nil)
	if ret == true {
		t.Fatal("must be false")
	}
}

func TestString(t *testing.T) {
	TearUp(t)
	var ret bool
	str := "ab"
	obj := "abc"
	ret = ts.Evaluate(String, obj, StartsWith, str)
	if ret == false {
		t.Fatal("must be true")
	}
	str = "a.c"
	obj = "abc"
	ret = ts.Evaluate(String, obj, Regex, str)
	if ret == false {
		t.Fatal("must be true")
	}
	obj = "abc"
	ret = ts.Evaluate(String, obj, In, []string{"abc", "bcd"})
	if ret == false {
		t.Fatal("must be true")
	}

	obj = "abc"
	ret = ts.Evaluate(String, obj, NotIn, []string{"abc", "bcd"})
	if ret {
		t.Fatal("must be false")
	}
}

func TestStrings(t *testing.T) {
	TearUp(t)
	var ret bool
	ret = ts.Evaluate(Strings, []string{"abc"}, StartsWith, "ab")
	if !ret {
		t.Fatal("must be true")
	}
	ret = ts.Evaluate(Strings, []string{"123", "abc"}, Regex, "a.c")
	if !ret {
		t.Fatal("must be true")
	}
	ret = ts.Evaluate(Strings, []string{"abc", "efg"}, In, []string{"abc", "bcd"})
	if !ret {
		t.Fatal("must be true")
	}
	ret = ts.Evaluate(Strings, []string{"abc"}, NotIn, []string{"abc", "bcd"})
	if ret {
		t.Fatal("must be false")
	}
	ret = ts.Evaluate(Strings, []string{"abc", "eft", "bcd"}, Superset, []string{"abc", "bcd"})
	if !ret {
		t.Fatal("must be true")
	}
	ret = ts.Evaluate(Strings, []string{"abc"}, Subset, []string{"abc", "bcd"})
	if !ret {
		t.Fatal("must be true")
	}
	ret = ts.Evaluate(Strings, []string{"1", "2"}, Eq, []int{1, 2})
	if !ret {
		t.Fatal("must be true")
	}
}

func TestNumber(t *testing.T) {
	TearUp(t)
	var ret bool
	str := "-20.0"
	obj := "-20"
	ret = ts.Evaluate(Number, obj, Eq, str)
	if !ret {
		t.Fatal("must be true")
	}
	str = "1"
	obj = "5.4"
	ret = ts.Evaluate(Number, obj, Gt, str)
	if !ret {
		t.Fatal("must be true")
	}
	obj = "5"
	ret = ts.Evaluate(Number, obj, In, []int{25, 5})
	if !ret {
		t.Fatal("must be true")
	}

	ret = ts.Evaluate(Number, obj, NotIn, []string{"25.0", "5.0"})
	if ret {
		t.Fatal("must be false")
	}

	obj = "4"
	ret = ts.Evaluate(Number, obj, InRange, []int{1, 20})
	if !ret {
		t.Fatal("must be true")
	}

	ret = ts.Evaluate(Number, obj, NotInRange, []string{"1", "20"})
	if ret {
		t.Fatal("must be true")
	}
}
