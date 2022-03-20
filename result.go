package anyhow

// Result is the type used for returning and propagating errors.
// It is an enum with the variants, Ok(T), representing success and containing a value,
// and Err(E), representing error and containing an error value.
//
// docs: https://doc.rust-lang.org/std/result/
type Result[T any] struct {
	T T
	E error
}

// Err equal to generate a Result[T]{E: err}
func Err[T any](err error) Result[T] {
	return Result[T]{E: err}
}

// Ok equal to generate a Result[T]{T: t}
func Ok[T any](t T) Result[T] {
	return Result[T]{T: t}
}
