package anyhow

// Result1 ...
type Result1[T1 any] struct {
	v1  T1
	err error
}

// Unwrap return data or panic(err)
func (r *Result1[T1]) Unwrap() T1 {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1
}

// Unpack return (v1, v2, ..., err)
func (r *Result1[T1]) Unpack() (T1, error) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.err
}

// Err return err
func (r *Result1[T1]) Err() error {
	return r.err
}

// Value return v1, v2, ...
func (r *Result1[T1]) Value() T1 {
	return r.v1
}

// IsErr check if err is not nil
func (r *Result1[T1]) IsErr() bool {
	return r.err != nil
}

// IsOk check if err is nil
func (r *Result1[T1]) IsOk() bool {
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

// Then1 ...
func Then1[T1, U1 any](r Result1[T1], op func(T1) Result1[U1]) Result1[U1] {
	if r.err != nil {
		return Err1[U1](r.err)
	}
	return op(r.V1())
}

func (r *Result1[T1]) V1() T1 {
	return r.v1
}

// Result2 ...
type Result2[T1, T2 any] struct {
	v1  T1
	v2  T2
	err error
}

// Unwrap return data or panic(err)
func (r *Result2[T1, T2]) Unwrap() (T1, T2) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2
}

// Unpack return (v1, v2, ..., err)
func (r *Result2[T1, T2]) Unpack() (T1, T2, error) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.err
}

// Err return err
func (r *Result2[T1, T2]) Err() error {
	return r.err
}

// Value return v1, v2, ...
func (r *Result2[T1, T2]) Value() (T1, T2) {
	return r.v1, r.v2
}

// IsErr check if err is not nil
func (r *Result2[T1, T2]) IsErr() bool {
	return r.err != nil
}

// IsOk check if err is nil
func (r *Result2[T1, T2]) IsOk() bool {
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

// Then2 ...
func Then2[T1, T2, U1, U2 any](r Result2[T1, T2], op func(T1, T2) Result2[U1, U2]) Result2[U1, U2] {
	if r.err != nil {
		return Err2[U1, U2](r.err)
	}
	return op(r.V1(), r.V2())
}

func (r *Result2[T1, T2]) V1() T1 {
	return r.v1
}

func (r *Result2[T1, T2]) V2() T2 {
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
func (r *Result3[T1, T2, T3]) Unwrap() (T1, T2, T3) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3
}

// Unpack return (v1, v2, ..., err)
func (r *Result3[T1, T2, T3]) Unpack() (T1, T2, T3, error) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.err
}

// Err return err
func (r *Result3[T1, T2, T3]) Err() error {
	return r.err
}

// Value return v1, v2, ...
func (r *Result3[T1, T2, T3]) Value() (T1, T2, T3) {
	return r.v1, r.v2, r.v3
}

// IsErr check if err is not nil
func (r *Result3[T1, T2, T3]) IsErr() bool {
	return r.err != nil
}

// IsOk check if err is nil
func (r *Result3[T1, T2, T3]) IsOk() bool {
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

// Then3 ...
func Then3[T1, T2, T3, U1, U2, U3 any](r Result3[T1, T2, T3], op func(T1, T2, T3) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if r.err != nil {
		return Err3[U1, U2, U3](r.err)
	}
	return op(r.V1(), r.V2(), r.V3())
}

func (r *Result3[T1, T2, T3]) V1() T1 {
	return r.v1
}

func (r *Result3[T1, T2, T3]) V2() T2 {
	return r.v2
}

func (r *Result3[T1, T2, T3]) V3() T3 {
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
func (r *Result4[T1, T2, T3, T4]) Unwrap() (T1, T2, T3, T4) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4
}

// Unpack return (v1, v2, ..., err)
func (r *Result4[T1, T2, T3, T4]) Unpack() (T1, T2, T3, T4, error) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.err
}

// Err return err
func (r *Result4[T1, T2, T3, T4]) Err() error {
	return r.err
}

