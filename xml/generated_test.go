// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package xml

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/coveo/gotemplate/errors"
)

var strFixture = xmlList(xmlListHelper.NewStringList(strings.Split("Hello World, I'm Foo Bar!", " ")...).AsArray())

func Test_list_Append(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		l      xmlIList
		values []interface{}
		want   xmlIList
	}{
		{"Empty", xmlList{}, []interface{}{1, 2, 3}, xmlList{1, 2, 3}},
		{"List of int", xmlList{1, 2, 3}, []interface{}{4, 5}, xmlList{1, 2, 3, 4, 5}},
		{"List of string", strFixture, []interface{}{"That's all folks!"}, xmlList{"Hello", "World,", "I'm", "Foo", "Bar!", "That's all folks!"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Append(tt.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlList.Append():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_Prepend(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		l      xmlIList
		values []interface{}
		want   xmlIList
	}{
		{"Empty", xmlList{}, []interface{}{1, 2, 3}, xmlList{1, 2, 3}},
		{"List of int", xmlList{1, 2, 3}, []interface{}{4, 5}, xmlList{4, 5, 1, 2, 3}},
		{"List of string", strFixture, []interface{}{"That's all folks!"}, xmlList{"That's all folks!", "Hello", "World,", "I'm", "Foo", "Bar!"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Prepend(tt.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlList.Prepend():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_AsArray(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    xmlList
		want []interface{}
	}{
		{"Empty List", xmlList{}, []interface{}{}},
		{"List of int", xmlList{1, 2, 3}, []interface{}{1, 2, 3}},
		{"List of string", strFixture, []interface{}{"Hello", "World,", "I'm", "Foo", "Bar!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.AsArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlList.AsList():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_XmlList_Strings(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    xmlList
		want []string
	}{
		{"Empty List", xmlList{}, []string{}},
		{"List of int", xmlList{1, 2, 3}, []string{"1", "2", "3"}},
		{"List of string", strFixture, []string{"Hello", "World,", "I'm", "Foo", "Bar!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Strings(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlList.Strings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list_Capacity(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    xmlIList
		want int
	}{
		{"Empty List with 100 spaces", xmlListHelper.CreateList(0, 100), 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Capacity(); got != tt.want {
				t.Errorf("XmlList.Capacity() = %v, want %v", got, tt.want)
			}
			if tt.l.Capacity() != tt.l.Cap() {
				t.Errorf("Cap and Capacity return different values")
			}
		})
	}
}

func Test_list_Clone(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    xmlList
		want xmlIList
	}{
		{"Empty List", xmlList{}, xmlList{}},
		{"List of int", xmlList{1, 2, 3}, xmlList{1, 2, 3}},
		{"List of string", strFixture, xmlList{"Hello", "World,", "I'm", "Foo", "Bar!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlList.Clone():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_Get(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		l     xmlList
		index int
		want  interface{}
	}{
		{"Empty List", xmlList{}, 0, nil},
		{"Negative index", xmlList{}, -1, nil},
		{"List of int", xmlList{1, 2, 3}, 0, 1},
		{"List of string", strFixture, 1, "World,"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Get(tt.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlList.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list_Len(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    xmlList
		want int
	}{
		{"Empty List", xmlList{}, 0},
		{"List of int", xmlList{1, 2, 3}, 3},
		{"List of string", strFixture, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Len(); got != tt.want {
				t.Errorf("XmlList.Len() = %v, want %v", got, tt.want)
			}
			if tt.l.Len() != tt.l.Count() {
				t.Errorf("Len and Count return different values")
			}
		})
	}
}

func Test_CreateList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		args    []int
		want    xmlIList
		wantErr bool
	}{
		{"Empty", nil, xmlList{}, false},
		{"With nil elements", []int{10}, make(xmlList, 10), false},
		{"With capacity", []int{0, 10}, make(xmlList, 0, 10), false},
		{"Too much args", []int{0, 10, 1}, nil, true},
	}
	for _, tt := range tests {
		var got xmlIList
		var err error
		func() {
			defer func() { err = errors.Trap(err, recover()) }()
			got = xmlListHelper.CreateList(tt.args...)
		}()
		if (err != nil) != tt.wantErr {
			t.Errorf("CreateList() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if err != nil {
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("CreateList():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
		}
		if got.Capacity() != tt.want.Cap() {
			t.Errorf("CreateList() capacity:\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got.Cap(), tt.want.Capacity())
		}
	}
}

func Test_list_Create(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    xmlList
		args []int
		want xmlIList
	}{
		{"Empty", nil, nil, xmlList{}},
		{"Existing List", xmlList{1, 2}, nil, xmlList{}},
		{"With Empty spaces", xmlList{1, 2}, []int{5}, xmlList{nil, nil, nil, nil, nil}},
		{"With Capacity", xmlList{1, 2}, []int{0, 5}, xmlListHelper.CreateList(0, 5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.l.Create(tt.args...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlList.Create():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
			if got.Capacity() != tt.want.Capacity() {
				t.Errorf("XmlList.Create() capacity:\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got.Capacity(), tt.want.Capacity())
			}
		})
	}
}

func Test_list_New(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    xmlList
		args []interface{}
		want xmlIList
	}{
		{"Empty", nil, nil, xmlList{}},
		{"Existing List", xmlList{1, 2}, nil, xmlList{}},
		{"With elements", xmlList{1, 2}, []interface{}{3, 4, 5}, xmlList{3, 4, 5}},
		{"With strings", xmlList{1, 2}, []interface{}{"Hello", "World"}, xmlList{"Hello", "World"}},
		{"With nothing", xmlList{1, 2}, []interface{}{}, xmlList{}},
		{"With nil", xmlList{1, 2}, nil, xmlList{}},
		{"Adding array", xmlList{1, 2}, []interface{}{xmlList{3, 4}}, xmlList{3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.New(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlList.Create():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_CreateDict(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		l       xmlList
		args    []int
		want    xmlIDict
		wantErr bool
	}{
		{"Empty", nil, nil, xmlDict{}, false},
		{"With capacity", nil, []int{10}, xmlDict{}, false},
		{"With too much parameter", nil, []int{10, 1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got xmlIDict
			var err error
			func() {
				defer func() { err = errors.Trap(err, recover()) }()
				got = tt.l.CreateDict(tt.args...)
			}()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlList.CreateDict():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("XmlList.CreateDict() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_list_Contains(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    xmlList
		args []interface{}
		want bool
	}{
		{"Empty List", nil, []interface{}{}, false},
		{"Search nothing", xmlList{1}, nil, true},
		{"Search nothing 2", xmlList{1}, []interface{}{}, true},
		{"Not there", xmlList{1}, []interface{}{2}, false},
		{"Included", xmlList{1, 2}, []interface{}{2}, true},
		{"Partially there", xmlList{1, 2}, []interface{}{2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Contains(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlList.Contains():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_Without(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    xmlList
		args []interface{}
		want xmlList
	}{
		{"Empty List", nil, []interface{}{}, xmlList{}},
		{"Remove nothing", xmlList{1}, nil, xmlList{1}},
		{"Remove nothing 2", xmlList{1}, []interface{}{}, xmlList{1}},
		{"Not there", xmlList{1}, []interface{}{2}, xmlList{1}},
		{"Included", xmlList{1, 2}, []interface{}{2}, xmlList{1}},
		{"Partially there", xmlList{1, 2}, []interface{}{2, 3}, xmlList{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Without(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlList.Without():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_Unique(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    xmlList
		want xmlList
	}{
		{"Empty List", nil, xmlList{}},
		{"Remove nothing", xmlList{1}, xmlList{1}},
		{"Duplicates following", xmlList{1, 1, 2, 3}, xmlList{1, 2, 3}},
		{"Duplicates not following", xmlList{1, 2, 3, 1, 2, 3, 4}, xmlList{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Unique(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlList.Unique():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}
func Test_list_Reverse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    xmlList
		want xmlIList
	}{
		{"Empty List", xmlList{}, xmlList{}},
		{"List of int", xmlList{1, 2, 3}, xmlList{3, 2, 1}},
		{"List of string", strFixture, xmlList{"Bar!", "Foo", "I'm", "World,", "Hello"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.l.Clone()
			if got := l.Reverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlList.Reverse():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_Set(t *testing.T) {
	t.Parallel()

	type args struct {
		i int
		v interface{}
	}
	tests := []struct {
		name    string
		l       xmlIList
		args    args
		want    xmlIList
		wantErr bool
	}{
		{"Empty", xmlList{}, args{2, 1}, xmlList{nil, nil, 1}, false},
		{"List of int", xmlList{1, 2, 3}, args{0, 10}, xmlList{10, 2, 3}, false},
		{"List of string", strFixture, args{2, "You're"}, xmlList{"Hello", "World,", "You're", "Foo", "Bar!"}, false},
		{"Negative", xmlList{}, args{-1, "negative value"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.Clone().Set(tt.args.i, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("XmlList.Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlList.Set():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

var mapFixture = map[string]interface{}{
	"int":     123,
	"float":   1.23,
	"string":  "Foo bar",
	"list":    []interface{}{1, "two"},
	"listInt": []int{1, 2, 3},
	"map": map[string]interface{}{
		"sub1": 1,
		"sub2": "two",
	},
	"mapInt": map[int]interface{}{
		1: 1,
		2: "two",
	},
}

var dictFixture = xmlDict(xmlDictHelper.AsDictionary(mapFixture).AsMap())

func dumpKeys(t *testing.T, d1, d2 xmlIDict) {
	t.Parallel()

	for key := range d1.AsMap() {
		v1, v2 := d1.Get(key), d2.Get(key)
		if reflect.DeepEqual(v1, v2) {
			continue
		}
		t.Logf("'%[1]v' = %[2]v (%[2]T) vs %[3]v (%[3]T)", key, v1, v2)
	}
}

func Test_dict_AsMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    xmlDict
		want map[string]interface{}
	}{
		{"Nil", nil, nil},
		{"Empty", xmlDict{}, map[string]interface{}{}},
		{"Map", dictFixture, map[string]interface{}(dictFixture)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.AsMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlDict.AsMap():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_dict_Clone(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    xmlDict
		keys []interface{}
		want xmlIDict
	}{
		{"Nil", nil, nil, xmlDict{}},
		{"Empty", xmlDict{}, nil, xmlDict{}},
		{"Map", dictFixture, nil, dictFixture},
		{"Map with Fields", dictFixture, []interface{}{"int", "list"}, xmlDict(dictFixture).Omit("float", "string", "listInt", "map", "mapInt")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.d.Clone(tt.keys...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlDict.Clone():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
				dumpKeys(t, got, tt.want)
			}

			// Ensure that the copy is distinct from the original
			got.Set("NewFields", "Test")
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Should be different:\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
			if !got.Has("NewFields") || !reflect.DeepEqual(got.Get("NewFields"), "Test") {
				t.Errorf("Element has not been added")
			}
			if got.Len() != tt.want.Count()+1 {
				t.Errorf("Len and Count don't return the same value")
			}
		})
	}
}

func Test_XmlDict_CreateList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		d            xmlDict
		args         []int
		want         xmlIList
		wantLen      int
		wantCapacity int
	}{
		{"Nil", nil, nil, xmlList{}, 0, 0},
		{"Empty", xmlDict{}, nil, xmlList{}, 0, 0},
		{"Map", dictFixture, nil, xmlList{}, 0, 0},
		{"Map with size", dictFixture, []int{3}, xmlList{nil, nil, nil}, 3, 3},
		{"Map with capacity", dictFixture, []int{0, 10}, xmlList{}, 0, 10},
		{"Map with size&capacity", dictFixture, []int{3, 10}, xmlList{nil, nil, nil}, 3, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.d.CreateList(tt.args...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlDict.CreateList() = %v, want %v", got, tt.want)
			}
			if got.Len() != tt.wantLen || got.Cap() != tt.wantCapacity {
				t.Errorf("XmlDict.CreateList() size: %d, %d vs %d, %d", got.Len(), got.Cap(), tt.wantLen, tt.wantCapacity)
			}
		})
	}
}

func Test_dict_Create(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		d       xmlDict
		args    []int
		want    xmlIDict
		wantErr bool
	}{
		{"Empty", nil, nil, xmlDict{}, false},
		{"With capacity", nil, []int{10}, xmlDict{}, false},
		{"With too much parameter", nil, []int{10, 1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got xmlIDict
			var err error
			func() {
				defer func() { err = errors.Trap(err, recover()) }()
				got = tt.d.Create(tt.args...)
			}()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlDict.Create():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("XmlList.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_dict_Default(t *testing.T) {
	t.Parallel()

	type args struct {
		key    interface{}
		defVal interface{}
	}
	tests := []struct {
		name string
		d    xmlDict
		args args
		want interface{}
	}{
		{"Empty", nil, args{"Foo", "Bar"}, "Bar"},
		{"Map int", dictFixture, args{"int", 1}, 123},
		{"Map float", dictFixture, args{"float", 1}, 1.23},
		{"Map Non existant", dictFixture, args{"Foo", "Bar"}, "Bar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Default(tt.args.key, tt.args.defVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlDict.Default() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dict_Delete(t *testing.T) {
	t.Parallel()

	type args struct {
		key  interface{}
		keys []interface{}
	}
	tests := []struct {
		name    string
		d       xmlDict
		args    args
		want    xmlIDict
		wantErr bool
	}{
		{"Empty", nil, args{}, xmlDict{}, true},
		{"Map", dictFixture, args{}, dictFixture, true},
		{"Non existant key", dictFixture, args{"Test", nil}, dictFixture, true},
		{"Map with keys", dictFixture, args{"int", []interface{}{"list"}}, dictFixture.Clone("float", "string", "listInt", "map", "mapInt"), false},
		{"Map with keys + non existant", dictFixture, args{"int", []interface{}{"list", "Test"}}, dictFixture.Clone("float", "string", "listInt", "map", "mapInt"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d.Clone()
			got, err := d.Delete(tt.args.key, tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("XmlDict.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlDict.Delete():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
				dumpKeys(t, got, tt.want)
			}
		})
	}
}

func Test_dict_Flush(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    xmlDict
		keys []interface{}
		want xmlIDict
	}{
		{"Empty", nil, nil, xmlDict{}},
		{"Map", dictFixture, nil, xmlDict{}},
		{"Non existant key", dictFixture, []interface{}{"Test"}, dictFixture},
		{"Map with keys", dictFixture, []interface{}{"int", "list"}, dictFixture.Clone("float", "string", "listInt", "map", "mapInt")},
		{"Map with keys + non existant", dictFixture, []interface{}{"int", "list", "Test"}, dictFixture.Clone("float", "string", "listInt", "map", "mapInt")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d.Clone()
			got := d.Flush(tt.keys...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlDict.Flush():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
				dumpKeys(t, got, tt.want)
			}
			if !reflect.DeepEqual(d, got) {
				t.Errorf("Should be equal after: %v, want %v", d, got)
				dumpKeys(t, d, got)
			}
		})
	}
}

func Test_dict_Keys(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    xmlDict
		want xmlIList
	}{
		{"Empty", nil, xmlList{}},
		{"Map", dictFixture, xmlList{"float", "int", "list", "listInt", "map", "mapInt", "string"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.GetKeys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlDict.GetKeys():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_dict_KeysAsString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    xmlDict
		want []string
	}{
		{"Empty", nil, []string{}},
		{"Map", dictFixture, []string{"float", "int", "list", "listInt", "map", "mapInt", "string"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.KeysAsString(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlDict.KeysAsString():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_dict_Merge(t *testing.T) {
	t.Parallel()

	adding1 := xmlDict{
		"int":        1000,
		"Add1Int":    1,
		"Add1String": "string",
	}
	adding2 := xmlDict{
		"Add2Int":    1,
		"Add2String": "string",
		"map": map[string]interface{}{
			"sub1":   2,
			"newVal": "NewValue",
		},
	}
	type args struct {
		xmlDict xmlIDict
		dicts   []xmlIDict
	}
	tests := []struct {
		name string
		d    xmlDict
		args args
		want xmlIDict
	}{
		{"Empty", nil, args{nil, []xmlIDict{}}, xmlDict{}},
		{"Add map to empty", nil, args{dictFixture, []xmlIDict{}}, dictFixture},
		{"Add map to same map", dictFixture, args{dictFixture, []xmlIDict{}}, dictFixture},
		{"Add empty to map", dictFixture, args{nil, []xmlIDict{}}, dictFixture},
		{"Add new1 to map", dictFixture, args{adding1, []xmlIDict{}}, dictFixture.Clone().Merge(adding1)},
		{"Add new2 to map", dictFixture, args{adding2, []xmlIDict{}}, dictFixture.Clone().Merge(adding2)},
		{"Add new1 & new2 to map", dictFixture, args{adding1, []xmlIDict{adding2}}, dictFixture.Clone().Merge(adding1, adding2)},
		{"Add new1 & new2 to map", dictFixture, args{adding1, []xmlIDict{adding2}}, dictFixture.Clone().Merge(adding1).Merge(adding2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d.Clone()
			got := d.Merge(tt.args.xmlDict, tt.args.dicts...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlDict.Merge():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
				dumpKeys(t, got, tt.want)
			}
		})
	}
}

func Test_dict_Values(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    xmlDict
		want xmlIList
	}{
		{"Empty", nil, xmlList{}},
		{"Map", dictFixture, xmlList{1.23, 123, xmlList{1, "two"}, xmlList{1, 2, 3}, xmlDict{"sub1": 1, "sub2": "two"}, xmlDict{"1": 1, "2": "two"}, "Foo bar"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.GetValues(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlDict.GetValues():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_dict_Add(t *testing.T) {
	t.Parallel()

	type args struct {
		key interface{}
		v   interface{}
	}
	tests := []struct {
		name string
		d    xmlDict
		args args
		want xmlIDict
	}{
		{"Empty", nil, args{"A", 1}, xmlDict{"A": 1}},
		{"With element", xmlDict{"A": 1}, args{"A", 2}, xmlDict{"A": xmlList{1, 2}}},
		{"With element, another value", xmlDict{"A": 1}, args{"B", 2}, xmlDict{"A": 1, "B": 2}},
		{"With list element", xmlDict{"A": xmlList{1, 2}}, args{"A", 3}, xmlDict{"A": xmlList{1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Add(tt.args.key, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlDict.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dict_Set(t *testing.T) {
	t.Parallel()

	type args struct {
		key interface{}
		v   interface{}
	}
	tests := []struct {
		name string
		d    xmlDict
		args args
		want xmlIDict
	}{
		{"Empty", nil, args{"A", 1}, xmlDict{"A": 1}},
		{"With element", xmlDict{"A": 1}, args{"A", 2}, xmlDict{"A": 2}},
		{"With element, another value", xmlDict{"A": 1}, args{"B", 2}, xmlDict{"A": 1, "B": 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Set(tt.args.key, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlDict.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dict_Transpose(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    xmlDict
		want xmlIDict
	}{
		{"Empty", nil, xmlDict{}},
		{"Base", xmlDict{"A": 1}, xmlDict{"1": "A"}},
		{"Multiple", xmlDict{"A": 1, "B": 2, "C": 1}, xmlDict{"1": xmlList{"A", "C"}, "2": "B"}},
		{"List", xmlDict{"A": []int{1, 2, 3}, "B": 2, "C": 3}, xmlDict{"1": "A", "2": xmlList{"A", "B"}, "3": xmlList{"A", "C"}}},
		{"Complex", xmlDict{"A": xmlDict{"1": 1, "2": 2}, "B": 2, "C": 3}, xmlDict{"2": "B", "3": "C", fmt.Sprint(xmlDict{"1": 1, "2": 2}): "A"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Transpose(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlDict.Transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}