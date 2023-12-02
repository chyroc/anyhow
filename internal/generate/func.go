//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"text/template"
)

var tmplFunc = `
// Then{{len .L}}{{len .R}} ...
func Then{{len .L}}{{len .R}}[{{range $val := .L}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}, {{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }} any](r Result{{len .L}}[{{range $val := .L}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}], op func({{range $val := .L}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}) Result{{len .R}}[{{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}]) Result{{len .R}}[{{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}] {
	if r.err != nil {
		return Err{{len .R}}[{{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}](r.err)
	}
	return op(r.Value())
}

// Or{{len .L}}{{len .R}} ...
func Or{{len .L}}{{len .R}}[{{range $val := .L}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}, {{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }} any](r Result{{len .L}}[{{range $val := .L}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}], {{range $val := .R}}{{if gt $val 1 }}, {{end}}defaultV{{$val}} U{{$val}}{{ end }}, op func({{range $val := .L}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}) Result{{len .R}}[{{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}]) Result{{len .R}}[{{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}] {
	if r.err != nil {
		return Ok{{len .R}}[{{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}]({{range $val := .R}}{{if gt $val 1 }}, {{end}}defaultV{{$val}}{{ end }})
	}
	return op(r.Value())
}

`

func generateFunc(n int) string {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "package anyhow\n")

	var left []int
	t := template.Must(template.New("").Parse(tmplFunc))
	for i := 1; i <= n; i++ {
		left = append(left, i)
		right := []int{}
		for j := 1; j <= n; j++ {
			right = append(right, j)
			if err := t.Execute(&buf, map[string]any{
				"L": left,
				"R": right,
			}); err != nil {
				log.Fatal(err)
			}
		}
	}

	return string(buf.Bytes())

	b, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