// Value return v1, v2, ...
func (r *Result4[T1, T2, T3, T4]) Value() (T1, T2, T3, T4) {
	return r.v1, r.v2, r.v3, r.v4
}

// IsErr check if err is not nil
func (r *Result4[T1, T2, T3, T4]) IsErr() bool {
	return r.err != nil
}

// IsOk check if err is nil
func (r *Result4[T1, T2, T3, T4]) IsOk() bool {
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

// Then4 ...
func Then4[T1, T2, T3, T4, U1, U2, U3, U4 any](r Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if r.err != nil {
		return Err4[U1, U2, U3, U4](r.err)
	}
	return op(r.V1(), r.V2(), r.V3(), r.V4())
}

func (r *Result4[T1, T2, T3, T4]) V1() T1 {
	return r.v1
}

func (r *Result4[T1, T2, T3, T4]) V2() T2 {
	return r.v2
}

func (r *Result4[T1, T2, T3, T4]) V3() T3 {
	return r.v3
}

func (r *Result4[T1, T2, T3, T4]) V4() T4 {
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
func (r *Result5[T1, T2, T3, T4, T5]) Unwrap() (T1, T2, T3, T4, T5) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5
}

// Unpack return (v1, v2, ..., err)
func (r *Result5[T1, T2, T3, T4, T5]) Unpack() (T1, T2, T3, T4, T5, error) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.err
}

// Err return err
func (r *Result5[T1, T2, T3, T4, T5]) Err() error {
	return r.err
}

// Value return v1, v2, ...
func (r *Result5[T1, T2, T3, T4, T5]) Value() (T1, T2, T3, T4, T5) {
	return r.v1, r.v2, r.v3, r.v4, r.v5
}

// IsErr check if err is not nil
func (r *Result5[T1, T2, T3, T4, T5]) IsErr() bool {
	return r.err != nil
}

// IsOk check if err is nil
func (r *Result5[T1, T2, T3, T4, T5]) IsOk() bool {
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

// Then5 ...
func Then5[T1, T2, T3, T4, T5, U1, U2, U3, U4, U5 any](r Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if r.err != nil {
		return Err5[U1, U2, U3, U4, U5](r.err)
	}
	return op(r.V1(), r.V2(), r.V3(), r.V4(), r.V5())
}

func (r *Result5[T1, T2, T3, T4, T5]) V1() T1 {
	return r.v1
}

func (r *Result5[T1, T2, T3, T4, T5]) V2() T2 {
	return r.v2
}

func (r *Result5[T1, T2, T3, T4, T5]) V3() T3 {
	return r.v3
}

func (r *Result5[T1, T2, T3, T4, T5]) V4() T4 {
	return r.v4
}

func (r *Result5[T1, T2, T3, T4, T5]) V5() T5 {
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
func (r *Result6[T1, T2, T3, T4, T5, T6]) Unwrap() (T1, T2, T3, T4, T5, T6) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6
}

// Unpack return (v1, v2, ..., err)
func (r *Result6[T1, T2, T3, T4, T5, T6]) Unpack() (T1, T2, T3, T4, T5, T6, error) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6, r.err
}

// Err return err
func (r *Result6[T1, T2, T3, T4, T5, T6]) Err() error {
	return r.err
}

// Value return v1, v2, ...
func (r *Result6[T1, T2, T3, T4, T5, T6]) Value() (T1, T2, T3, T4, T5, T6) {
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6
}

// IsErr check if err is not nil
func (r *Result6[T1, T2, T3, T4, T5, T6]) IsErr() bool {
	return r.err != nil
}

