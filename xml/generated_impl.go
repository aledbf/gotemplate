// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package xml

import (
	"strings"

	"github.com/coveooss/gotemplate/v3/collections"
	"github.com/coveooss/multilogger/errors"
)

// List implementation of IGenericList for xmlList
type List = xmlList
type xmlIList = collections.IGenericList
type xmlList []interface{}

var xmlLower = strings.ToLower("Xml") // This is required because genny capitalize the type name in strings

func (l xmlList) AsArray() []interface{}               { return []interface{}(l) }
func (l xmlList) Cap() int                             { return cap(l) }
func (l xmlList) Capacity() int                        { return cap(l) }
func (l xmlList) Clone() xmlIList                      { return xmlListHelper.Clone(l) }
func (l xmlList) Count() int                           { return len(l) }
func (l xmlList) Create(args ...int) xmlIList          { return xmlListHelper.CreateList(args...) }
func (l xmlList) CreateDict(args ...int) xmlIDict      { return xmlListHelper.CreateDictionary(args...) }
func (l xmlList) Find(element interface{}) xmlIList    { return xmlListHelper.Find(l, element, false) }
func (l xmlList) First() interface{}                   { return xmlListHelper.GetIndexes(l, 0) }
func (l xmlList) Get(indexes ...int) interface{}       { return xmlListHelper.GetIndexes(l, indexes...) }
func (l xmlList) GetKinds() xmlIList                   { return xmlListHelper.GetTypes(l, true) }
func (l xmlList) GetTypes() xmlIList                   { return xmlListHelper.GetTypes(l, false) }
func (l xmlList) Has(values ...interface{}) bool       { return l.Contains(values...) }
func (l xmlList) HasStrict(values ...interface{}) bool { return l.ContainsStrict(values...) }
func (l xmlList) Join(sep interface{}) str             { return l.StringArray().Join(sep) }
func (l xmlList) Last() interface{}                    { return xmlListHelper.GetIndexes(l, len(l)-1) }
func (l xmlList) Len() int                             { return len(l) }
func (l xmlList) New(args ...interface{}) xmlIList     { return xmlListHelper.NewList(args...) }
func (l xmlList) Reverse() xmlIList                    { return xmlListHelper.Reverse(l) }
func (l xmlList) RemoveEmpty() xmlIList                { return xmlListHelper.RemoveEmpty(l) }
func (l xmlList) RemoveNil() xmlIList                  { return xmlListHelper.RemoveNil(l) }
func (l xmlList) StringArray() strArray                { return xmlListHelper.GetStringArray(l) }
func (l xmlList) Strings() []string                    { return xmlListHelper.GetStrings(l) }
func (l xmlList) Type() str                            { return xmlListHelper.Type(l) }
func (l xmlList) TypeName() str                        { return str(xmlLower) }
func (l xmlList) Unique() xmlIList                     { return xmlListHelper.Unique(l) }

func (l xmlList) GetHelpers() (collections.IDictionaryHelper, collections.IListHelper) {
	return xmlDictHelper, xmlListHelper
}

func (l xmlList) Append(values ...interface{}) xmlIList {
	return xmlListHelper.Add(l, false, values...)
}

func (l xmlList) Contains(values ...interface{}) bool {
	return xmlListHelper.Contains(l, false, values...)
}

func (l xmlList) ContainsStrict(values ...interface{}) bool {
	return xmlListHelper.Contains(l, true, values...)
}

func (l xmlList) FindStrict(element interface{}) xmlIList {
	return xmlListHelper.Find(l, element, true)
}

func (l xmlList) Intersect(values ...interface{}) xmlIList {
	return xmlListHelper.Intersect(l, values...)
}

func (l xmlList) Pop(indexes ...int) (interface{}, xmlIList) {
	if len(indexes) == 0 {
		indexes = []int{len(l) - 1}
	}
	return l.Get(indexes...), l.Remove(indexes...)
}

func (l xmlList) Prepend(values ...interface{}) xmlIList {
	return xmlListHelper.Add(l, true, values...)
}

func (l xmlList) Remove(indexes ...int) xmlIList {
	return xmlListHelper.Remove(l, indexes...)
}

func (l xmlList) Set(i int, v interface{}) (xmlIList, error) {
	return xmlListHelper.SetIndex(l, i, v)
}

func (l xmlList) Union(values ...interface{}) xmlIList {
	return xmlListHelper.Add(l, false, values...).Unique()
}

