package anyhow

// Result1 ...
type Result1[T1 any] struct {
	v1  T1
	err error
}

// Unwrap return data or panic(err)
func (r Result1[T1]) Unwrap() T1 {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1
}

// UnwrapOr return v1, v2, ...
func (r Result1[T1]) UnwrapOr(v1 T1) T1 {
	if r.err != nil {
		return v1
	}
	return r.v1
}

// Unpack return (v1, v2, ..., err)
func (r Result1[T1]) Unpack() (T1, error) {
	return r.v1, r.err
}

// Err return err
func (r Result1[T1]) Err() error {
	return r.err
}

// Value return v1, v2, ...
func (r Result1[T1]) Value() T1 {
	return r.v1
}

// IsErr check if err is not nil
func (r Result1[T1]) IsErr() bool {
	return r.err != nil
}

// IsOk check if err is nil
func (r Result1[T1]) IsOk() bool {
	return r.err == nil
}

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

// Then1 ...
func Then1[T1, U1 any](r Result1[T1], op func(T1) Result1[U1]) Result1[U1] {
	if r.err != nil {
		return Err1[U1](r.err)
	}
	return op(r.V1())
}

func (r Result1[T1]) V1() T1 {
	return r.v1
}

// Result2 ...
type Result2[T1, T2 any] struct {
	v1  T1
	v2  T2
	err error
}

// Unwrap return data or panic(err)
func (r Result2[T1, T2]) Unwrap() (T1, T2) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2
}

// UnwrapOr return v1, v2, ...
func (r Result2[T1, T2]) UnwrapOr(v1 T1, v2 T2) (T1, T2) {
	if r.err != nil {
		return v1, v2
	}
	return r.v1, r.v2
}

// Unpack return (v1, v2, ..., err)
func (r Result2[T1, T2]) Unpack() (T1, T2, error) {
	return r.v1, r.v2, r.err
}

// Err return err
func (r Result2[T1, T2]) Err() error {
	return r.err
}

// Value return v1, v2, ...
func (r Result2[T1, T2]) Value() (T1, T2) {
	return r.v1, r.v2
}

// IsErr check if err is not nil
func (r Result2[T1, T2]) IsErr() bool {
	return r.err != nil
}

// IsOk check if err is nil
func (r Result2[T1, T2]) IsOk() bool {
	return r.err == nil
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

// Then2 ...
func Then2[T1, T2, U1, U2 any](r Result2[T1, T2], op func(T1, T2) Result2[U1, U2]) Result2[U1, U2] {
	if r.err != nil {
		return Err2[U1, U2](r.err)
	}
	return op(r.V1(), r.V2())
}

func (r Result2[T1, T2]) V1() T1 {
	return r.v1
}

func (r Result2[T1, T2]) V2() T2 {
	return r.v2
}

// Result3 ...
type Result3[T1, T2, T3 any] struct {
	v1  T1
	v2  T2
	v3  T3
	err error
}

// Unwrap return data or panic(err)
func (r Result3[T1, T2, T3]) Unwrap() (T1, T2, T3) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3
}

// UnwrapOr return v1, v2, ...
func (r Result3[T1, T2, T3]) UnwrapOr(v1 T1, v2 T2, v3 T3) (T1, T2, T3) {
	if r.err != nil {
		return v1, v2, v3
	}
	return r.v1, r.v2, r.v3
}

// Unpack return (v1, v2, ..., err)
func (r Result3[T1, T2, T3]) Unpack() (T1, T2, T3, error) {
	return r.v1, r.v2, r.v3, r.err
}

// Err return err
func (r Result3[T1, T2, T3]) Err() error {
	return r.err
}

// Value return v1, v2, ...
func (r Result3[T1, T2, T3]) Value() (T1, T2, T3) {
	return r.v1, r.v2, r.v3
}

// IsErr check if err is not nil
func (r Result3[T1, T2, T3]) IsErr() bool {
	return r.err != nil
}

// IsOk check if err is nil
func (r Result3[T1, T2, T3]) IsOk() bool {
	return r.err == nil
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

// Then3 ...
func Then3[T1, T2, T3, U1, U2, U3 any](r Result3[T1, T2, T3], op func(T1, T2, T3) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if r.err != nil {
		return Err3[U1, U2, U3](r.err)
	}
	return op(r.V1(), r.V2(), r.V3())
}

func (r Result3[T1, T2, T3]) V1() T1 {
	return r.v1
}

func (r Result3[T1, T2, T3]) V2() T2 {
	return r.v2
}

func (r Result3[T1, T2, T3]) V3() T3 {
	return r.v3
}

// Result4 ...
type Result4[T1, T2, T3, T4 any] struct {
	v1  T1
	v2  T2
	v3  T3
	v4  T4
	err error
}

// Unwrap return data or panic(err)
func (r Result4[T1, T2, T3, T4]) Unwrap() (T1, T2, T3, T4) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4
}

// UnwrapOr return v1, v2, ...
func (r Result4[T1, T2, T3, T4]) UnwrapOr(v1 T1, v2 T2, v3 T3, v4 T4) (T1, T2, T3, T4) {
	if r.err != nil {
		return v1, v2, v3, v4
	}
	return r.v1, r.v2, r.v3, r.v4
}

// Unpack return (v1, v2, ..., err)
func (r Result4[T1, T2, T3, T4]) Unpack() (T1, T2, T3, T4, error) {
	return r.v1, r.v2, r.v3, r.v4, r.err
}