// IsOk check if err is nil
func (r *Result6[T1, T2, T3, T4, T5, T6]) IsOk() bool {
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

// Then6 ...
func Then6[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4, U5, U6 any](r Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if r.err != nil {
		return Err6[U1, U2, U3, U4, U5, U6](r.err)
	}
	return op(r.V1(), r.V2(), r.V3(), r.V4(), r.V5(), r.V6())
}

func (r *Result6[T1, T2, T3, T4, T5, T6]) V1() T1 {
	return r.v1
}

func (r *Result6[T1, T2, T3, T4, T5, T6]) V2() T2 {
	return r.v2
}

func (r *Result6[T1, T2, T3, T4, T5, T6]) V3() T3 {
	return r.v3
}

func (r *Result6[T1, T2, T3, T4, T5, T6]) V4() T4 {
	return r.v4
}

func (r *Result6[T1, T2, T3, T4, T5, T6]) V5() T5 {
	return r.v5
}

func (r *Result6[T1, T2, T3, T4, T5, T6]) V6() T6 {
	return r.v6
}

// Result7 ...
type Result7[T1, T2, T3, T4, T5, T6, T7 any] struct {
	v1  T1
	v2  T2
	v3  T3
	v4  T4
	v5  T5
	v6  T6
	v7  T7
	err error
}

// Unwrap return data or panic(err)
func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) Unwrap() (T1, T2, T3, T4, T5, T6, T7) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6, r.v7
}

// Unpack return (v1, v2, ..., err)
func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) Unpack() (T1, T2, T3, T4, T5, T6, T7, error) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6, r.v7, r.err
}

// Err return err
func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) Err() error {
	return r.err
}

// Value return v1, v2, ...
func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) Value() (T1, T2, T3, T4, T5, T6, T7) {
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6, r.v7
}

// IsErr check if err is not nil
func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) IsErr() bool {
	return r.err != nil
}

// IsOk check if err is nil
func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) IsOk() bool {
	return r.err == nil
}

// Ok7 pack v1, v2, ... to Result7
func Ok7[T1, T2, T3, T4, T5, T6, T7 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) Result7[T1, T2, T3, T4, T5, T6, T7] {
	return Result7[T1, T2, T3, T4, T5, T6, T7]{v1: v1, v2: v2, v3: v3, v4: v4, v5: v5, v6: v6, v7: v7, err: nil}
}

// Err7 pack err to Result7
func Err7[T1, T2, T3, T4, T5, T6, T7 any](err error) Result7[T1, T2, T3, T4, T5, T6, T7] {
	return Result7[T1, T2, T3, T4, T5, T6, T7]{err: err}
}

// Then7 ...
func Then7[T1, T2, T3, T4, T5, T6, T7, U1, U2, U3, U4, U5, U6, U7 any](r Result7[T1, T2, T3, T4, T5, T6, T7], op func(T1, T2, T3, T4, T5, T6, T7) Result7[U1, U2, U3, U4, U5, U6, U7]) Result7[U1, U2, U3, U4, U5, U6, U7] {
	if r.err != nil {
		return Err7[U1, U2, U3, U4, U5, U6, U7](r.err)
	}
	return op(r.V1(), r.V2(), r.V3(), r.V4(), r.V5(), r.V6(), r.V7())
}

func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) V1() T1 {
	return r.v1
}

func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) V2() T2 {
	return r.v2
}

func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) V3() T3 {
	return r.v3
}

func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) V4() T4 {
	return r.v4
}

func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) V5() T5 {
	return r.v5
}

func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) V6() T6 {
	return r.v6
}

func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) V7() T7 {
	return r.v7
}

// Result8 ...
type Result8[T1, T2, T3, T4, T5, T6, T7, T8 any] struct {
	v1  T1
	v2  T2
	v3  T3
	v4  T4
	v5  T5
	v6  T6
	v7  T7
	v8  T8
	err error
}

// Unwrap return data or panic(err)
func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) Unwrap() (T1, T2, T3, T4, T5, T6, T7, T8) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6, r.v7, r.v8
}

// Unpack return (v1, v2, ..., err)
func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) Unpack() (T1, T2, T3, T4, T5, T6, T7, T8, error) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6, r.v7, r.v8, r.err
}

// Err return err
func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) Err() error {
	return r.err
}

// Value return v1, v2, ...
func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) Value() (T1, T2, T3, T4, T5, T6, T7, T8) {
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6, r.v7, r.v8
}

// IsErr check if err is not nil
func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) IsErr() bool {
	return r.err != nil
}

