package typesystem

import (
	"encoding/json"
	"testing"
)

var ts = NewTypeSystem()

func TearUp(t *testing.T) {

}






// func TestEmpty(t *testing.T) {
// 	TearUp(t)
// 	var ret bool
// 	ret = ts.Evaluate(Boolean, nil, Empty, nil)
// 	if ret == false {
// 		t.Fatal("must be true")
// 	}

// 	ret = ts.Evaluate(String, nil, Empty, nil)
// 	if ret == false {
// 		t.Fatal("must be true")
// 	}

// 	ret = ts.Evaluate(Number, nil, Empty, nil)
// 	if ret == false {
// 		t.Fatal("must be true")
// 	}

// 	ret = ts.Evaluate(Strings, nil, Empty, nil)
// 	if ret == false {
// 		t.Fatal("must be true")
// 	}
// }

func TestBool(t *testing.T) {
	TearUp(t)
	var truestr = "true"
	ret, err := ts.Evaluate(Boolean, truestr, True, "true")
	if ret == false {
		t.Fatal("must be true")
	}
	if err != nil {
		t.Fatal(err)
	}

	ret, err = ts.Evaluate(Boolean, truestr, False, "false")
	if ret == true {
		t.Fatal("must be false")
	}

	if err != nil {
		t.Fatal(err)
	}

	ret, err = ts.Evaluate(Boolean, nil, True, "false")
	if ret == true {
		t.Fatal("must be false")
	}

	if err != nil {
		t.Fatal(err)
	}
}

// func TestString(t *testing.T) {
// 	TearUp(t)
// 	var ret bool
// 	str := "\"ab\""
// 	obj := "abc"
// 	ret, err := ts.Evaluate(String, obj, StartsWith, str)
// 	if ret == false {
// 		t.Fatal("must be true")
// 	}

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	str = "a.c"
// 	obj = "abc"
// 	ret, err = ts.Evaluate(String, obj, Regex, str)
// 	if ret == false {
// 		t.Fatal("must be true")
// 	}

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	obj = "abc"
// 	strbyte, _ := json.Marshal([]string{"abc", "bcd"})
// 	ret, err = ts.Evaluate(String, obj, In, string(strbyte))
// 	if ret == false {
// 		t.Fatal("must be true")
// 	}

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	obj = "abc"
// 	strbyte, _ = json.Marshal([]string{"abc", "bcd"})
// 	ret, err = ts.Evaluate(String, obj, NotIn, string(strbyte))
// 	if ret {
// 		t.Fatal("must be false")
// 	}

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestStrings(t *testing.T) {
// 	TearUp(t)
// 	var ret bool
// 	ret, err := ts.Evaluate(Strings, []string{"abc"}, StartsWith, "ab")
// 	if !ret {
// 		t.Fatal("must be true")
// 	}

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	ret, err = ts.Evaluate(Strings, []string{"123", "abc"}, Regex, "a.c")
// 	if !ret {
// 		t.Fatal("must be true")
// 	}

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	strb, _ := json.Marshal([]string{"abc", "bcd"})
// 	str := string(strb)
// 	ret, err = ts.Evaluate(Strings, []string{"abc", "efg"}, In, str)
// 	if !ret {
// 		t.Fatal("must be true")
// 	}

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	ret, err = ts.Evaluate(Strings, []string{"abc"}, NotIn, str)
// 	if ret {
// 		t.Fatal("must be false")
// 	}

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	ret, err = ts.Evaluate(Strings, []string{"abc", "eft", "bcd"}, Superset, str)
// 	if !ret {
// 		t.Fatal("must be true")
// 	}

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	ret, err = ts.Evaluate(Strings, []string{"abc"}, Subset, str)
// 	if !ret {
// 		t.Fatal("must be true")
// 	}

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	strb, _ = json.Marshal([]int{1, 2})
// 	str = string(strb)
// 	ret, err = ts.Evaluate(Strings, []string{"1", "2"}, Eq, str)
// 	if !ret {
// 		t.Fatal("must be true")
// 	}

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

func TestNumber(t *testing.T) {
	TearUp(t)
	var ret bool
	str := "-20.0"
	obj := "-20"
	ret, err := ts.Evaluate(Number, obj, Eq, str)
	if !ret {
		t.Fatal("must be true")
	}
	if err != nil {
		t.Fatal(err)
	}

	str = "1"
	obj = "5.4"
	ret, err = ts.Evaluate(Number, obj, Gt, str)
	if !ret {
		t.Fatal("must be true")
	}

	if err != nil {
		t.Fatal(err)
	}

	obj = "5"
	strb, _ := json.Marshal([]int{25, 5})
	str = string(strb)
	ret, err = ts.Evaluate(Number, obj, In, str)
	if !ret {
		t.Fatal("must be true")
	}

	if err != nil {
		t.Fatal(err)
	}

	ret, err = ts.Evaluate(Number, obj, NotIn, str)
	if ret {
		t.Fatal("must be false")
	}

	if err != nil {
		t.Fatal(err)
	}

	obj = "4"
	strb, _ = json.Marshal([]int{1, 20})
	str = string(strb)
	ret, err = ts.Evaluate(Number, obj, InRange, str)
	if !ret {
		t.Fatal("must be true")
	}

	if err != nil {
		t.Fatal(err)
	}

	// ret, err = ts.Evaluate(Number, obj, NotInRange, str)
	//if ret {
	//t.Fatal("must be false")
	//}

	if err != nil {
		t.Fatal(err)
	}


	obj = "114"
	strb, _ = json.Marshal([]int{1, 20})
	str = string(strb)
	ret, err = ts.Evaluate(Number, obj, NotInRange, str)
	if !ret {
		t.Fatal("must be true")
	}

	if err != nil {
		t.Fatal(err)
	}
}
