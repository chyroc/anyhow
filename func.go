package anyhow

// And11 Returns res if the result is Ok, otherwise returns the Err value of self
func And11[T1, U1 any](self Result1[T1], res Result1[U1]) Result1[U1] {
	if self.IsOk() {
		return res
	}
	return Err1[U1](self.err)
}

// AndThen11 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen11[T1, U1 any](self Result1[T1], op func(T1) Result1[U1]) Result1[U1] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err1[U1](self.err)
}

// Map11 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map11[T1, U1 any](self Result1[T1], op func(T1) U1) Result1[U1] {
	if self.err != nil {
		return Err1[U1](self.err)
	}
	return Ok1[U1](op(self.Value()))
}

// MapOr11 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr11[T1, U1 any](self Result1[T1], defaultV1 U1, op func(T1) Result1[U1]) Result1[U1] {
	if self.err != nil {
		return Ok1[U1](defaultV1)
	}
	return op(self.Value())
}

// And12 Returns res if the result is Ok, otherwise returns the Err value of self
func And12[T1, U1, U2 any](self Result1[T1], res Result2[U1, U2]) Result2[U1, U2] {
	if self.IsOk() {
		return res
	}
	return Err2[U1, U2](self.err)
}

// AndThen12 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen12[T1, U1, U2 any](self Result1[T1], op func(T1) Result2[U1, U2]) Result2[U1, U2] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err2[U1, U2](self.err)
}

// Map12 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map12[T1, U1, U2 any](self Result1[T1], op func(T1) (U1, U2)) Result2[U1, U2] {
	if self.err != nil {
		return Err2[U1, U2](self.err)
	}
	return Ok2[U1, U2](op(self.Value()))
}

// MapOr12 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr12[T1, U1, U2 any](self Result1[T1], defaultV1 U1, defaultV2 U2, op func(T1) Result2[U1, U2]) Result2[U1, U2] {
	if self.err != nil {
		return Ok2[U1, U2](defaultV1, defaultV2)
	}
	return op(self.Value())
}

// And13 Returns res if the result is Ok, otherwise returns the Err value of self
func And13[T1, U1, U2, U3 any](self Result1[T1], res Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if self.IsOk() {
		return res
	}
	return Err3[U1, U2, U3](self.err)
}

// AndThen13 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen13[T1, U1, U2, U3 any](self Result1[T1], op func(T1) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err3[U1, U2, U3](self.err)
}

// Map13 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map13[T1, U1, U2, U3 any](self Result1[T1], op func(T1) (U1, U2, U3)) Result3[U1, U2, U3] {
	if self.err != nil {
		return Err3[U1, U2, U3](self.err)
	}
	return Ok3[U1, U2, U3](op(self.Value()))
}

// MapOr13 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr13[T1, U1, U2, U3 any](self Result1[T1], defaultV1 U1, defaultV2 U2, defaultV3 U3, op func(T1) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if self.err != nil {
		return Ok3[U1, U2, U3](defaultV1, defaultV2, defaultV3)
	}
	return op(self.Value())
}

// And14 Returns res if the result is Ok, otherwise returns the Err value of self
func And14[T1, U1, U2, U3, U4 any](self Result1[T1], res Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if self.IsOk() {
		return res
	}
	return Err4[U1, U2, U3, U4](self.err)
}

// AndThen14 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen14[T1, U1, U2, U3, U4 any](self Result1[T1], op func(T1) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err4[U1, U2, U3, U4](self.err)
}

// Map14 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map14[T1, U1, U2, U3, U4 any](self Result1[T1], op func(T1) (U1, U2, U3, U4)) Result4[U1, U2, U3, U4] {
	if self.err != nil {
		return Err4[U1, U2, U3, U4](self.err)
	}
	return Ok4[U1, U2, U3, U4](op(self.Value()))
}

// MapOr14 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr14[T1, U1, U2, U3, U4 any](self Result1[T1], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, op func(T1) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if self.err != nil {
		return Ok4[U1, U2, U3, U4](defaultV1, defaultV2, defaultV3, defaultV4)
	}
	return op(self.Value())
}

// And15 Returns res if the result is Ok, otherwise returns the Err value of self
func And15[T1, U1, U2, U3, U4, U5 any](self Result1[T1], res Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if self.IsOk() {
		return res
	}
	return Err5[U1, U2, U3, U4, U5](self.err)
}

// AndThen15 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen15[T1, U1, U2, U3, U4, U5 any](self Result1[T1], op func(T1) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err5[U1, U2, U3, U4, U5](self.err)
}

// Map15 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map15[T1, U1, U2, U3, U4, U5 any](self Result1[T1], op func(T1) (U1, U2, U3, U4, U5)) Result5[U1, U2, U3, U4, U5] {
	if self.err != nil {
		return Err5[U1, U2, U3, U4, U5](self.err)
	}
	return Ok5[U1, U2, U3, U4, U5](op(self.Value()))
}

// MapOr15 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr15[T1, U1, U2, U3, U4, U5 any](self Result1[T1], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, op func(T1) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if self.err != nil {
		return Ok5[U1, U2, U3, U4, U5](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5)
	}
	return op(self.Value())
}

