package anyhow

import "fmt"

// IntoErr Returns the contained Err value, but never panics
func (r Result1[T1]) IntoErr() error {
	return r.err
}

// IntoOk Returns the contained Ok value, but never panics
func (r Result1[T1]) IntoOk() T1 {
	return r.v1
}

// Expect Returns the contained Ok value, otherwise panics with a panic message including the passed message, and the content of the Err
func (r Result1[T1]) Expect(msg string) T1 {
	if r.IsOk() {
		return r.v1
	}
	panic(fmt.Errorf("%s: %v", msg, r.err))
}

// ExpectErr Returns the contained Err value, otherwise panics with a panic message including the passed message, and the content of the Ok
func (r Result1[T1]) ExpectErr(msg string) error {
	if r.IsErr() {
		return r.err
	}
	panic(fmt.Errorf("%s: %v", msg, r.v1))
}

// Inspect calls a function with a reference to the contained value if Ok
func (r Result1[T1]) Inspect(f func(T1)) Result1[T1] {
	if r.IsOk() {
		f(r.IntoOk())
	}
	return r
}

// InspectErr calls a function with a reference to the contained value if Err
func (r Result1[T1]) InspectErr(f func(error)) Result1[T1] {
	if r.IsErr() {
		f(r.err)
	}
	return r
}

// MapErr Maps a R[T] to R[T] by applying a function to a contained Err value, leaving an Ok value untouched
func (r Result1[T1]) MapErr(op func(error) error) Result1[T1] {
	if r.IsErr() {
		return Err1[T1](op(r.err))
	}
	return r
}

// IsErr returns true if the result is Err
func (r Result1[T1]) IsErr() bool {
	return r.err != nil
}

// IsOk returns true if the result is Ok
func (r Result1[T1]) IsOk() bool {
	return r.err == nil
}

// Or Returns res if the result is Err, otherwise returns the Ok value of self
func (r Result1[T1]) Or(res Result1[T1]) Result1[T1] {
	if r.IsErr() {
		return res
	}
	return r
}

// OrElse Calls op if the result is Err, otherwise returns the Ok value of self
func (r Result1[T1]) OrElse(op func(error) Result1[T1]) Result1[T1] {
	if r.IsErr() {
		return op(r.err)
	}
	return r
}

// Unwrap Returns the contained Ok value, otherwise panics with Err
func (r Result1[T1]) Unwrap() T1 {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1
}

// UnwrapOr Returns the contained Ok value or a provided default
func (r Result1[T1]) UnwrapOr(v1 T1) T1 {
	if r.err != nil {
		return v1
	}
	return r.v1
}

// UnwrapOrDefault Returns the contained Ok value or a default value
func (r Result1[T1]) UnwrapOrDefault() T1 {
	if r.err != nil {
		return zero[T1]()
	}
	return r.v1
}

// UnwrapOrElse Returns the contained Ok value or computes it from a closure
func (r Result1[T1]) UnwrapOrElse(op func(error) T1) T1 {
	if r.err != nil {
		return op(r.err)
	}
	return r.v1
}

// Unpack return (v1, v2, ..., err)
func (r Result1[T1]) Unpack() (T1, error) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.err
}

func (r Result1[T1]) V1() T1 {
	return r.v1
}

// IntoErr Returns the contained Err value, but never panics
func (r Result2[T1, T2]) IntoErr() error {
	return r.err
}

// IntoOk Returns the contained Ok value, but never panics
func (r Result2[T1, T2]) IntoOk() (T1, T2) {
	return r.v1, r.v2
}

// Expect Returns the contained Ok value, otherwise panics with a panic message including the passed message, and the content of the Err
func (r Result2[T1, T2]) Expect(msg string) (T1, T2) {
	if r.IsOk() {
		return r.v1, r.v2
	}
	panic(fmt.Errorf("%s: %v", msg, r.err))
}

