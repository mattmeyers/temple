package temple

import (
	htmltmpl "html/template"
	texttmpl "text/template"
)

var TextFuncMap texttmpl.FuncMap = texttmpl.FuncMap(HTMLFuncMap)

var HTMLFuncMap htmltmpl.FuncMap = htmltmpl.FuncMap{
	"ToString":  ToString,
	"Commas":    Commas,
	"IsNumeric": IsNumeric,
}