// And16 Returns res if the result is Ok, otherwise returns the Err value of self
func And16[T1, U1, U2, U3, U4, U5, U6 any](self Result1[T1], res Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if self.IsOk() {
		return res
	}
	return Err6[U1, U2, U3, U4, U5, U6](self.err)
}

// AndThen16 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen16[T1, U1, U2, U3, U4, U5, U6 any](self Result1[T1], op func(T1) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err6[U1, U2, U3, U4, U5, U6](self.err)
}

// Map16 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map16[T1, U1, U2, U3, U4, U5, U6 any](self Result1[T1], op func(T1) (U1, U2, U3, U4, U5, U6)) Result6[U1, U2, U3, U4, U5, U6] {
	if self.err != nil {
		return Err6[U1, U2, U3, U4, U5, U6](self.err)
	}
	return Ok6[U1, U2, U3, U4, U5, U6](op(self.Value()))
}

// MapOr16 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr16[T1, U1, U2, U3, U4, U5, U6 any](self Result1[T1], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, defaultV6 U6, op func(T1) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if self.err != nil {
		return Ok6[U1, U2, U3, U4, U5, U6](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5, defaultV6)
	}
	return op(self.Value())
}

// And21 Returns res if the result is Ok, otherwise returns the Err value of self
func And21[T1, T2, U1 any](self Result2[T1, T2], res Result1[U1]) Result1[U1] {
	if self.IsOk() {
		return res
	}
	return Err1[U1](self.err)
}

// AndThen21 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen21[T1, T2, U1 any](self Result2[T1, T2], op func(T1, T2) Result1[U1]) Result1[U1] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err1[U1](self.err)
}

// Map21 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map21[T1, T2, U1 any](self Result2[T1, T2], op func(T1, T2) U1) Result1[U1] {
	if self.err != nil {
		return Err1[U1](self.err)
	}
	return Ok1[U1](op(self.Value()))
}

// MapOr21 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr21[T1, T2, U1 any](self Result2[T1, T2], defaultV1 U1, op func(T1, T2) Result1[U1]) Result1[U1] {
	if self.err != nil {
		return Ok1[U1](defaultV1)
	}
	return op(self.Value())
}

// And22 Returns res if the result is Ok, otherwise returns the Err value of self
func And22[T1, T2, U1, U2 any](self Result2[T1, T2], res Result2[U1, U2]) Result2[U1, U2] {
	if self.IsOk() {
		return res
	}
	return Err2[U1, U2](self.err)
}

// AndThen22 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen22[T1, T2, U1, U2 any](self Result2[T1, T2], op func(T1, T2) Result2[U1, U2]) Result2[U1, U2] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err2[U1, U2](self.err)
}

// Map22 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map22[T1, T2, U1, U2 any](self Result2[T1, T2], op func(T1, T2) (U1, U2)) Result2[U1, U2] {
	if self.err != nil {
		return Err2[U1, U2](self.err)
	}
	return Ok2[U1, U2](op(self.Value()))
}

// MapOr22 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr22[T1, T2, U1, U2 any](self Result2[T1, T2], defaultV1 U1, defaultV2 U2, op func(T1, T2) Result2[U1, U2]) Result2[U1, U2] {
	if self.err != nil {
		return Ok2[U1, U2](defaultV1, defaultV2)
	}
	return op(self.Value())
}

// And23 Returns res if the result is Ok, otherwise returns the Err value of self
func And23[T1, T2, U1, U2, U3 any](self Result2[T1, T2], res Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if self.IsOk() {
		return res
	}
	return Err3[U1, U2, U3](self.err)
}

// AndThen23 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen23[T1, T2, U1, U2, U3 any](self Result2[T1, T2], op func(T1, T2) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err3[U1, U2, U3](self.err)
}

// Map23 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map23[T1, T2, U1, U2, U3 any](self Result2[T1, T2], op func(T1, T2) (U1, U2, U3)) Result3[U1, U2, U3] {
	if self.err != nil {
		return Err3[U1, U2, U3](self.err)
	}
	return Ok3[U1, U2, U3](op(self.Value()))
}

// MapOr23 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr23[T1, T2, U1, U2, U3 any](self Result2[T1, T2], defaultV1 U1, defaultV2 U2, defaultV3 U3, op func(T1, T2) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if self.err != nil {
		return Ok3[U1, U2, U3](defaultV1, defaultV2, defaultV3)
	}
	return op(self.Value())
}

// And24 Returns res if the result is Ok, otherwise returns the Err value of self
func And24[T1, T2, U1, U2, U3, U4 any](self Result2[T1, T2], res Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if self.IsOk() {
		return res
	}
	return Err4[U1, U2, U3, U4](self.err)
}

// AndThen24 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen24[T1, T2, U1, U2, U3, U4 any](self Result2[T1, T2], op func(T1, T2) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err4[U1, U2, U3, U4](self.err)
}

// Map24 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map24[T1, T2, U1, U2, U3, U4 any](self Result2[T1, T2], op func(T1, T2) (U1, U2, U3, U4)) Result4[U1, U2, U3, U4] {
	if self.err != nil {
		return Err4[U1, U2, U3, U4](self.err)
	}
	return Ok4[U1, U2, U3, U4](op(self.Value()))
}