// ExpectErr Returns the contained Err value, otherwise panics with a panic message including the passed message, and the content of the Ok
func (r Result2[T1, T2]) ExpectErr(msg string) error {
	if r.IsErr() {
		return r.err
	}
	panic(fmt.Errorf("%s: %v, %v", msg, r.v1, r.v2))
}

// Inspect calls a function with a reference to the contained value if Ok
func (r Result2[T1, T2]) Inspect(f func(T1, T2)) Result2[T1, T2] {
	if r.IsOk() {
		f(r.IntoOk())
	}
	return r
}

// InspectErr calls a function with a reference to the contained value if Err
func (r Result2[T1, T2]) InspectErr(f func(error)) Result2[T1, T2] {
	if r.IsErr() {
		f(r.err)
	}
	return r
}

// MapErr Maps a R[T] to R[T] by applying a function to a contained Err value, leaving an Ok value untouched
func (r Result2[T1, T2]) MapErr(op func(error) error) Result2[T1, T2] {
	if r.IsErr() {
		return Err2[T1, T2](op(r.err))
	}
	return r
}

// IsErr returns true if the result is Err
func (r Result2[T1, T2]) IsErr() bool {
	return r.err != nil
}

// IsOk returns true if the result is Ok
func (r Result2[T1, T2]) IsOk() bool {
	return r.err == nil
}

// Or Returns res if the result is Err, otherwise returns the Ok value of self
func (r Result2[T1, T2]) Or(res Result2[T1, T2]) Result2[T1, T2] {
	if r.IsErr() {
		return res
	}
	return r
}

// OrElse Calls op if the result is Err, otherwise returns the Ok value of self
func (r Result2[T1, T2]) OrElse(op func(error) Result2[T1, T2]) Result2[T1, T2] {
	if r.IsErr() {
		return op(r.err)
	}
	return r
}

// Unwrap Returns the contained Ok value, otherwise panics with Err
func (r Result2[T1, T2]) Unwrap() (T1, T2) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2
}

// UnwrapOr Returns the contained Ok value or a provided default
func (r Result2[T1, T2]) UnwrapOr(v1 T1, v2 T2) (T1, T2) {
	if r.err != nil {
		return v1, v2
	}
	return r.v1, r.v2
}

// UnwrapOrDefault Returns the contained Ok value or a default value
func (r Result2[T1, T2]) UnwrapOrDefault() (T1, T2) {
	if r.err != nil {
		return zero[T1](), zero[T2]()
	}
	return r.v1, r.v2
}

// UnwrapOrElse Returns the contained Ok value or computes it from a closure
func (r Result2[T1, T2]) UnwrapOrElse(op func(error) (T1, T2)) (T1, T2) {
	if r.err != nil {
		return op(r.err)
	}
	return r.v1, r.v2
}

// Unpack return (v1, v2, ..., err)
func (r Result2[T1, T2]) Unpack() (T1, T2, error) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.err
}

func (r Result2[T1, T2]) V1() T1 {
	return r.v1
}

func (r Result2[T1, T2]) V2() T2 {
	return r.v2
}

// IntoErr Returns the contained Err value, but never panics
func (r Result3[T1, T2, T3]) IntoErr() error {
	return r.err
}

// IntoOk Returns the contained Ok value, but never panics
func (r Result3[T1, T2, T3]) IntoOk() (T1, T2, T3) {
	return r.v1, r.v2, r.v3
}

// Expect Returns the contained Ok value, otherwise panics with a panic message including the passed message, and the content of the Err
func (r Result3[T1, T2, T3]) Expect(msg string) (T1, T2, T3) {
	if r.IsOk() {
		return r.v1, r.v2, r.v3
	}
	panic(fmt.Errorf("%s: %v", msg, r.err))
}

// ExpectErr Returns the contained Err value, otherwise panics with a panic message including the passed message, and the content of the Ok
func (r Result3[T1, T2, T3]) ExpectErr(msg string) error {
	if r.IsErr() {
		return r.err
	}
	panic(fmt.Errorf("%s: %v, %v, %v", msg, r.v1, r.v2, r.v3))
}

