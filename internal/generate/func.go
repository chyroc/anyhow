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
// And{{len .L}}{{len .R}} Returns res if the result is Ok, otherwise returns the Err value of self
func And{{len .L}}{{len .R}}[{{range $val := .L}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}, {{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }} any](self Result{{len .L}}[{{range $val := .L}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}], res Result{{len .R}}[{{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}]) Result{{len .R}}[{{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}] {
	if self.IsOk() {
		return res
	}
	return Err{{len .R}}[{{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}](self.err)
}

// AndThen{{len .L}}{{len .R}} Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen{{len .L}}{{len .R}}[{{range $val := .L}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}, {{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }} any](self Result{{len .L}}[{{range $val := .L}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}], op func({{range $val := .L}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}) Result{{len .R}}[{{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}]) Result{{len .R}}[{{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}] {
	if self.IsOk() {
		return op(self.IntoOk())	
	}
	return Err{{len .R}}[{{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}](self.err)
}

// Map{{len .L}}{{len .R}} Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map{{len .L}}{{len .R}}[{{range $val := .L}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}, {{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }} any](self Result{{len .L}}[{{range $val := .L}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}], op func({{range $val := .L}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}) ({{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }})) Result{{len .R}}[{{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}] {
	if self.err != nil {
		return Err{{len .R}}[{{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}](self.err)
	}
	return Ok{{len .R}}[{{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}](op(self.IntoOk()))
}

// MapOr{{len .L}}{{len .R}} Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr{{len .L}}{{len .R}}[{{range $val := .L}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}, {{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }} any](self Result{{len .L}}[{{range $val := .L}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}], {{range $val := .R}}{{if gt $val 1 }}, {{end}}defaultV{{$val}} U{{$val}}{{ end }}, op func({{range $val := .L}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}) Result{{len .R}}[{{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}]) Result{{len .R}}[{{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}] {
	if self.err != nil {
		return Ok{{len .R}}[{{range $val := .R}}{{if gt $val 1 }}, {{end}}U{{$val}}{{ end }}]({{range $val := .R}}{{if gt $val 1 }}, {{end}}defaultV{{$val}}{{ end }})
	}
	return op(self.IntoOk())
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
