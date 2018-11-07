package typesystem

import (
	"encoding/json"
	"fmt"
	"testing"
)

var ts = NewTypeSystem()

func TearUp(t *testing.T) {

}

func TestDateConvToEls(t *testing.T) {
	fmt.Println("---- Date ----")
	query, err := ts.datets.ConvToEls("age", "eq", "10")
	fmt.Println(query)
	if err != nil {
		t.Fatal(err)
	}

	query, err = ts.datets.ConvToEls("age", "ne", "10")
	fmt.Println(query)
	if err != nil {
		t.Fatal(err)
	}

	query, err = ts.datets.ConvToEls("age", "gt", "10")
	fmt.Println(query)
	if err != nil {
		t.Fatal(err)
	}

	query, err = ts.datets.ConvToEls("age", "lt", "10")
	fmt.Println(query)
	if err != nil {
		t.Fatal(err)
	}

	query, err = ts.datets.ConvToEls("age", "gte", "10")
	fmt.Println(query)
	if err != nil {
		t.Fatal(err)
	}

	arrInt := []int{5, 25}
	arrInts, _ := json.Marshal(arrInt)
	query, err = ts.datets.ConvToEls("age", "inRange", string(arrInts))
	fmt.Println(query)
	if err != nil {
		t.Fatal(err)
	}

	query, err = ts.datets.ConvToEls("age", "notInRange", string(arrInts))
	fmt.Println(query)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBoolConvToEls(t *testing.T) {
	fmt.Println("---- Boolean ----")
	query, err := ts.boolts.ConvToEls("sex", "true", "")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(query)

	query, err = ts.boolts.ConvToEls("sex", "false", "")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(query)

}

func TestStrConvToEls(t *testing.T) {
	fmt.Println("---- String ----")
	query, err := ts.stringts.ConvToEls("name", "eq", "\"viet\"")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(query)

	query, err = ts.stringts.ConvToEls("name", "ne", "\"viet\"")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(query)

	query, err = ts.stringts.ConvToEls("name", "begin", "\"vi\"")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(query)

	query, err = ts.stringts.ConvToEls("name", "notBegin", "\"vi\"")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(query)

	query, err = ts.stringts.ConvToEls("name", "end", "\"t\"")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(query)

	query, err = ts.stringts.ConvToEls("name", "notEnd", "\"t\"")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(query)

	query, err = ts.stringts.ConvToEls("name", "con", "\"ie\"")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(query)

	query, err = ts.stringts.ConvToEls("name", "notCon", "\"viet\"")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(query)

}

func TestNumberConvToEls(t *testing.T) {
	fmt.Println("---- Number ----")
	query, err := ts.numberts.ConvToEls("ag\"e", "eq", "10")
	fmt.Println(query)
	if err != nil {
		t.Fatal(err)
	}

	query, err = ts.numberts.ConvToEls("age", "ne", "10")
	fmt.Println(query)
	if err != nil {
		t.Fatal(err)
	}

	query, err = ts.numberts.ConvToEls("age", "gt", "10")
	fmt.Println(query)
	if err != nil {
		t.Fatal(err)
	}

	query, err = ts.numberts.ConvToEls("age", "lt", "10")
	fmt.Println(query)
	if err != nil {
		t.Fatal(err)
	}

	query, err = ts.numberts.ConvToEls("age", "gte", "10")
	fmt.Println(query)
	if err != nil {
		t.Fatal(err)
	}

	arrInt := []int{5, 25}
	arrInts, _ := json.Marshal(arrInt)
	query, err = ts.numberts.ConvToEls("age", "inRange", string(arrInts))
	fmt.Println(query)
	if err != nil {
		t.Fatal(err)
	}

	query, err = ts.numberts.ConvToEls("age", "notInRange", string(arrInts))
	fmt.Println(query)
	if err != nil {
		t.Fatal(err)
	}
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

	ret, err = ts.Evaluate(Number, obj, NotInRange, str)
	if ret {
		t.Fatal("must be true")
	}

	if err != nil {
		t.Fatal(err)
	}
}
