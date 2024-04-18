package anyhow

// Result1 ...
type Result1[T1 any] struct {
	v1  T1
	err error
}

// Result2 ...
type Result2[T1, T2 any] struct {
	v1  T1
	v2  T2
	err error
}

// Result3 ...
type Result3[T1, T2, T3 any] struct {
	v1  T1
	v2  T2
	v3  T3
	err error
}

// Result4 ...
type Result4[T1, T2, T3, T4 any] struct {
	v1  T1
	v2  T2
	v3  T3
	v4  T4
	err error
}

// Result5 ...
type Result5[T1, T2, T3, T4, T5 any] struct {
	v1  T1
	v2  T2
	v3  T3
	v4  T4
	v5  T5
	err error
}

// Result6 ...
type Result6[T1, T2, T3, T4, T5, T6 any] struct {
	v1  T1
	v2  T2
	v3  T3
	v4  T4
	v5  T5
	v6  T6
	err error
}
