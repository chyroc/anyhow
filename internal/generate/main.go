//go:build ignore

package main

import (
	"os"
)

var tmpl = `
// Result{{len .}} ...
type Result{{len .}}[{{range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }} any] struct {
    {{- range $val := .}}
	v{{$val}}  T{{$val}}
    {{- end }}
	err error
}

// Unwrap return data or panic(err)
func (r *Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) Unwrap() ({{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}) {
	if r.err != nil {
		panic(r.err)
	}
	return {{ range $val := .}}{{if gt $val 1 }}, {{end}}r.v{{$val}}{{ end }}
}

// Unpack return (v1, v2, ..., err)
func (r *Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) Unpack() ({{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}, error) {
	if r.err != nil {
		panic(r.err)
	}
	return {{ range $val := .}}{{if gt $val 1 }}, {{end}}r.v{{$val}}{{ end }}, r.err
}

// Err return err
func (r *Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) Err() error {
	return r.err
}

// Value return v1, v2, ...
func (r *Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) Value() ({{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}) {
	return {{ range $val := .}}{{if gt $val 1 }}, {{end}}r.v{{$val}}{{ end }}
}

// IsErr check if err is not nil
func (r *Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) IsErr() bool {
	return r.err != nil
}

// IsOk check if err is nil
func (r *Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) IsOk() bool {
	return r.err == nil
}

// Ok{{len .}} pack v1, v2, ... to Result{{len .}}
func Ok{{len .}}[{{range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }} any]({{range $val := .}}{{if gt $val 1 }}, {{end}}v{{$val}} T{{$val}}{{ end }}) Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}] {
    return Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]{{"{"}}{{range $val := .}}{{if gt $val 1 }}, {{end}}v{{$val}}: v{{$val}}{{ end }}, err: nil}
}

// Err{{len .}} pack err to Result{{len .}}
func Err{{len .}}[{{range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }} any](err error) Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}] {
	return Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]{err: err{{"}"}}
}

// Then{{len .}} ...
func Then{{len .}}[{{range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}, {{range $val := .}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }} any](r Result{{len .}}[{{range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}], op func({{range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}) Result{{len .}}[{{range $val := .}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}]) Result{{len .}}[{{range $val := .}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}] {
	if r.err != nil {
		return Err{{len .}}[{{range $val := .}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}](r.err)
	}
	return op({{range $val := .}}{{if gt $val 1 }}, {{end}}r.V{{$val}}(){{ end }})
}


{{- $ns := .}}
{{ range $val := $ns}}
func (r *Result{{len $ns}}[{{ range $val := $ns}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) V{{$val}}() T{{$val}} {
	return r.v{{$val}}
}
{{ end }}
`

func main() {
	{
		f, err := os.Create("struct.go")
		assert(err)
		f.Write([]byte(generateStruct(6)))
		f.Close()
	}
	{
		f, err := os.Create("func.go")
		assert(err)
		f.Write([]byte(generateFunc(6)))
		f.Close()
	}
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}
