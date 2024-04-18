//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"text/template"
)

var tmplConstructor = `
// Ok{{len .}} pack v1, v2, ... to Result{{len .}}
func Ok{{len .}}[{{range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }} any]({{range $val := .}}{{if gt $val 1 }}, {{end}}v{{$val}} T{{$val}}{{ end }}) Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}] {
    return Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]{{"{"}}{{range $val := .}}{{if gt $val 1 }}, {{end}}v{{$val}}: v{{$val}}{{ end }}, err: nil}
}

// Err{{len .}} pack err to Result{{len .}}
func Err{{len .}}[{{range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }} any](err error) Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}] {
	return Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]{err: err{{"}"}}
}

// Of{{len .}} pack v1, v2, ..., err to Result{{len .}}
func Of{{len .}}[{{range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }} any]({{range $val := .}}{{if gt $val 1 }}, {{end}}v{{$val}} T{{$val}}{{ end }}, err error) Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}] {
    return Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]{{"{"}}{{range $val := .}}{{if gt $val 1 }}, {{end}}v{{$val}}: v{{$val}}{{ end }}, err: err}
}
`

func generateConstructor(n int) string {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "package anyhow\n")

	var ns []int
	t := template.Must(template.New("").Parse(tmplConstructor))
	for i := 1; i <= n; i++ {
		ns = append(ns, i)
		err := t.Execute(&buf, ns)
		if err != nil {
			log.Fatal(err)
		}
	}
	// return string(buf.String())

	b, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