// Err return err
func (r Result4[T1, T2, T3, T4]) Err() error {
	return r.err
}

// Value return v1, v2, ...
func (r Result4[T1, T2, T3, T4]) Value() (T1, T2, T3, T4) {
	return r.v1, r.v2, r.v3, r.v4
}

// IsErr check if err is not nil
func (r Result4[T1, T2, T3, T4]) IsErr() bool {
	return r.err != nil
}

// IsOk check if err is nil
func (r Result4[T1, T2, T3, T4]) IsOk() bool {
	return r.err == nil
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

// Then4 ...
func Then4[T1, T2, T3, T4, U1, U2, U3, U4 any](r Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if r.err != nil {
		return Err4[U1, U2, U3, U4](r.err)
	}
	return op(r.V1(), r.V2(), r.V3(), r.V4())
}

func (r Result4[T1, T2, T3, T4]) V1() T1 {
	return r.v1
}

func (r Result4[T1, T2, T3, T4]) V2() T2 {
	return r.v2
}

func (r Result4[T1, T2, T3, T4]) V3() T3 {
	return r.v3
}

func (r Result4[T1, T2, T3, T4]) V4() T4 {
	return r.v4
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

// Unwrap return data or panic(err)
func (r Result5[T1, T2, T3, T4, T5]) Unwrap() (T1, T2, T3, T4, T5) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5
}

// UnwrapOr return v1, v2, ...
func (r Result5[T1, T2, T3, T4, T5]) UnwrapOr(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) (T1, T2, T3, T4, T5) {
	if r.err != nil {
		return v1, v2, v3, v4, v5
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5
}

// Unpack return (v1, v2, ..., err)
func (r Result5[T1, T2, T3, T4, T5]) Unpack() (T1, T2, T3, T4, T5, error) {
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.err
}

// Err return err
func (r Result5[T1, T2, T3, T4, T5]) Err() error {
	return r.err
}

// Value return v1, v2, ...
func (r Result5[T1, T2, T3, T4, T5]) Value() (T1, T2, T3, T4, T5) {
	return r.v1, r.v2, r.v3, r.v4, r.v5
}

// IsErr check if err is not nil
func (r Result5[T1, T2, T3, T4, T5]) IsErr() bool {
	return r.err != nil
}

// IsOk check if err is nil
func (r Result5[T1, T2, T3, T4, T5]) IsOk() bool {
	return r.err == nil
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

// Then5 ...
func Then5[T1, T2, T3, T4, T5, U1, U2, U3, U4, U5 any](r Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if r.err != nil {
		return Err5[U1, U2, U3, U4, U5](r.err)
	}
	return op(r.V1(), r.V2(), r.V3(), r.V4(), r.V5())
}

func (r Result5[T1, T2, T3, T4, T5]) V1() T1 {
	return r.v1
}

func (r Result5[T1, T2, T3, T4, T5]) V2() T2 {
	return r.v2
}

func (r Result5[T1, T2, T3, T4, T5]) V3() T3 {
	return r.v3
}

func (r Result5[T1, T2, T3, T4, T5]) V4() T4 {
	return r.v4
}

func (r Result5[T1, T2, T3, T4, T5]) V5() T5 {
	return r.v5
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

// Unwrap return data or panic(err)
func (r Result6[T1, T2, T3, T4, T5, T6]) Unwrap() (T1, T2, T3, T4, T5, T6) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6
}

// UnwrapOr return v1, v2, ...
func (r Result6[T1, T2, T3, T4, T5, T6]) UnwrapOr(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) (T1, T2, T3, T4, T5, T6) {
	if r.err != nil {
		return v1, v2, v3, v4, v5, v6
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6
}

// Unpack return (v1, v2, ..., err)
func (r Result6[T1, T2, T3, T4, T5, T6]) Unpack() (T1, T2, T3, T4, T5, T6, error) {
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6, r.err
}

// Err return err
func (r Result6[T1, T2, T3, T4, T5, T6]) Err() error {
	return r.err
}

// Value return v1, v2, ...
func (r Result6[T1, T2, T3, T4, T5, T6]) Value() (T1, T2, T3, T4, T5, T6) {
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6
}

// IsErr check if err is not nil
func (r Result6[T1, T2, T3, T4, T5, T6]) IsErr() bool {
	return r.err != nil
}

// IsOk check if err is nil
func (r Result6[T1, T2, T3, T4, T5, T6]) IsOk() bool {
	return r.err == nil
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

// Then6 ...
func Then6[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4, U5, U6 any](r Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if r.err != nil {
		return Err6[U1, U2, U3, U4, U5, U6](r.err)
	}
	return op(r.V1(), r.V2(), r.V3(), r.V4(), r.V5(), r.V6())
}

func (r Result6[T1, T2, T3, T4, T5, T6]) V1() T1 {
	return r.v1
}

func (r Result6[T1, T2, T3, T4, T5, T6]) V2() T2 {
	return r.v2
}

func (r Result6[T1, T2, T3, T4, T5, T6]) V3() T3 {
	return r.v3
}

func (r Result6[T1, T2, T3, T4, T5, T6]) V4() T4 {
	return r.v4
}

func (r Result6[T1, T2, T3, T4, T5, T6]) V5() T5 {
	return r.v5
}

func (r Result6[T1, T2, T3, T4, T5, T6]) V6() T6 {
	return r.v6
}
