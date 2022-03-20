package anyhow

import (
	"fmt"
)

// And Returns res if the result is Ok, otherwise returns the Err value of self.
// func And[T, U any](r Result[T], t Result[U]) Result[U] {
// 	if r.E != nil {
// 		return Err[U](r.E)
// 	}
// 	if t.E != nil {
// 		return Err[U](t.E)
// 	}
// 	return Ok[U](t.T)
// }

// AndThen Calls op if the result is Ok, otherwise returns the Err value of self.
//
// This function can be used for control flow based on Result values.
// func AndThen[T, U any](r Result[T], op func(T) Result[U]) Result[U] {
// 	if r.E != nil {
// 		return Result[U]{E: r.E}
// 	}
// 	return op(r.T)
// }

// as_deref

// as_deref_mut

// as_mut

// as_ref

// cloned

// contains

// contains_err

// copied

// Err Converts from Result<T, E> to Option<E>.
//
// Converts self into an Option<E>, consuming self, and discarding the success value, if any.
func (r Result[T]) Err() error {
	return r.E
}

// Expect Returns the contained Ok value, consuming the self value.
//
// Panics: Panics if the value is an Err, with a panic message including the passed message, and the content of the Err.
func (r Result[T]) Expect(msg string) T {
	if r.E != nil {
		panic(fmt.Sprintf("%s: %s", msg, r.E))
	}
	return r.T
}

func (r Result[T]) ExpectErr(msg string) error {
	if r.E == nil {
		panic(fmt.Sprintf("%s: %v", msg, r.T))
	}
	return r.E
}

// Flatten

// Inspect Calls the provided closure with a reference to the contained value (if Ok).
func (r Result[T]) Inspect(f func(T)) Result[T] {
	if r.E == nil {
		f(r.T)
	}
	return r
}

// InspectErr Calls the provided closure with a reference to the contained error (if Err).
func (r Result[T]) InspectErr(f func(err error)) Result[T] {
	if r.E != nil {
		f(r.E)
	}
	return r
}

// IntoErr Returns the contained Err value, but never panics.
//
// Unlike unwrap_err, this method is known to never panic on the result types it is implemented for.
// Therefore, it can be used instead of unwrap_err as a maintainability safeguard that will fail to
// compile if the ok type of the Result is later changed to a type that can actually occur.
func (r Result[T]) IntoErr() error {
	return r.E
}

// IntoOk Returns the contained Ok value, but never panics.
//
// Unlike unwrap, this method is known to never panic on the result types it is implemented for.
// Therefore, it can be used instead of unwrap as a maintainability safeguard that will fail to
// compile if the error type of the Result is later changed to an error that can actually occur.
func (r Result[T]) IntoOk() T {
	return r.T
}

// into_ok_or_err

// IsErr Returns true if the result is Err.
func (r Result[T]) IsErr() bool {
	return r.E != nil
}

// IsOk Returns true if the result is Ok.
func (r Result[T]) IsOk() bool {
	return r.E == nil
}

// iter

// iter_mut

// Map maps a Result<T, E> to Result<U, E> by applying a function to a contained Ok value,
// leaving an Err value untouched.
//
// This function can be used to compose the results of two functions.
// func Map[T, U any](r Result[T], op func(T) U) Result[U] {
// 	if r.E != nil {
// 		return Result[U]{E: r.E}
// 	}
// 	return Result[U]{T: op(r.T)}
// }

// map_err

// MapOr Returns the provided default (if Err), or applies a function to the contained value (if Ok),
//
// Arguments passed to map_or are eagerly evaluated; if you are passing the result of a function call,
// it is recommended to use map_or_else, which is lazily evaluated.
// func MapOr[T, U any](r Result[T], defaultValue U, f func(T) U) U {
// 	if r.E != nil {
// 		return defaultValue
// 	}
// 	return f(r.T)
// }

// MapOrElse Maps a Result<T, E> to U by applying fallback function default to a contained Err value,
// or function f to a contained Ok value.
//
// This function can be used to unpack a successful result while handling an error.
// func MapOrElse[T, U any](r Result[T], fallback func(err error) U, f func(T) U) U {
// 	if r.E != nil {
// 		return fallback(r.E)
// 	}
// 	return f(r.T)
// }

// map_err

// Ok Converts from Result<T, E> to Option<T>.
//
// Converts self into an Option<T>, consuming self, and discarding the error, if any.
func (r Result[T]) Ok() *T {
	if r.E != nil {
		return nil
	}
	return &r.T
}

// Or Returns res if the result is Err, otherwise returns the Ok value of self.
//
// Arguments passed to or are eagerly evaluated; if you are passing the result of a function call,
// it is recommended to use or_else, which is lazily evaluated.
func (r Result[T]) Or(t Result[T]) Result[T] {
	if r.E == nil {
		return Ok[T](r.T)
	}
	if t.E == nil {
		return Ok[T](t.T)
	}
	return Err[T](t.E)
}

// or_else

// transpose

// Unwrap Returns the contained Ok value, consuming the self value.
//
// Because this function may panic, its use is generally discouraged.
// Instead, prefer to use pattern matching and handle the Err case explicitly,
// or call unwrap_or, unwrap_or_else, or unwrap_or_default.
//
// Panics: Panics if the value is an Err, with a panic message provided by the Err’s value.
func (r Result[T]) Unwrap() T {
	if r.E != nil {
		panic(r.E)
	}
	return r.T
}

// UnwrapErr Returns the contained Err value, consuming the self value.
//
// Panics: Panics if the value is an Ok, with a custom panic message provided by the Ok’s value.
func (r Result[T]) UnwrapErr() error {
	if r.E != nil {
		return r.E
	}
	panic(r.T)
}

// UnwrapErrUnchecked Returns the contained Err value, consuming the self value,
// without checking that the value is not an Ok.
func (r Result[T]) UnwrapErrUnchecked() error {
	return r.E
}

// UnwrapOr Returns the contained Ok value or a provided default.
//
// Arguments passed to unwrap_or are eagerly evaluated; if you are passing the result of a function call,
// it is recommended to use unwrap_or_else, which is lazily evaluated.
func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.E != nil {
		return defaultValue
	}
	return r.T
}

// unwrap_or_default

// UnwrapOrElse Returns the contained Ok value or computes it from a closure.
func (r Result[T]) UnwrapOrElse(op func(err error) T) T {
	if r.E != nil {
		return op(r.E)
	}
	return r.T
}

// UnwrapUnchecked Returns the contained Ok value, consuming the self value, without checking that the value is not an Err.
func (r Result[T]) UnwrapUnchecked() T {
	return r.T
}
