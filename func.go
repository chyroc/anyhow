package anyhow

//
//// AndThen Calls op if the result is Ok, otherwise returns the Err value of self.
////
//// This function can be used for control flow based on Result values.
////
//// 将 r 作为参数传递到 op 中
//
//// Or Returns the provided default (if Err), or applies a function to the contained value (if Ok),
////
//// Arguments passed to map_or are eagerly evaluated; if you are passing the result of a function call,
//// it is recommended to use map_or_else, which is lazily evaluated.
//func Or[T, U any](r Result[T], defaultValue U, f func(T) U) U {
//	if r.E != nil {
//		return defaultValue
//	}
//	return f(r.T)
//}
//
//// OrElse Maps a Result<T, E> to U by applying fallback function default to a contained Err value,
//// or function f to a contained Ok value.
////
//// This function can be used to unpack a successful result while handling an error.
//func OrElse[T, U any](r Result[T], fallback func(err error) U, f func(T) U) U {
//	if r.E != nil {
//		return fallback(r.E)
//	}
//	return f(r.T)
//}
