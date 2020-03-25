package temple

import (
	htmltmpl "html/template"
	texttmpl "text/template"
)

type FuncMap map[string]interface{}

func (f FuncMap) Text() texttmpl.FuncMap {
	return texttmpl.FuncMap(f)
}

func (f FuncMap) HTML() htmltmpl.FuncMap {
	return htmltmpl.FuncMap(f)
}

func MergeFuncMaps(base FuncMap, adds ...FuncMap) FuncMap {
	for _, a := range adds {
		for k, v := range a {
			base[k] = v
		}
	}
	return base
}

func FullFuncMap() FuncMap {
	return MergeFuncMaps(
		StringsFuncs,
		NumbersFuncs,
		ConversionFuncs,
	)
}

var StringsFuncs FuncMap = FuncMap{
	"ToString":  ToString,
	"Commas":    Commas,
	"IsNumeric": IsNumeric,
}

var NumbersFuncs FuncMap = FuncMap{
	// Float64 functions
	"Max":      Max,
	"Min":      Min,
	"SliceMax": SliceMax,
	"SliceMin": SliceMin,
	// int functions
	"IntMax":      IntMax,
	"IntMin":      IntMin,
	"IntSliceMax": IntSliceMax,
	"IntSliceMin": IntSliceMin,
	// uint functions
	"UintMax":      UintMax,
	"UintMin":      UintMin,
	"UintSliceMax": UintSliceMax,
	"UintSliceMin": UintSliceMin,
}

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