// Inspect calls a function with a reference to the contained value if Ok
func (r Result3[T1, T2, T3]) Inspect(f func(T1, T2, T3)) Result3[T1, T2, T3] {
	if r.IsOk() {
		f(r.IntoOk())
	}
	return r
}

// InspectErr calls a function with a reference to the contained value if Err
func (r Result3[T1, T2, T3]) InspectErr(f func(error)) Result3[T1, T2, T3] {
	if r.IsErr() {
		f(r.err)
	}
	return r
}

// MapErr Maps a R[T] to R[T] by applying a function to a contained Err value, leaving an Ok value untouched
func (r Result3[T1, T2, T3]) MapErr(op func(error) error) Result3[T1, T2, T3] {
	if r.IsErr() {
		return Err3[T1, T2, T3](op(r.err))
	}
	return r
}

// IsErr returns true if the result is Err
func (r Result3[T1, T2, T3]) IsErr() bool {
	return r.err != nil
}

// IsOk returns true if the result is Ok
func (r Result3[T1, T2, T3]) IsOk() bool {
	return r.err == nil
}

// Or Returns res if the result is Err, otherwise returns the Ok value of self
func (r Result3[T1, T2, T3]) Or(res Result3[T1, T2, T3]) Result3[T1, T2, T3] {
	if r.IsErr() {
		return res
	}
	return r
}

// OrElse Calls op if the result is Err, otherwise returns the Ok value of self
func (r Result3[T1, T2, T3]) OrElse(op func(error) Result3[T1, T2, T3]) Result3[T1, T2, T3] {
	if r.IsErr() {
		return op(r.err)
	}
	return r
}

// Unwrap Returns the contained Ok value, otherwise panics with Err
func (r Result3[T1, T2, T3]) Unwrap() (T1, T2, T3) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3
}

// UnwrapOr Returns the contained Ok value or a provided default
func (r Result3[T1, T2, T3]) UnwrapOr(v1 T1, v2 T2, v3 T3) (T1, T2, T3) {
	if r.err != nil {
		return v1, v2, v3
	}
	return r.v1, r.v2, r.v3
}

// UnwrapOrDefault Returns the contained Ok value or a default value
func (r Result3[T1, T2, T3]) UnwrapOrDefault() (T1, T2, T3) {
	if r.err != nil {
		return zero[T1](), zero[T2](), zero[T3]()
	}
	return r.v1, r.v2, r.v3
}

// UnwrapOrElse Returns the contained Ok value or computes it from a closure
func (r Result3[T1, T2, T3]) UnwrapOrElse(op func(error) (T1, T2, T3)) (T1, T2, T3) {
	if r.err != nil {
		return op(r.err)
	}
	return r.v1, r.v2, r.v3
}

// Unpack return (v1, v2, ..., err)
func (r Result3[T1, T2, T3]) Unpack() (T1, T2, T3, error) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.err
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

// IntoErr Returns the contained Err value, but never panics
func (r Result4[T1, T2, T3, T4]) IntoErr() error {
	return r.err
}

// IntoOk Returns the contained Ok value, but never panics
func (r Result4[T1, T2, T3, T4]) IntoOk() (T1, T2, T3, T4) {
	return r.v1, r.v2, r.v3, r.v4
}

// Expect Returns the contained Ok value, otherwise panics with a panic message including the passed message, and the content of the Err
func (r Result4[T1, T2, T3, T4]) Expect(msg string) (T1, T2, T3, T4) {
	if r.IsOk() {
		return r.v1, r.v2, r.v3, r.v4
	}
	panic(fmt.Errorf("%s: %v", msg, r.err))
}

// ExpectErr Returns the contained Err value, otherwise panics with a panic message including the passed message, and the content of the Ok
func (r Result4[T1, T2, T3, T4]) ExpectErr(msg string) error {
	if r.IsErr() {
		return r.err
	}
	panic(fmt.Errorf("%s: %v, %v, %v, %v", msg, r.v1, r.v2, r.v3, r.v4))
}

