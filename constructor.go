package anyhow

// Ok1 pack v1, v2, ... to Result1
func Ok1[T1 any](v1 T1) Result1[T1] {
	return Result1[T1]{v1: v1, err: nil}
}

// Err1 pack err to Result1
func Err1[T1 any](err error) Result1[T1] {
	return Result1[T1]{err: err}
}

// Of1 pack v1, v2, ..., err to Result1
func Of1[T1 any](v1 T1, err error) Result1[T1] {
	return Result1[T1]{v1: v1, err: err}
}

// Ok2 pack v1, v2, ... to Result2
func Ok2[T1, T2 any](v1 T1, v2 T2) Result2[T1, T2] {
	return Result2[T1, T2]{v1: v1, v2: v2, err: nil}
}

// Err2 pack err to Result2
func Err2[T1, T2 any](err error) Result2[T1, T2] {
	return Result2[T1, T2]{err: err}
}

// Of2 pack v1, v2, ..., err to Result2
func Of2[T1, T2 any](v1 T1, v2 T2, err error) Result2[T1, T2] {
	return Result2[T1, T2]{v1: v1, v2: v2, err: err}
}

// Ok3 pack v1, v2, ... to Result3
func Ok3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3) Result3[T1, T2, T3] {
	return Result3[T1, T2, T3]{v1: v1, v2: v2, v3: v3, err: nil}
}

// Err3 pack err to Result3
func Err3[T1, T2, T3 any](err error) Result3[T1, T2, T3] {
	return Result3[T1, T2, T3]{err: err}
}

// Of3 pack v1, v2, ..., err to Result3
func Of3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3, err error) Result3[T1, T2, T3] {
	return Result3[T1, T2, T3]{v1: v1, v2: v2, v3: v3, err: err}
}

// Ok4 pack v1, v2, ... to Result4
func Ok4[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4) Result4[T1, T2, T3, T4] {
	return Result4[T1, T2, T3, T4]{v1: v1, v2: v2, v3: v3, v4: v4, err: nil}
}

// Err4 pack err to Result4
func Err4[T1, T2, T3, T4 any](err error) Result4[T1, T2, T3, T4] {
	return Result4[T1, T2, T3, T4]{err: err}
}

// Of4 pack v1, v2, ..., err to Result4
func Of4[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4, err error) Result4[T1, T2, T3, T4] {
	return Result4[T1, T2, T3, T4]{v1: v1, v2: v2, v3: v3, v4: v4, err: err}
}

// Ok5 pack v1, v2, ... to Result5
func Ok5[T1, T2, T3, T4, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) Result5[T1, T2, T3, T4, T5] {
	return Result5[T1, T2, T3, T4, T5]{v1: v1, v2: v2, v3: v3, v4: v4, v5: v5, err: nil}
}

// Err5 pack err to Result5
func Err5[T1, T2, T3, T4, T5 any](err error) Result5[T1, T2, T3, T4, T5] {
	return Result5[T1, T2, T3, T4, T5]{err: err}
}

// Of5 pack v1, v2, ..., err to Result5
func Of5[T1, T2, T3, T4, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, err error) Result5[T1, T2, T3, T4, T5] {
	return Result5[T1, T2, T3, T4, T5]{v1: v1, v2: v2, v3: v3, v4: v4, v5: v5, err: err}
}

// Ok6 pack v1, v2, ... to Result6
func Ok6[T1, T2, T3, T4, T5, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) Result6[T1, T2, T3, T4, T5, T6] {
	return Result6[T1, T2, T3, T4, T5, T6]{v1: v1, v2: v2, v3: v3, v4: v4, v5: v5, v6: v6, err: nil}
}

// Err6 pack err to Result6
func Err6[T1, T2, T3, T4, T5, T6 any](err error) Result6[T1, T2, T3, T4, T5, T6] {
	return Result6[T1, T2, T3, T4, T5, T6]{err: err}
}

// Of6 pack v1, v2, ..., err to Result6
func Of6[T1, T2, T3, T4, T5, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, err error) Result6[T1, T2, T3, T4, T5, T6] {
	return Result6[T1, T2, T3, T4, T5, T6]{v1: v1, v2: v2, v3: v3, v4: v4, v5: v5, v6: v6, err: err}
}