// MapOr24 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr24[T1, T2, U1, U2, U3, U4 any](self Result2[T1, T2], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, op func(T1, T2) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if self.err != nil {
		return Ok4[U1, U2, U3, U4](defaultV1, defaultV2, defaultV3, defaultV4)
	}
	return op(self.Value())
}

// And25 Returns res if the result is Ok, otherwise returns the Err value of self
func And25[T1, T2, U1, U2, U3, U4, U5 any](self Result2[T1, T2], res Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if self.IsOk() {
		return res
	}
	return Err5[U1, U2, U3, U4, U5](self.err)
}

// AndThen25 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen25[T1, T2, U1, U2, U3, U4, U5 any](self Result2[T1, T2], op func(T1, T2) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err5[U1, U2, U3, U4, U5](self.err)
}

// Map25 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map25[T1, T2, U1, U2, U3, U4, U5 any](self Result2[T1, T2], op func(T1, T2) (U1, U2, U3, U4, U5)) Result5[U1, U2, U3, U4, U5] {
	if self.err != nil {
		return Err5[U1, U2, U3, U4, U5](self.err)
	}
	return Ok5[U1, U2, U3, U4, U5](op(self.Value()))
}

// MapOr25 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr25[T1, T2, U1, U2, U3, U4, U5 any](self Result2[T1, T2], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, op func(T1, T2) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if self.err != nil {
		return Ok5[U1, U2, U3, U4, U5](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5)
	}
	return op(self.Value())
}

// And26 Returns res if the result is Ok, otherwise returns the Err value of self
func And26[T1, T2, U1, U2, U3, U4, U5, U6 any](self Result2[T1, T2], res Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if self.IsOk() {
		return res
	}
	return Err6[U1, U2, U3, U4, U5, U6](self.err)
}

// AndThen26 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen26[T1, T2, U1, U2, U3, U4, U5, U6 any](self Result2[T1, T2], op func(T1, T2) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err6[U1, U2, U3, U4, U5, U6](self.err)
}

// Map26 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map26[T1, T2, U1, U2, U3, U4, U5, U6 any](self Result2[T1, T2], op func(T1, T2) (U1, U2, U3, U4, U5, U6)) Result6[U1, U2, U3, U4, U5, U6] {
	if self.err != nil {
		return Err6[U1, U2, U3, U4, U5, U6](self.err)
	}
	return Ok6[U1, U2, U3, U4, U5, U6](op(self.Value()))
}

// MapOr26 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr26[T1, T2, U1, U2, U3, U4, U5, U6 any](self Result2[T1, T2], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, defaultV6 U6, op func(T1, T2) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if self.err != nil {
		return Ok6[U1, U2, U3, U4, U5, U6](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5, defaultV6)
	}
	return op(self.Value())
}

// And31 Returns res if the result is Ok, otherwise returns the Err value of self
func And31[T1, T2, T3, U1 any](self Result3[T1, T2, T3], res Result1[U1]) Result1[U1] {
	if self.IsOk() {
		return res
	}
	return Err1[U1](self.err)
}

// AndThen31 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen31[T1, T2, T3, U1 any](self Result3[T1, T2, T3], op func(T1, T2, T3) Result1[U1]) Result1[U1] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err1[U1](self.err)
}

// Map31 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map31[T1, T2, T3, U1 any](self Result3[T1, T2, T3], op func(T1, T2, T3) U1) Result1[U1] {
	if self.err != nil {
		return Err1[U1](self.err)
	}
	return Ok1[U1](op(self.Value()))
}

// MapOr31 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr31[T1, T2, T3, U1 any](self Result3[T1, T2, T3], defaultV1 U1, op func(T1, T2, T3) Result1[U1]) Result1[U1] {
	if self.err != nil {
		return Ok1[U1](defaultV1)
	}
	return op(self.Value())
}

// And32 Returns res if the result is Ok, otherwise returns the Err value of self
func And32[T1, T2, T3, U1, U2 any](self Result3[T1, T2, T3], res Result2[U1, U2]) Result2[U1, U2] {
	if self.IsOk() {
		return res
	}
	return Err2[U1, U2](self.err)
}

// AndThen32 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen32[T1, T2, T3, U1, U2 any](self Result3[T1, T2, T3], op func(T1, T2, T3) Result2[U1, U2]) Result2[U1, U2] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err2[U1, U2](self.err)
}

// Map32 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map32[T1, T2, T3, U1, U2 any](self Result3[T1, T2, T3], op func(T1, T2, T3) (U1, U2)) Result2[U1, U2] {
	if self.err != nil {
		return Err2[U1, U2](self.err)
	}
	return Ok2[U1, U2](op(self.Value()))
}

// MapOr32 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr32[T1, T2, T3, U1, U2 any](self Result3[T1, T2, T3], defaultV1 U1, defaultV2 U2, op func(T1, T2, T3) Result2[U1, U2]) Result2[U1, U2] {
	if self.err != nil {
		return Ok2[U1, U2](defaultV1, defaultV2)
	}
	return op(self.Value())
}

