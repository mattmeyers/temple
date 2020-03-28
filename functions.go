package temple

import (
	htmltmpl "html/template"
	texttmpl "text/template"
)

// FuncMap maps template function names to functions. This
// type mirrors the types with the same name in the
// text/template and html/template packages. The Text()
// and HTML() methods can be used respectively to achieve
// these other FuncMap types.
type FuncMap map[string]interface{}

// Text casts a temple.FuncMap to a text/template.FuncMap.
func (f FuncMap) Text() texttmpl.FuncMap {
	return texttmpl.FuncMap(f)
}

// HTML casts a temple.FuncMap to a html/template.FuncMap.
func (f FuncMap) HTML() htmltmpl.FuncMap {
	return htmltmpl.FuncMap(f)
}

// MergeFuncMaps combines multiple FuncMap structures. If the same
// function name is found in multiple FuncMaps, then the last
// occurence will appear in the resulting FuncMap.
func MergeFuncMaps(base FuncMap, additions ...FuncMap) FuncMap {
	for _, a := range additions {
		for k, v := range a {
			base[k] = v
		}
	}
	return base
}

// FullFuncMap merges all standard FuncMaps.
func FullFuncMap() FuncMap {
	return MergeFuncMaps(
		StringsFuncs,
		NumbersFuncs,
		ConversionFuncs,
		CollectionFuncs,
	)
}

// StringsFuncs maps all string related functions provided
// by temple.
var StringsFuncs FuncMap = FuncMap{
	"ToString":  ToString,
	"Commas":    Commas,
	"IsNumeric": IsNumeric,
	"Join":      Join,
}

// NumbersFuncs maps all number and math related functions
// provided by temple.
var NumbersFuncs FuncMap = FuncMap{
	"Max":   Max,
	"Min":   Min,
	"Ceil":  Ceil,
	"Floor": Floor,
	"Mod":   Mod,
	"Sum":   Sum,
	"Diff":  Diff,
	"Mul":   Mul,
	"Div":   Div,
}

// ConversionFuncs maps all type conversion related functions
// provided by temple.
var ConversionFuncs FuncMap = FuncMap{
	"ToInt":     ToInt,
	"ToInt8":    ToInt8,
	"ToInt16":   ToInt16,
	"ToInt32":   ToInt32,
	"ToInt64":   ToInt64,
	"ToUint":    ToUint,
	"Touint8":   Touint8,
	"ToUint16":  ToUint16,
	"ToUint32":  ToUint32,
	"ToUint64":  ToUint64,
	"ToFloat32": ToFloat32,
	"ToFloat64": ToFloat64,

	"ToIntSlice":     ToIntSlice,
	"ToInt8Slice":    ToInt8Slice,
	"ToInt16Slice":   ToInt16Slice,
	"ToInt32Slice":   ToInt32Slice,
	"ToInt64Slice":   ToInt64Slice,
	"ToUintSlice":    ToUintSlice,
	"Touint8Slice":   Touint8Slice,
	"ToUint16Slice":  ToUint16Slice,
	"ToUint32Slice":  ToUint32Slice,
	"ToUint64Slice":  ToUint64Slice,
	"ToFloat32Slice": ToFloat32Slice,
	"ToFloat64Slice": ToFloat64Slice,
}

// CollectionFuncs maps all type conversion related functions
// provided by temple.
var CollectionFuncs FuncMap = FuncMap{
	"NewSlice": NewSlice,
	"NewSet":   NewSet,
	"Contains": Contains,
}