// IsOk check if err is nil
func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) IsOk() bool {
	return r.err == nil
}

// Ok8 pack v1, v2, ... to Result8
func Ok8[T1, T2, T3, T4, T5, T6, T7, T8 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) Result8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Result8[T1, T2, T3, T4, T5, T6, T7, T8]{v1: v1, v2: v2, v3: v3, v4: v4, v5: v5, v6: v6, v7: v7, v8: v8, err: nil}
}

// Err8 pack err to Result8
func Err8[T1, T2, T3, T4, T5, T6, T7, T8 any](err error) Result8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Result8[T1, T2, T3, T4, T5, T6, T7, T8]{err: err}
}

// Then8 ...
func Then8[T1, T2, T3, T4, T5, T6, T7, T8, U1, U2, U3, U4, U5, U6, U7, U8 any](r Result8[T1, T2, T3, T4, T5, T6, T7, T8], op func(T1, T2, T3, T4, T5, T6, T7, T8) Result8[U1, U2, U3, U4, U5, U6, U7, U8]) Result8[U1, U2, U3, U4, U5, U6, U7, U8] {
	if r.err != nil {
		return Err8[U1, U2, U3, U4, U5, U6, U7, U8](r.err)
	}
	return op(r.V1(), r.V2(), r.V3(), r.V4(), r.V5(), r.V6(), r.V7(), r.V8())
}

func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) V1() T1 {
	return r.v1
}

func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) V2() T2 {
	return r.v2
}

func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) V3() T3 {
	return r.v3
}

func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) V4() T4 {
	return r.v4
}

func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) V5() T5 {
	return r.v5
}

func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) V6() T6 {
	return r.v6
}

func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) V7() T7 {
	return r.v7
}

func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) V8() T8 {
	return r.v8
}

// Result9 ...
type Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] struct {
	v1  T1
	v2  T2
	v3  T3
	v4  T4
	v5  T5
	v6  T6
	v7  T7
	v8  T8
	v9  T9
	err error
}

// Unwrap return data or panic(err)
func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Unwrap() (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6, r.v7, r.v8, r.v9
}

// Unpack return (v1, v2, ..., err)
func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Unpack() (T1, T2, T3, T4, T5, T6, T7, T8, T9, error) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6, r.v7, r.v8, r.v9, r.err
}

// Err return err
func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Err() error {
	return r.err
}

// Value return v1, v2, ...
func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Value() (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6, r.v7, r.v8, r.v9
}

// IsErr check if err is not nil
func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) IsErr() bool {
	return r.err != nil
}

// IsOk check if err is nil
func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) IsOk() bool {
	return r.err == nil
}

// Ok9 pack v1, v2, ... to Result9
func Ok9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9) Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{v1: v1, v2: v2, v3: v3, v4: v4, v5: v5, v6: v6, v7: v7, v8: v8, v9: v9, err: nil}
}

// Err9 pack err to Result9
func Err9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](err error) Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{err: err}
}

// Then9 ...
func Then9[T1, T2, T3, T4, T5, T6, T7, T8, T9, U1, U2, U3, U4, U5, U6, U7, U8, U9 any](r Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9], op func(T1, T2, T3, T4, T5, T6, T7, T8, T9) Result9[U1, U2, U3, U4, U5, U6, U7, U8, U9]) Result9[U1, U2, U3, U4, U5, U6, U7, U8, U9] {
	if r.err != nil {
		return Err9[U1, U2, U3, U4, U5, U6, U7, U8, U9](r.err)
	}
	return op(r.V1(), r.V2(), r.V3(), r.V4(), r.V5(), r.V6(), r.V7(), r.V8(), r.V9())
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) V1() T1 {
	return r.v1
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) V2() T2 {
	return r.v2
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) V3() T3 {
	return r.v3
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) V4() T4 {
	return r.v4
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) V5() T5 {
	return r.v5
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) V6() T6 {
	return r.v6
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) V7() T7 {
	return r.v7
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) V8() T8 {
	return r.v8
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) V9() T9 {
	return r.v9
}