// And33 Returns res if the result is Ok, otherwise returns the Err value of self
func And33[T1, T2, T3, U1, U2, U3 any](self Result3[T1, T2, T3], res Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if self.IsOk() {
		return res
	}
	return Err3[U1, U2, U3](self.err)
}

// AndThen33 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen33[T1, T2, T3, U1, U2, U3 any](self Result3[T1, T2, T3], op func(T1, T2, T3) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err3[U1, U2, U3](self.err)
}

// Map33 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map33[T1, T2, T3, U1, U2, U3 any](self Result3[T1, T2, T3], op func(T1, T2, T3) (U1, U2, U3)) Result3[U1, U2, U3] {
	if self.err != nil {
		return Err3[U1, U2, U3](self.err)
	}
	return Ok3[U1, U2, U3](op(self.Value()))
}

// MapOr33 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr33[T1, T2, T3, U1, U2, U3 any](self Result3[T1, T2, T3], defaultV1 U1, defaultV2 U2, defaultV3 U3, op func(T1, T2, T3) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if self.err != nil {
		return Ok3[U1, U2, U3](defaultV1, defaultV2, defaultV3)
	}
	return op(self.Value())
}

// And34 Returns res if the result is Ok, otherwise returns the Err value of self
func And34[T1, T2, T3, U1, U2, U3, U4 any](self Result3[T1, T2, T3], res Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if self.IsOk() {
		return res
	}
	return Err4[U1, U2, U3, U4](self.err)
}

// AndThen34 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen34[T1, T2, T3, U1, U2, U3, U4 any](self Result3[T1, T2, T3], op func(T1, T2, T3) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err4[U1, U2, U3, U4](self.err)
}

// Map34 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map34[T1, T2, T3, U1, U2, U3, U4 any](self Result3[T1, T2, T3], op func(T1, T2, T3) (U1, U2, U3, U4)) Result4[U1, U2, U3, U4] {
	if self.err != nil {
		return Err4[U1, U2, U3, U4](self.err)
	}
	return Ok4[U1, U2, U3, U4](op(self.Value()))
}

// MapOr34 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr34[T1, T2, T3, U1, U2, U3, U4 any](self Result3[T1, T2, T3], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, op func(T1, T2, T3) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if self.err != nil {
		return Ok4[U1, U2, U3, U4](defaultV1, defaultV2, defaultV3, defaultV4)
	}
	return op(self.Value())
}

// And35 Returns res if the result is Ok, otherwise returns the Err value of self
func And35[T1, T2, T3, U1, U2, U3, U4, U5 any](self Result3[T1, T2, T3], res Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if self.IsOk() {
		return res
	}
	return Err5[U1, U2, U3, U4, U5](self.err)
}

// AndThen35 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen35[T1, T2, T3, U1, U2, U3, U4, U5 any](self Result3[T1, T2, T3], op func(T1, T2, T3) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err5[U1, U2, U3, U4, U5](self.err)
}

// Map35 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map35[T1, T2, T3, U1, U2, U3, U4, U5 any](self Result3[T1, T2, T3], op func(T1, T2, T3) (U1, U2, U3, U4, U5)) Result5[U1, U2, U3, U4, U5] {
	if self.err != nil {
		return Err5[U1, U2, U3, U4, U5](self.err)
	}
	return Ok5[U1, U2, U3, U4, U5](op(self.Value()))
}

// MapOr35 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr35[T1, T2, T3, U1, U2, U3, U4, U5 any](self Result3[T1, T2, T3], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, op func(T1, T2, T3) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if self.err != nil {
		return Ok5[U1, U2, U3, U4, U5](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5)
	}
	return op(self.Value())
}

// And36 Returns res if the result is Ok, otherwise returns the Err value of self
func And36[T1, T2, T3, U1, U2, U3, U4, U5, U6 any](self Result3[T1, T2, T3], res Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if self.IsOk() {
		return res
	}
	return Err6[U1, U2, U3, U4, U5, U6](self.err)
}

// AndThen36 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen36[T1, T2, T3, U1, U2, U3, U4, U5, U6 any](self Result3[T1, T2, T3], op func(T1, T2, T3) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err6[U1, U2, U3, U4, U5, U6](self.err)
}

// Map36 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map36[T1, T2, T3, U1, U2, U3, U4, U5, U6 any](self Result3[T1, T2, T3], op func(T1, T2, T3) (U1, U2, U3, U4, U5, U6)) Result6[U1, U2, U3, U4, U5, U6] {
	if self.err != nil {
		return Err6[U1, U2, U3, U4, U5, U6](self.err)
	}
	return Ok6[U1, U2, U3, U4, U5, U6](op(self.Value()))
}

// MapOr36 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr36[T1, T2, T3, U1, U2, U3, U4, U5, U6 any](self Result3[T1, T2, T3], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, defaultV6 U6, op func(T1, T2, T3) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if self.err != nil {
		return Ok6[U1, U2, U3, U4, U5, U6](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5, defaultV6)
	}
	return op(self.Value())
}

// And41 Returns res if the result is Ok, otherwise returns the Err value of self
func And41[T1, T2, T3, T4, U1 any](self Result4[T1, T2, T3, T4], res Result1[U1]) Result1[U1] {
	if self.IsOk() {
		return res
	}
	return Err1[U1](self.err)
}