// Inspect calls a function with a reference to the contained value if Ok
func (r Result4[T1, T2, T3, T4]) Inspect(f func(T1, T2, T3, T4)) Result4[T1, T2, T3, T4] {
	if r.IsOk() {
		f(r.IntoOk())
	}
	return r
}

// InspectErr calls a function with a reference to the contained value if Err
func (r Result4[T1, T2, T3, T4]) InspectErr(f func(error)) Result4[T1, T2, T3, T4] {
	if r.IsErr() {
		f(r.err)
	}
	return r
}

// MapErr Maps a R[T] to R[T] by applying a function to a contained Err value, leaving an Ok value untouched
func (r Result4[T1, T2, T3, T4]) MapErr(op func(error) error) Result4[T1, T2, T3, T4] {
	if r.IsErr() {
		return Err4[T1, T2, T3, T4](op(r.err))
	}
	return r
}

// IsErr returns true if the result is Err
func (r Result4[T1, T2, T3, T4]) IsErr() bool {
	return r.err != nil
}

// IsOk returns true if the result is Ok
func (r Result4[T1, T2, T3, T4]) IsOk() bool {
	return r.err == nil
}

// Or Returns res if the result is Err, otherwise returns the Ok value of self
func (r Result4[T1, T2, T3, T4]) Or(res Result4[T1, T2, T3, T4]) Result4[T1, T2, T3, T4] {
	if r.IsErr() {
		return res
	}
	return r
}

// OrElse Calls op if the result is Err, otherwise returns the Ok value of self
func (r Result4[T1, T2, T3, T4]) OrElse(op func(error) Result4[T1, T2, T3, T4]) Result4[T1, T2, T3, T4] {
	if r.IsErr() {
		return op(r.err)
	}
	return r
}

// Unwrap Returns the contained Ok value, otherwise panics with Err
func (r Result4[T1, T2, T3, T4]) Unwrap() (T1, T2, T3, T4) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4
}

// UnwrapOr Returns the contained Ok value or a provided default
func (r Result4[T1, T2, T3, T4]) UnwrapOr(v1 T1, v2 T2, v3 T3, v4 T4) (T1, T2, T3, T4) {
	if r.err != nil {
		return v1, v2, v3, v4
	}
	return r.v1, r.v2, r.v3, r.v4
}

// UnwrapOrDefault Returns the contained Ok value or a default value
func (r Result4[T1, T2, T3, T4]) UnwrapOrDefault() (T1, T2, T3, T4) {
	if r.err != nil {
		return zero[T1](), zero[T2](), zero[T3](), zero[T4]()
	}
	return r.v1, r.v2, r.v3, r.v4
}

