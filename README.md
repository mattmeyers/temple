# temple

`temple` is a library of Go template functions and a CLI tool for compiling these templates.

## CLI

```sh
Usage: temple [Options] template1 template2...

Compile Go templates from the command line

  -d string
        a JSON file containing the template data
  -html
        use html/template for template parsing
  -o string
        the output filename
```

### Usage

Given the template file `report.tmpl`

```html
<ul>
  {{- range .Prices}}
  <li>{{. | ToString 2 | Commas}}</li>
  {{- end}}
</ul>
```

and the data file `data.json`

```json
{
  "Prices": [123, 1234567.56, 0.56]
}
```

Running the command

```sh
temple --html -d data.json -o report.html report.tmpl
```

will generate the file `report.html` containing

```html
<ul>
  <li>123.00</li>
  <li>1,234,567.56</li>
  <li>0.56</li>
</ul>
```