// AndThen41 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen41[T1, T2, T3, T4, U1 any](self Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) Result1[U1]) Result1[U1] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err1[U1](self.err)
}

// Map41 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map41[T1, T2, T3, T4, U1 any](self Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) U1) Result1[U1] {
	if self.err != nil {
		return Err1[U1](self.err)
	}
	return Ok1[U1](op(self.Value()))
}

// MapOr41 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr41[T1, T2, T3, T4, U1 any](self Result4[T1, T2, T3, T4], defaultV1 U1, op func(T1, T2, T3, T4) Result1[U1]) Result1[U1] {
	if self.err != nil {
		return Ok1[U1](defaultV1)
	}
	return op(self.Value())
}

// And42 Returns res if the result is Ok, otherwise returns the Err value of self
func And42[T1, T2, T3, T4, U1, U2 any](self Result4[T1, T2, T3, T4], res Result2[U1, U2]) Result2[U1, U2] {
	if self.IsOk() {
		return res
	}
	return Err2[U1, U2](self.err)
}

// AndThen42 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen42[T1, T2, T3, T4, U1, U2 any](self Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) Result2[U1, U2]) Result2[U1, U2] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err2[U1, U2](self.err)
}

// Map42 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map42[T1, T2, T3, T4, U1, U2 any](self Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) (U1, U2)) Result2[U1, U2] {
	if self.err != nil {
		return Err2[U1, U2](self.err)
	}
	return Ok2[U1, U2](op(self.Value()))
}

// MapOr42 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr42[T1, T2, T3, T4, U1, U2 any](self Result4[T1, T2, T3, T4], defaultV1 U1, defaultV2 U2, op func(T1, T2, T3, T4) Result2[U1, U2]) Result2[U1, U2] {
	if self.err != nil {
		return Ok2[U1, U2](defaultV1, defaultV2)
	}
	return op(self.Value())
}

// And43 Returns res if the result is Ok, otherwise returns the Err value of self
func And43[T1, T2, T3, T4, U1, U2, U3 any](self Result4[T1, T2, T3, T4], res Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if self.IsOk() {
		return res
	}
	return Err3[U1, U2, U3](self.err)
}

// AndThen43 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen43[T1, T2, T3, T4, U1, U2, U3 any](self Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err3[U1, U2, U3](self.err)
}

// Map43 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map43[T1, T2, T3, T4, U1, U2, U3 any](self Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) (U1, U2, U3)) Result3[U1, U2, U3] {
	if self.err != nil {
		return Err3[U1, U2, U3](self.err)
	}
	return Ok3[U1, U2, U3](op(self.Value()))
}

// MapOr43 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr43[T1, T2, T3, T4, U1, U2, U3 any](self Result4[T1, T2, T3, T4], defaultV1 U1, defaultV2 U2, defaultV3 U3, op func(T1, T2, T3, T4) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if self.err != nil {
		return Ok3[U1, U2, U3](defaultV1, defaultV2, defaultV3)
	}
	return op(self.Value())
}

// And44 Returns res if the result is Ok, otherwise returns the Err value of self
func And44[T1, T2, T3, T4, U1, U2, U3, U4 any](self Result4[T1, T2, T3, T4], res Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if self.IsOk() {
		return res
	}
	return Err4[U1, U2, U3, U4](self.err)
}

// AndThen44 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen44[T1, T2, T3, T4, U1, U2, U3, U4 any](self Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err4[U1, U2, U3, U4](self.err)
}

// Map44 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map44[T1, T2, T3, T4, U1, U2, U3, U4 any](self Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) (U1, U2, U3, U4)) Result4[U1, U2, U3, U4] {
	if self.err != nil {
		return Err4[U1, U2, U3, U4](self.err)
	}
	return Ok4[U1, U2, U3, U4](op(self.Value()))
}

// MapOr44 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr44[T1, T2, T3, T4, U1, U2, U3, U4 any](self Result4[T1, T2, T3, T4], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, op func(T1, T2, T3, T4) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if self.err != nil {
		return Ok4[U1, U2, U3, U4](defaultV1, defaultV2, defaultV3, defaultV4)
	}
	return op(self.Value())
}

// And45 Returns res if the result is Ok, otherwise returns the Err value of self
func And45[T1, T2, T3, T4, U1, U2, U3, U4, U5 any](self Result4[T1, T2, T3, T4], res Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if self.IsOk() {
		return res
	}
	return Err5[U1, U2, U3, U4, U5](self.err)
}

// AndThen45 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen45[T1, T2, T3, T4, U1, U2, U3, U4, U5 any](self Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err5[U1, U2, U3, U4, U5](self.err)
}

// Map45 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map45[T1, T2, T3, T4, U1, U2, U3, U4, U5 any](self Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) (U1, U2, U3, U4, U5)) Result5[U1, U2, U3, U4, U5] {
	if self.err != nil {
		return Err5[U1, U2, U3, U4, U5](self.err)
	}
	return Ok5[U1, U2, U3, U4, U5](op(self.Value()))
}

// MapOr45 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr45[T1, T2, T3, T4, U1, U2, U3, U4, U5 any](self Result4[T1, T2, T3, T4], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, op func(T1, T2, T3, T4) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if self.err != nil {
		return Ok5[U1, U2, U3, U4, U5](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5)
	}
	return op(self.Value())
}