// UnwrapOrElse Returns the contained Ok value or computes it from a closure
func (r Result4[T1, T2, T3, T4]) UnwrapOrElse(op func(error) (T1, T2, T3, T4)) (T1, T2, T3, T4) {
	if r.err != nil {
		return op(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4
}

// Unpack return (v1, v2, ..., err)
func (r Result4[T1, T2, T3, T4]) Unpack() (T1, T2, T3, T4, error) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.err
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

// IntoErr Returns the contained Err value, but never panics
func (r Result5[T1, T2, T3, T4, T5]) IntoErr() error {
	return r.err
}

// IntoOk Returns the contained Ok value, but never panics
func (r Result5[T1, T2, T3, T4, T5]) IntoOk() (T1, T2, T3, T4, T5) {
	return r.v1, r.v2, r.v3, r.v4, r.v5
}

// Expect Returns the contained Ok value, otherwise panics with a panic message including the passed message, and the content of the Err
func (r Result5[T1, T2, T3, T4, T5]) Expect(msg string) (T1, T2, T3, T4, T5) {
	if r.IsOk() {
		return r.v1, r.v2, r.v3, r.v4, r.v5
	}
	panic(fmt.Errorf("%s: %v", msg, r.err))
}

// ExpectErr Returns the contained Err value, otherwise panics with a panic message including the passed message, and the content of the Ok
func (r Result5[T1, T2, T3, T4, T5]) ExpectErr(msg string) error {
	if r.IsErr() {
		return r.err
	}
	panic(fmt.Errorf("%s: %v, %v, %v, %v, %v", msg, r.v1, r.v2, r.v3, r.v4, r.v5))
}

// Inspect calls a function with a reference to the contained value if Ok
func (r Result5[T1, T2, T3, T4, T5]) Inspect(f func(T1, T2, T3, T4, T5)) Result5[T1, T2, T3, T4, T5] {
	if r.IsOk() {
		f(r.IntoOk())
	}
	return r
}

// InspectErr calls a function with a reference to the contained value if Err
func (r Result5[T1, T2, T3, T4, T5]) InspectErr(f func(error)) Result5[T1, T2, T3, T4, T5] {
	if r.IsErr() {
		f(r.err)
	}
	return r
}

// MapErr Maps a R[T] to R[T] by applying a function to a contained Err value, leaving an Ok value untouched
func (r Result5[T1, T2, T3, T4, T5]) MapErr(op func(error) error) Result5[T1, T2, T3, T4, T5] {
	if r.IsErr() {
		return Err5[T1, T2, T3, T4, T5](op(r.err))
	}
	return r
}

// IsErr returns true if the result is Err
func (r Result5[T1, T2, T3, T4, T5]) IsErr() bool {
	return r.err != nil
}

// IsOk returns true if the result is Ok
func (r Result5[T1, T2, T3, T4, T5]) IsOk() bool {
	return r.err == nil
}

// Or Returns res if the result is Err, otherwise returns the Ok value of self
func (r Result5[T1, T2, T3, T4, T5]) Or(res Result5[T1, T2, T3, T4, T5]) Result5[T1, T2, T3, T4, T5] {
	if r.IsErr() {
		return res
	}
	return r
}

// OrElse Calls op if the result is Err, otherwise returns the Ok value of self
func (r Result5[T1, T2, T3, T4, T5]) OrElse(op func(error) Result5[T1, T2, T3, T4, T5]) Result5[T1, T2, T3, T4, T5] {
	if r.IsErr() {
		return op(r.err)
	}
	return r
}

// Unwrap Returns the contained Ok value, otherwise panics with Err
func (r Result5[T1, T2, T3, T4, T5]) Unwrap() (T1, T2, T3, T4, T5) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5
}

// UnwrapOr Returns the contained Ok value or a provided default
func (r Result5[T1, T2, T3, T4, T5]) UnwrapOr(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) (T1, T2, T3, T4, T5) {
	if r.err != nil {
		return v1, v2, v3, v4, v5
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5
}

// UnwrapOrDefault Returns the contained Ok value or a default value
func (r Result5[T1, T2, T3, T4, T5]) UnwrapOrDefault() (T1, T2, T3, T4, T5) {
	if r.err != nil {
		return zero[T1](), zero[T2](), zero[T3](), zero[T4](), zero[T5]()
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5
}

// UnwrapOrElse Returns the contained Ok value or computes it from a closure
func (r Result5[T1, T2, T3, T4, T5]) UnwrapOrElse(op func(error) (T1, T2, T3, T4, T5)) (T1, T2, T3, T4, T5) {
	if r.err != nil {
		return op(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5
}

// Unpack return (v1, v2, ..., err)
func (r Result5[T1, T2, T3, T4, T5]) Unpack() (T1, T2, T3, T4, T5, error) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.err
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

// IntoErr Returns the contained Err value, but never panics
func (r Result6[T1, T2, T3, T4, T5, T6]) IntoErr() error {
	return r.err
}

// IntoOk Returns the contained Ok value, but never panics
func (r Result6[T1, T2, T3, T4, T5, T6]) IntoOk() (T1, T2, T3, T4, T5, T6) {
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6
}

// Expect Returns the contained Ok value, otherwise panics with a panic message including the passed message, and the content of the Err
func (r Result6[T1, T2, T3, T4, T5, T6]) Expect(msg string) (T1, T2, T3, T4, T5, T6) {
	if r.IsOk() {
		return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6
	}
	panic(fmt.Errorf("%s: %v", msg, r.err))
}

// ExpectErr Returns the contained Err value, otherwise panics with a panic message including the passed message, and the content of the Ok
func (r Result6[T1, T2, T3, T4, T5, T6]) ExpectErr(msg string) error {
	if r.IsErr() {
		return r.err
	}
	panic(fmt.Errorf("%s: %v, %v, %v, %v, %v, %v", msg, r.v1, r.v2, r.v3, r.v4, r.v5, r.v6))
}

// Inspect calls a function with a reference to the contained value if Ok
func (r Result6[T1, T2, T3, T4, T5, T6]) Inspect(f func(T1, T2, T3, T4, T5, T6)) Result6[T1, T2, T3, T4, T5, T6] {
	if r.IsOk() {
		f(r.IntoOk())
	}
	return r
}

// InspectErr calls a function with a reference to the contained value if Err
func (r Result6[T1, T2, T3, T4, T5, T6]) InspectErr(f func(error)) Result6[T1, T2, T3, T4, T5, T6] {
	if r.IsErr() {
		f(r.err)
	}
	return r
}

// MapErr Maps a R[T] to R[T] by applying a function to a contained Err value, leaving an Ok value untouched
func (r Result6[T1, T2, T3, T4, T5, T6]) MapErr(op func(error) error) Result6[T1, T2, T3, T4, T5, T6] {
	if r.IsErr() {
		return Err6[T1, T2, T3, T4, T5, T6](op(r.err))
	}
	return r
}

// IsErr returns true if the result is Err
func (r Result6[T1, T2, T3, T4, T5, T6]) IsErr() bool {
	return r.err != nil
}

// IsOk returns true if the result is Ok
func (r Result6[T1, T2, T3, T4, T5, T6]) IsOk() bool {
	return r.err == nil
}

// Or Returns res if the result is Err, otherwise returns the Ok value of self
func (r Result6[T1, T2, T3, T4, T5, T6]) Or(res Result6[T1, T2, T3, T4, T5, T6]) Result6[T1, T2, T3, T4, T5, T6] {
	if r.IsErr() {
		return res
	}
	return r
}

// OrElse Calls op if the result is Err, otherwise returns the Ok value of self
func (r Result6[T1, T2, T3, T4, T5, T6]) OrElse(op func(error) Result6[T1, T2, T3, T4, T5, T6]) Result6[T1, T2, T3, T4, T5, T6] {
	if r.IsErr() {
		return op(r.err)
	}
	return r
}

// Unwrap Returns the contained Ok value, otherwise panics with Err
func (r Result6[T1, T2, T3, T4, T5, T6]) Unwrap() (T1, T2, T3, T4, T5, T6) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6
}

// UnwrapOr Returns the contained Ok value or a provided default
func (r Result6[T1, T2, T3, T4, T5, T6]) UnwrapOr(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) (T1, T2, T3, T4, T5, T6) {
	if r.err != nil {
		return v1, v2, v3, v4, v5, v6
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6
}

// UnwrapOrDefault Returns the contained Ok value or a default value
func (r Result6[T1, T2, T3, T4, T5, T6]) UnwrapOrDefault() (T1, T2, T3, T4, T5, T6) {
	if r.err != nil {
		return zero[T1](), zero[T2](), zero[T3](), zero[T4](), zero[T5](), zero[T6]()
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6
}

// UnwrapOrElse Returns the contained Ok value or computes it from a closure
func (r Result6[T1, T2, T3, T4, T5, T6]) UnwrapOrElse(op func(error) (T1, T2, T3, T4, T5, T6)) (T1, T2, T3, T4, T5, T6) {
	if r.err != nil {
		return op(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6
}

// Unpack return (v1, v2, ..., err)
func (r Result6[T1, T2, T3, T4, T5, T6]) Unpack() (T1, T2, T3, T4, T5, T6, error) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6, r.err
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
