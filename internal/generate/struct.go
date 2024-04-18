//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"text/template"
)

var tmplStruct = `
// Result{{len .}} ...
type Result{{len .}}[{{range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }} any] struct {
    {{- range $val := .}}
	v{{$val}}  T{{$val}}
    {{- end }}
	err error
}
`

func generateStruct(n int) string {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "package anyhow\n")

	var ns []int
	t := template.Must(template.New("").Parse(tmplStruct))
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