// And46 Returns res if the result is Ok, otherwise returns the Err value of self
func And46[T1, T2, T3, T4, U1, U2, U3, U4, U5, U6 any](self Result4[T1, T2, T3, T4], res Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if self.IsOk() {
		return res
	}
	return Err6[U1, U2, U3, U4, U5, U6](self.err)
}

// AndThen46 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen46[T1, T2, T3, T4, U1, U2, U3, U4, U5, U6 any](self Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err6[U1, U2, U3, U4, U5, U6](self.err)
}

// Map46 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map46[T1, T2, T3, T4, U1, U2, U3, U4, U5, U6 any](self Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) (U1, U2, U3, U4, U5, U6)) Result6[U1, U2, U3, U4, U5, U6] {
	if self.err != nil {
		return Err6[U1, U2, U3, U4, U5, U6](self.err)
	}
	return Ok6[U1, U2, U3, U4, U5, U6](op(self.Value()))
}

// MapOr46 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr46[T1, T2, T3, T4, U1, U2, U3, U4, U5, U6 any](self Result4[T1, T2, T3, T4], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, defaultV6 U6, op func(T1, T2, T3, T4) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if self.err != nil {
		return Ok6[U1, U2, U3, U4, U5, U6](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5, defaultV6)
	}
	return op(self.Value())
}

// And51 Returns res if the result is Ok, otherwise returns the Err value of self
func And51[T1, T2, T3, T4, T5, U1 any](self Result5[T1, T2, T3, T4, T5], res Result1[U1]) Result1[U1] {
	if self.IsOk() {
		return res
	}
	return Err1[U1](self.err)
}

// AndThen51 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen51[T1, T2, T3, T4, T5, U1 any](self Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) Result1[U1]) Result1[U1] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err1[U1](self.err)
}

// Map51 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map51[T1, T2, T3, T4, T5, U1 any](self Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) U1) Result1[U1] {
	if self.err != nil {
		return Err1[U1](self.err)
	}
	return Ok1[U1](op(self.Value()))
}

// MapOr51 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr51[T1, T2, T3, T4, T5, U1 any](self Result5[T1, T2, T3, T4, T5], defaultV1 U1, op func(T1, T2, T3, T4, T5) Result1[U1]) Result1[U1] {
	if self.err != nil {
		return Ok1[U1](defaultV1)
	}
	return op(self.Value())
}

// And52 Returns res if the result is Ok, otherwise returns the Err value of self
func And52[T1, T2, T3, T4, T5, U1, U2 any](self Result5[T1, T2, T3, T4, T5], res Result2[U1, U2]) Result2[U1, U2] {
	if self.IsOk() {
		return res
	}
	return Err2[U1, U2](self.err)
}

// AndThen52 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen52[T1, T2, T3, T4, T5, U1, U2 any](self Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) Result2[U1, U2]) Result2[U1, U2] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err2[U1, U2](self.err)
}

// Map52 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map52[T1, T2, T3, T4, T5, U1, U2 any](self Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) (U1, U2)) Result2[U1, U2] {
	if self.err != nil {
		return Err2[U1, U2](self.err)
	}
	return Ok2[U1, U2](op(self.Value()))
}

// MapOr52 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr52[T1, T2, T3, T4, T5, U1, U2 any](self Result5[T1, T2, T3, T4, T5], defaultV1 U1, defaultV2 U2, op func(T1, T2, T3, T4, T5) Result2[U1, U2]) Result2[U1, U2] {
	if self.err != nil {
		return Ok2[U1, U2](defaultV1, defaultV2)
	}
	return op(self.Value())
}

// And53 Returns res if the result is Ok, otherwise returns the Err value of self
func And53[T1, T2, T3, T4, T5, U1, U2, U3 any](self Result5[T1, T2, T3, T4, T5], res Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if self.IsOk() {
		return res
	}
	return Err3[U1, U2, U3](self.err)
}

// AndThen53 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen53[T1, T2, T3, T4, T5, U1, U2, U3 any](self Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err3[U1, U2, U3](self.err)
}

// Map53 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map53[T1, T2, T3, T4, T5, U1, U2, U3 any](self Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) (U1, U2, U3)) Result3[U1, U2, U3] {
	if self.err != nil {
		return Err3[U1, U2, U3](self.err)
	}
	return Ok3[U1, U2, U3](op(self.Value()))
}

// MapOr53 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr53[T1, T2, T3, T4, T5, U1, U2, U3 any](self Result5[T1, T2, T3, T4, T5], defaultV1 U1, defaultV2 U2, defaultV3 U3, op func(T1, T2, T3, T4, T5) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if self.err != nil {
		return Ok3[U1, U2, U3](defaultV1, defaultV2, defaultV3)
	}
	return op(self.Value())
}

// And54 Returns res if the result is Ok, otherwise returns the Err value of self
func And54[T1, T2, T3, T4, T5, U1, U2, U3, U4 any](self Result5[T1, T2, T3, T4, T5], res Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if self.IsOk() {
		return res
	}
	return Err4[U1, U2, U3, U4](self.err)
}

