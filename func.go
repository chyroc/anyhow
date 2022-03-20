package anyhow

// And Returns res if the result is Ok, otherwise returns the Err value of self.
func And[T, U any](r Result[T], t Result[U]) Result[U] {
	if r.E != nil {
		return Err[U](r.E)
	}
	if t.E != nil {
		return Err[U](t.E)
	}
	return Ok[U](t.T)
}

// AndThen Calls op if the result is Ok, otherwise returns the Err value of self.
//
// This function can be used for control flow based on Result values.
func AndThen[T, U any](r Result[T], op func(T) Result[U]) Result[U] {
	if r.E != nil {
		return Result[U]{E: r.E}
	}
	return op(r.T)
}

// Map maps a Result<T, E> to Result<U, E> by applying a function to a contained Ok value,
// leaving an Err value untouched.
//
// This function can be used to compose the results of two functions.
func Map[T, U any](r Result[T], op func(T) U) Result[U] {
	if r.E != nil {
		return Result[U]{E: r.E}
	}
	return Result[U]{T: op(r.T)}
}

// map_err

// MapOr Returns the provided default (if Err), or applies a function to the contained value (if Ok),
//
// Arguments passed to map_or are eagerly evaluated; if you are passing the result of a function call,
// it is recommended to use map_or_else, which is lazily evaluated.
func MapOr[T, U any](r Result[T], defaultValue U, f func(T) U) U {
	if r.E != nil {
		return defaultValue
	}
	return f(r.T)
}

// MapOrElse Maps a Result<T, E> to U by applying fallback function default to a contained Err value,
// or function f to a contained Ok value.
//
// This function can be used to unpack a successful result while handling an error.
func MapOrElse[T, U any](r Result[T], fallback func(err error) U, f func(T) U) U {
	if r.E != nil {
		return fallback(r.E)
	}
	return f(r.T)
}