func (l xmlList) Without(values ...interface{}) xmlIList {
	return xmlListHelper.Without(l, values...)
}

// Dictionary implementation of IDictionary for xmlDict
type Dictionary = xmlDict
type xmlIDict = collections.IDictionary
type xmlDict map[string]interface{}

func (d xmlDict) Add(key, v interface{}) xmlIDict     { return xmlDictHelper.Add(d, key, v) }
func (d xmlDict) AsMap() map[string]interface{}       { return (map[string]interface{})(d) }
func (d xmlDict) Clone(keys ...interface{}) xmlIDict  { return xmlDictHelper.Clone(d, keys) }
func (d xmlDict) Count() int                          { return len(d) }
func (d xmlDict) Create(args ...int) xmlIDict         { return xmlListHelper.CreateDictionary(args...) }
func (d xmlDict) CreateList(args ...int) xmlIList     { return xmlHelper.CreateList(args...) }
func (d xmlDict) Flush(keys ...interface{}) xmlIDict  { return xmlDictHelper.Flush(d, keys) }
func (d xmlDict) Get(keys ...interface{}) interface{} { return xmlDictHelper.Get(d, keys) }
func (d xmlDict) GetKeys() xmlIList                   { return xmlDictHelper.GetKeys(d) }
func (d xmlDict) GetKinds() xmlIDict                  { return xmlDictHelper.GetTypes(d, true) }
func (d xmlDict) GetTypes() xmlIDict                  { return xmlDictHelper.GetTypes(d, false) }
func (d xmlDict) GetValues() xmlIList                 { return xmlDictHelper.GetValues(d) }
func (d xmlDict) Has(keys ...interface{}) bool        { return xmlDictHelper.Has(d, keys) }
func (d xmlDict) KeysAsString() strArray              { return xmlDictHelper.KeysAsString(d) }
func (d xmlDict) Len() int                            { return len(d) }
func (d xmlDict) Native() interface{}                 { return must(collections.MarshalGo(d)) }
func (d xmlDict) Pop(keys ...interface{}) interface{} { return xmlDictHelper.Pop(d, keys) }
func (d xmlDict) Set(key, v interface{}) xmlIDict     { return xmlDictHelper.Set(d, key, v) }
func (d xmlDict) Transpose() xmlIDict                 { return xmlDictHelper.Transpose(d) }
func (d xmlDict) Type() str                           { return xmlDictHelper.Type(d) }
func (d xmlDict) TypeName() str                       { return str(xmlLower) }

func (d xmlDict) GetHelpers() (collections.IDictionaryHelper, collections.IListHelper) {
	return xmlDictHelper, xmlListHelper
}

func (d xmlDict) Default(key, defVal interface{}) interface{} {
	return xmlDictHelper.Default(d, key, defVal)
}

func (d xmlDict) Delete(key interface{}, otherKeys ...interface{}) (xmlIDict, error) {
	return xmlDictHelper.Delete(d, append([]interface{}{key}, otherKeys...))
}

func (d xmlDict) Merge(dict xmlIDict, otherDicts ...xmlIDict) xmlIDict {
	return xmlDictHelper.Merge(d, append([]xmlIDict{dict}, otherDicts...))
}

func (d xmlDict) Omit(key interface{}, otherKeys ...interface{}) xmlIDict {
	return xmlDictHelper.Omit(d, append([]interface{}{key}, otherKeys...))
}

// Generic helpers to simplify physical implementation
func xmlListConvert(list xmlIList) xmlIList { return xmlList(list.AsArray()) }
func xmlDictConvert(dict xmlIDict) xmlIDict { return xmlDict(dict.AsMap()) }
func needConversion(object interface{}, strict bool) bool {
	return needConversionImpl(object, strict, "Xml")
}

var xmlHelper = helperBase{ConvertList: xmlListConvert, ConvertDict: xmlDictConvert, NeedConversion: needConversion}
var xmlListHelper = helperList{BaseHelper: xmlHelper}
var xmlDictHelper = helperDict{BaseHelper: xmlHelper}

// DictionaryHelper gives public access to the basic dictionary functions
var DictionaryHelper collections.IDictionaryHelper = xmlDictHelper

// GenericListHelper gives public access to the basic list functions
var GenericListHelper collections.IListHelper = xmlListHelper

type (
	str      = collections.String
	strArray = collections.StringArray
)

var must = errors.Must
