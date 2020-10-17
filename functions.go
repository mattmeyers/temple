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

// Clear empties the FuncMap of all of its values.
func (f *FuncMap) Clear() { *f = make(FuncMap) }

// MergeFuncMaps combines multiple FuncMap structures. If the same
// function name is found in multiple FuncMaps, then the last
// occurence will appear in the resulting FuncMap.
func MergeFuncMaps(base FuncMap, additions ...FuncMap) FuncMap {
	if base == nil {
		base = make(FuncMap)
	}

	for _, a := range additions {
		if a == nil {
			continue
		}

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
	"Commas":     Commas,
	"IsNumeric":  IsNumeric,
	"Join":       Join,
	"FormatMask": FormatMask,
}

// NumbersFuncs maps all number and math related functions
// provided by temple.
var NumbersFuncs FuncMap = FuncMap{
	"Max":      Max,
	"IntMax":   IntMax,
	"FloatMax": FloatMax,
	"IntMin":   IntMin,
	"FloatMin": FloatMin,
	"Ceil":     Ceil,
	"Floor":    Floor,
	"Mod":      Mod,
	"Sum":      Sum,
	"Diff":     Diff,
	"Mul":      Mul,
	"Div":      Div,
}

// ConversionFuncs maps all type conversion related functions
// provided by temple.
var ConversionFuncs FuncMap = FuncMap{
	"ToInt":     ToInt,
	"ToFloat64": ToFloat64,
	"ToString":  ToString,

	"ToIntSlice":     ToIntSlice,
	"ToFloat64Slice": ToFloat64Slice,
	"ToStringSlice":  ToStringSlice,
}

// CollectionFuncs maps all type conversion related functions
// provided by temple.
var CollectionFuncs FuncMap = FuncMap{
	"NewList":  NewList,
	"NewSet":   NewSet,
	"Contains": Contains,
}