// AndThen54 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen54[T1, T2, T3, T4, T5, U1, U2, U3, U4 any](self Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err4[U1, U2, U3, U4](self.err)
}

// Map54 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map54[T1, T2, T3, T4, T5, U1, U2, U3, U4 any](self Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) (U1, U2, U3, U4)) Result4[U1, U2, U3, U4] {
	if self.err != nil {
		return Err4[U1, U2, U3, U4](self.err)
	}
	return Ok4[U1, U2, U3, U4](op(self.Value()))
}

// MapOr54 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr54[T1, T2, T3, T4, T5, U1, U2, U3, U4 any](self Result5[T1, T2, T3, T4, T5], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, op func(T1, T2, T3, T4, T5) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if self.err != nil {
		return Ok4[U1, U2, U3, U4](defaultV1, defaultV2, defaultV3, defaultV4)
	}
	return op(self.Value())
}

// And55 Returns res if the result is Ok, otherwise returns the Err value of self
func And55[T1, T2, T3, T4, T5, U1, U2, U3, U4, U5 any](self Result5[T1, T2, T3, T4, T5], res Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if self.IsOk() {
		return res
	}
	return Err5[U1, U2, U3, U4, U5](self.err)
}

// AndThen55 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen55[T1, T2, T3, T4, T5, U1, U2, U3, U4, U5 any](self Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err5[U1, U2, U3, U4, U5](self.err)
}

// Map55 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map55[T1, T2, T3, T4, T5, U1, U2, U3, U4, U5 any](self Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) (U1, U2, U3, U4, U5)) Result5[U1, U2, U3, U4, U5] {
	if self.err != nil {
		return Err5[U1, U2, U3, U4, U5](self.err)
	}
	return Ok5[U1, U2, U3, U4, U5](op(self.Value()))
}

// MapOr55 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr55[T1, T2, T3, T4, T5, U1, U2, U3, U4, U5 any](self Result5[T1, T2, T3, T4, T5], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, op func(T1, T2, T3, T4, T5) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if self.err != nil {
		return Ok5[U1, U2, U3, U4, U5](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5)
	}
	return op(self.Value())
}

// And56 Returns res if the result is Ok, otherwise returns the Err value of self
func And56[T1, T2, T3, T4, T5, U1, U2, U3, U4, U5, U6 any](self Result5[T1, T2, T3, T4, T5], res Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if self.IsOk() {
		return res
	}
	return Err6[U1, U2, U3, U4, U5, U6](self.err)
}

// AndThen56 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen56[T1, T2, T3, T4, T5, U1, U2, U3, U4, U5, U6 any](self Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err6[U1, U2, U3, U4, U5, U6](self.err)
}

// Map56 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map56[T1, T2, T3, T4, T5, U1, U2, U3, U4, U5, U6 any](self Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) (U1, U2, U3, U4, U5, U6)) Result6[U1, U2, U3, U4, U5, U6] {
	if self.err != nil {
		return Err6[U1, U2, U3, U4, U5, U6](self.err)
	}
	return Ok6[U1, U2, U3, U4, U5, U6](op(self.Value()))
}

// MapOr56 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr56[T1, T2, T3, T4, T5, U1, U2, U3, U4, U5, U6 any](self Result5[T1, T2, T3, T4, T5], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, defaultV6 U6, op func(T1, T2, T3, T4, T5) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if self.err != nil {
		return Ok6[U1, U2, U3, U4, U5, U6](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5, defaultV6)
	}
	return op(self.Value())
}

// And61 Returns res if the result is Ok, otherwise returns the Err value of self
func And61[T1, T2, T3, T4, T5, T6, U1 any](self Result6[T1, T2, T3, T4, T5, T6], res Result1[U1]) Result1[U1] {
	if self.IsOk() {
		return res
	}
	return Err1[U1](self.err)
}

// AndThen61 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen61[T1, T2, T3, T4, T5, T6, U1 any](self Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) Result1[U1]) Result1[U1] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err1[U1](self.err)
}

// Map61 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map61[T1, T2, T3, T4, T5, T6, U1 any](self Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) U1) Result1[U1] {
	if self.err != nil {
		return Err1[U1](self.err)
	}
	return Ok1[U1](op(self.Value()))
}

// MapOr61 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr61[T1, T2, T3, T4, T5, T6, U1 any](self Result6[T1, T2, T3, T4, T5, T6], defaultV1 U1, op func(T1, T2, T3, T4, T5, T6) Result1[U1]) Result1[U1] {
	if self.err != nil {
		return Ok1[U1](defaultV1)
	}
	return op(self.Value())
}

// And62 Returns res if the result is Ok, otherwise returns the Err value of self
func And62[T1, T2, T3, T4, T5, T6, U1, U2 any](self Result6[T1, T2, T3, T4, T5, T6], res Result2[U1, U2]) Result2[U1, U2] {
	if self.IsOk() {
		return res
	}
	return Err2[U1, U2](self.err)
}

// AndThen62 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen62[T1, T2, T3, T4, T5, T6, U1, U2 any](self Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) Result2[U1, U2]) Result2[U1, U2] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err2[U1, U2](self.err)
}

