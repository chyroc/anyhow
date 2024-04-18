//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"text/template"
)

var tmplMethod = `
// IntoErr Returns the contained Err value, but never panics
func (r Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) IntoErr() error {
	return r.err
}

// IntoOk Returns the contained Ok value, but never panics
func (r Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) IntoOk() ({{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}) {
	return {{ range $val := .}}{{if gt $val 1 }}, {{end}}r.v{{$val}}{{ end }}
}

// Expect Returns the contained Ok value, otherwise panics with a panic message including the passed message, and the content of the Err
func (r Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) Expect(msg string) ({{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}) {
	if r.IsOk() {
		return {{ range $val := .}}{{if gt $val 1 }}, {{end}}r.v{{$val}}{{ end }}
	}
	panic(fmt.Errorf("%s: %v", msg, r.err))
} 

// ExpectErr Returns the contained Err value, otherwise panics with a panic message including the passed message, and the content of the Ok
func (r Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) ExpectErr(msg string) error {
	if r.IsErr() {
		return r.err
	}
	panic(fmt.Errorf("%s: {{ range $val := .}}{{if gt $val 1 }}, {{end}}%v{{ end }}", msg, {{ range $val := .}}{{if gt $val 1 }}, {{end}}r.v{{$val}}{{ end }})) 
}

// Inspect calls a function with a reference to the contained value if Ok
func (r Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) Inspect(f func({{range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }})) Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}] {
	if r.IsOk() {
		f(r.IntoOk())
	}
	return r
}

// InspectErr calls a function with a reference to the contained value if Err
func (r Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) InspectErr(f func(error)) Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}] {
	if r.IsErr() {
		f(r.err)
	}
	return r
}

// MapErr Maps a R[T] to R[T] by applying a function to a contained Err value, leaving an Ok value untouched
func (r Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) MapErr(op func(error) error) Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}] {
	if r.IsErr() {
		return Err{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}](op(r.err))
	}
	return r
}

// IsErr returns true if the result is Err
func (r Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) IsErr() bool {
	return r.err != nil
}

// IsOk returns true if the result is Ok
func (r Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) IsOk() bool {
	return r.err == nil
}

// Or Returns res if the result is Err, otherwise returns the Ok value of self
func (r Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) Or(res Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) Result{{len .}}[{{ range $val :=.}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}] {
	if r.IsErr() {
		return res
	}
	return r
}

// OrElse Calls op if the result is Err, otherwise returns the Ok value of self
func (r Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) OrElse(op func(error) Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}] {
	if r.IsErr() {
		return op(r.err)
	}
	return r
}

// Unwrap Returns the contained Ok value, otherwise panics with Err
func (r Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) Unwrap() ({{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}) {
	if r.err != nil {
		panic(r.err)
	}
	return {{ range $val := .}}{{if gt $val 1 }}, {{end}}r.v{{$val}}{{ end }}
}

// UnwrapOr Returns the contained Ok value or a provided default
func (r Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) UnwrapOr({{range $val := .}}{{if gt $val 1 }}, {{end}}v{{$val}} T{{$val}}{{ end }}) ({{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}) {
	if r.err != nil {
		return {{range $val := .}}{{if gt $val 1 }}, {{end}}v{{$val}}{{ end }}
	}
	return {{ range $val := .}}{{if gt $val 1 }}, {{end}}r.v{{$val}}{{ end }}
}

// UnwrapOrDefault Returns the contained Ok value or a default value
func (r Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) UnwrapOrDefault() ({{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}) {
	if r.err != nil {
		return {{ range $val := .}}{{if gt $val 1 }}, {{end}}zero[T{{$val}}](){{ end }}
	}
	return {{ range $val := .}}{{if gt $val 1 }}, {{end}}r.v{{$val}}{{ end }}
}

// UnwrapOrElse Returns the contained Ok value or computes it from a closure
func (r Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) UnwrapOrElse(op func(error) ({{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }})) ({{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}) {
	if r.err != nil {
		return op(r.err)
	}
	return {{ range $val := .}}{{if gt $val 1 }}, {{end}}r.v{{$val}}{{ end }}
}

// Unpack return (v1, v2, ..., err)
func (r Result{{len .}}[{{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) Unpack() ({{ range $val := .}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}, error) {
	if r.err != nil {
		panic(r.err)
	}
	return {{ range $val := .}}{{if gt $val 1 }}, {{end}}r.v{{$val}}{{ end }}, r.err
}

{{- $ns := .}}
{{ range $val := $ns}}
func (r Result{{len $ns}}[{{ range $val := $ns}}{{if gt $val 1 }}, {{end}}T{{$val}}{{ end }}]) V{{$val}}() T{{$val}} {
	return r.v{{$val}}
}
{{ end }}
`

func generateMethod(n int) string {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "package anyhow\n\nimport \"fmt\"\n\n")

	var ns []int
	t := template.Must(template.New("").Parse(tmplMethod))
	for i := 1; i <= n; i++ {
		ns = append(ns, i)
		err := t.Execute(&buf, ns)
		if err != nil {
			log.Fatal(err)
		}
	}
	return string(buf.String())

	b, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