// Map62 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map62[T1, T2, T3, T4, T5, T6, U1, U2 any](self Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) (U1, U2)) Result2[U1, U2] {
	if self.err != nil {
		return Err2[U1, U2](self.err)
	}
	return Ok2[U1, U2](op(self.Value()))
}

// MapOr62 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr62[T1, T2, T3, T4, T5, T6, U1, U2 any](self Result6[T1, T2, T3, T4, T5, T6], defaultV1 U1, defaultV2 U2, op func(T1, T2, T3, T4, T5, T6) Result2[U1, U2]) Result2[U1, U2] {
	if self.err != nil {
		return Ok2[U1, U2](defaultV1, defaultV2)
	}
	return op(self.Value())
}

// And63 Returns res if the result is Ok, otherwise returns the Err value of self
func And63[T1, T2, T3, T4, T5, T6, U1, U2, U3 any](self Result6[T1, T2, T3, T4, T5, T6], res Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if self.IsOk() {
		return res
	}
	return Err3[U1, U2, U3](self.err)
}

// AndThen63 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen63[T1, T2, T3, T4, T5, T6, U1, U2, U3 any](self Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err3[U1, U2, U3](self.err)
}

// Map63 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map63[T1, T2, T3, T4, T5, T6, U1, U2, U3 any](self Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) (U1, U2, U3)) Result3[U1, U2, U3] {
	if self.err != nil {
		return Err3[U1, U2, U3](self.err)
	}
	return Ok3[U1, U2, U3](op(self.Value()))
}

// MapOr63 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr63[T1, T2, T3, T4, T5, T6, U1, U2, U3 any](self Result6[T1, T2, T3, T4, T5, T6], defaultV1 U1, defaultV2 U2, defaultV3 U3, op func(T1, T2, T3, T4, T5, T6) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if self.err != nil {
		return Ok3[U1, U2, U3](defaultV1, defaultV2, defaultV3)
	}
	return op(self.Value())
}

// And64 Returns res if the result is Ok, otherwise returns the Err value of self
func And64[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4 any](self Result6[T1, T2, T3, T4, T5, T6], res Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if self.IsOk() {
		return res
	}
	return Err4[U1, U2, U3, U4](self.err)
}

// AndThen64 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen64[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4 any](self Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err4[U1, U2, U3, U4](self.err)
}

// Map64 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map64[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4 any](self Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) (U1, U2, U3, U4)) Result4[U1, U2, U3, U4] {
	if self.err != nil {
		return Err4[U1, U2, U3, U4](self.err)
	}
	return Ok4[U1, U2, U3, U4](op(self.Value()))
}

// MapOr64 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr64[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4 any](self Result6[T1, T2, T3, T4, T5, T6], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, op func(T1, T2, T3, T4, T5, T6) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if self.err != nil {
		return Ok4[U1, U2, U3, U4](defaultV1, defaultV2, defaultV3, defaultV4)
	}
	return op(self.Value())
}

// And65 Returns res if the result is Ok, otherwise returns the Err value of self
func And65[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4, U5 any](self Result6[T1, T2, T3, T4, T5, T6], res Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if self.IsOk() {
		return res
	}
	return Err5[U1, U2, U3, U4, U5](self.err)
}

// AndThen65 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen65[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4, U5 any](self Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err5[U1, U2, U3, U4, U5](self.err)
}

// Map65 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map65[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4, U5 any](self Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) (U1, U2, U3, U4, U5)) Result5[U1, U2, U3, U4, U5] {
	if self.err != nil {
		return Err5[U1, U2, U3, U4, U5](self.err)
	}
	return Ok5[U1, U2, U3, U4, U5](op(self.Value()))
}

// MapOr65 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr65[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4, U5 any](self Result6[T1, T2, T3, T4, T5, T6], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, op func(T1, T2, T3, T4, T5, T6) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if self.err != nil {
		return Ok5[U1, U2, U3, U4, U5](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5)
	}
	return op(self.Value())
}

// And66 Returns res if the result is Ok, otherwise returns the Err value of self
func And66[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4, U5, U6 any](self Result6[T1, T2, T3, T4, T5, T6], res Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if self.IsOk() {
		return res
	}
	return Err6[U1, U2, U3, U4, U5, U6](self.err)
}

// AndThen66 Calls op if the result is Ok, otherwise returns the Err value of self
func AndThen66[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4, U5, U6 any](self Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if self.IsOk() {
		return op(self.Value())
	}
	return Err6[U1, U2, U3, U4, U5, U6](self.err)
}

// Map66 Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
func Map66[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4, U5, U6 any](self Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) (U1, U2, U3, U4, U5, U6)) Result6[U1, U2, U3, U4, U5, U6] {
	if self.err != nil {
		return Err6[U1, U2, U3, U4, U5, U6](self.err)
	}
	return Ok6[U1, U2, U3, U4, U5, U6](op(self.Value()))
}

// MapOr66 Returns the provided default (if Err), or applies a function to the contained value (if Ok)
func MapOr66[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4, U5, U6 any](self Result6[T1, T2, T3, T4, T5, T6], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, defaultV6 U6, op func(T1, T2, T3, T4, T5, T6) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if self.err != nil {
		return Ok6[U1, U2, U3, U4, U5, U6](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5, defaultV6)
	}
	return op(self.Value())
}
