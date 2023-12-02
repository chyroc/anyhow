package anyhow

// Then11 ...
func Then11[T1, U1 any](r Result1[T1], op func(T1) Result1[U1]) Result1[U1] {
	if r.err != nil {
		return Err1[U1](r.err)
	}
	return op(r.Value())
}

// Or11 ...
func Or11[T1, U1 any](r Result1[T1], defaultV1 U1, op func(T1) Result1[U1]) Result1[U1] {
	if r.err != nil {
		return Ok1[U1](defaultV1)
	}
	return op(r.Value())
}

// Then12 ...
func Then12[T1, U1, U2 any](r Result1[T1], op func(T1) Result2[U1, U2]) Result2[U1, U2] {
	if r.err != nil {
		return Err2[U1, U2](r.err)
	}
	return op(r.Value())
}

// Or12 ...
func Or12[T1, U1, U2 any](r Result1[T1], defaultV1 U1, defaultV2 U2, op func(T1) Result2[U1, U2]) Result2[U1, U2] {
	if r.err != nil {
		return Ok2[U1, U2](defaultV1, defaultV2)
	}
	return op(r.Value())
}

// Then13 ...
func Then13[T1, U1, U2, U3 any](r Result1[T1], op func(T1) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if r.err != nil {
		return Err3[U1, U2, U3](r.err)
	}
	return op(r.Value())
}

// Or13 ...
func Or13[T1, U1, U2, U3 any](r Result1[T1], defaultV1 U1, defaultV2 U2, defaultV3 U3, op func(T1) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if r.err != nil {
		return Ok3[U1, U2, U3](defaultV1, defaultV2, defaultV3)
	}
	return op(r.Value())
}

// Then14 ...
func Then14[T1, U1, U2, U3, U4 any](r Result1[T1], op func(T1) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if r.err != nil {
		return Err4[U1, U2, U3, U4](r.err)
	}
	return op(r.Value())
}

// Or14 ...
func Or14[T1, U1, U2, U3, U4 any](r Result1[T1], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, op func(T1) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if r.err != nil {
		return Ok4[U1, U2, U3, U4](defaultV1, defaultV2, defaultV3, defaultV4)
	}
	return op(r.Value())
}

// Then15 ...
func Then15[T1, U1, U2, U3, U4, U5 any](r Result1[T1], op func(T1) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if r.err != nil {
		return Err5[U1, U2, U3, U4, U5](r.err)
	}
	return op(r.Value())
}

// Or15 ...
func Or15[T1, U1, U2, U3, U4, U5 any](r Result1[T1], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, op func(T1) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if r.err != nil {
		return Ok5[U1, U2, U3, U4, U5](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5)
	}
	return op(r.Value())
}

// Then16 ...
func Then16[T1, U1, U2, U3, U4, U5, U6 any](r Result1[T1], op func(T1) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if r.err != nil {
		return Err6[U1, U2, U3, U4, U5, U6](r.err)
	}
	return op(r.Value())
}

// Or16 ...
func Or16[T1, U1, U2, U3, U4, U5, U6 any](r Result1[T1], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, defaultV6 U6, op func(T1) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if r.err != nil {
		return Ok6[U1, U2, U3, U4, U5, U6](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5, defaultV6)
	}
	return op(r.Value())
}

// Then21 ...
func Then21[T1, T2, U1 any](r Result2[T1, T2], op func(T1, T2) Result1[U1]) Result1[U1] {
	if r.err != nil {
		return Err1[U1](r.err)
	}
	return op(r.Value())
}

// Or21 ...
func Or21[T1, T2, U1 any](r Result2[T1, T2], defaultV1 U1, op func(T1, T2) Result1[U1]) Result1[U1] {
	if r.err != nil {
		return Ok1[U1](defaultV1)
	}
	return op(r.Value())
}

// Then22 ...
func Then22[T1, T2, U1, U2 any](r Result2[T1, T2], op func(T1, T2) Result2[U1, U2]) Result2[U1, U2] {
	if r.err != nil {
		return Err2[U1, U2](r.err)
	}
	return op(r.Value())
}

// Or22 ...
func Or22[T1, T2, U1, U2 any](r Result2[T1, T2], defaultV1 U1, defaultV2 U2, op func(T1, T2) Result2[U1, U2]) Result2[U1, U2] {
	if r.err != nil {
		return Ok2[U1, U2](defaultV1, defaultV2)
	}
	return op(r.Value())
}

// Then23 ...
func Then23[T1, T2, U1, U2, U3 any](r Result2[T1, T2], op func(T1, T2) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if r.err != nil {
		return Err3[U1, U2, U3](r.err)
	}
	return op(r.Value())
}

// Or23 ...
func Or23[T1, T2, U1, U2, U3 any](r Result2[T1, T2], defaultV1 U1, defaultV2 U2, defaultV3 U3, op func(T1, T2) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if r.err != nil {
		return Ok3[U1, U2, U3](defaultV1, defaultV2, defaultV3)
	}
	return op(r.Value())
}

// Then24 ...
func Then24[T1, T2, U1, U2, U3, U4 any](r Result2[T1, T2], op func(T1, T2) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if r.err != nil {
		return Err4[U1, U2, U3, U4](r.err)
	}
	return op(r.Value())
}

// Or24 ...
func Or24[T1, T2, U1, U2, U3, U4 any](r Result2[T1, T2], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, op func(T1, T2) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if r.err != nil {
		return Ok4[U1, U2, U3, U4](defaultV1, defaultV2, defaultV3, defaultV4)
	}
	return op(r.Value())
}

// Then25 ...
func Then25[T1, T2, U1, U2, U3, U4, U5 any](r Result2[T1, T2], op func(T1, T2) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if r.err != nil {
		return Err5[U1, U2, U3, U4, U5](r.err)
	}
	return op(r.Value())
}

// Or25 ...
func Or25[T1, T2, U1, U2, U3, U4, U5 any](r Result2[T1, T2], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, op func(T1, T2) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if r.err != nil {
		return Ok5[U1, U2, U3, U4, U5](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5)
	}
	return op(r.Value())
}

// Then26 ...
func Then26[T1, T2, U1, U2, U3, U4, U5, U6 any](r Result2[T1, T2], op func(T1, T2) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if r.err != nil {
		return Err6[U1, U2, U3, U4, U5, U6](r.err)
	}
	return op(r.Value())
}

// Or26 ...
func Or26[T1, T2, U1, U2, U3, U4, U5, U6 any](r Result2[T1, T2], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, defaultV6 U6, op func(T1, T2) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if r.err != nil {
		return Ok6[U1, U2, U3, U4, U5, U6](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5, defaultV6)
	}
	return op(r.Value())
}

// Then31 ...
func Then31[T1, T2, T3, U1 any](r Result3[T1, T2, T3], op func(T1, T2, T3) Result1[U1]) Result1[U1] {
	if r.err != nil {
		return Err1[U1](r.err)
	}
	return op(r.Value())
}

// Or31 ...
func Or31[T1, T2, T3, U1 any](r Result3[T1, T2, T3], defaultV1 U1, op func(T1, T2, T3) Result1[U1]) Result1[U1] {
	if r.err != nil {
		return Ok1[U1](defaultV1)
	}
	return op(r.Value())
}

// Then32 ...
func Then32[T1, T2, T3, U1, U2 any](r Result3[T1, T2, T3], op func(T1, T2, T3) Result2[U1, U2]) Result2[U1, U2] {
	if r.err != nil {
		return Err2[U1, U2](r.err)
	}
	return op(r.Value())
}

// Or32 ...
func Or32[T1, T2, T3, U1, U2 any](r Result3[T1, T2, T3], defaultV1 U1, defaultV2 U2, op func(T1, T2, T3) Result2[U1, U2]) Result2[U1, U2] {
	if r.err != nil {
		return Ok2[U1, U2](defaultV1, defaultV2)
	}
	return op(r.Value())
}

// Then33 ...
func Then33[T1, T2, T3, U1, U2, U3 any](r Result3[T1, T2, T3], op func(T1, T2, T3) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if r.err != nil {
		return Err3[U1, U2, U3](r.err)
	}
	return op(r.Value())
}

// Or33 ...
func Or33[T1, T2, T3, U1, U2, U3 any](r Result3[T1, T2, T3], defaultV1 U1, defaultV2 U2, defaultV3 U3, op func(T1, T2, T3) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if r.err != nil {
		return Ok3[U1, U2, U3](defaultV1, defaultV2, defaultV3)
	}
	return op(r.Value())
}

// Then34 ...
func Then34[T1, T2, T3, U1, U2, U3, U4 any](r Result3[T1, T2, T3], op func(T1, T2, T3) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if r.err != nil {
		return Err4[U1, U2, U3, U4](r.err)
	}
	return op(r.Value())
}

// Or34 ...
func Or34[T1, T2, T3, U1, U2, U3, U4 any](r Result3[T1, T2, T3], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, op func(T1, T2, T3) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if r.err != nil {
		return Ok4[U1, U2, U3, U4](defaultV1, defaultV2, defaultV3, defaultV4)
	}
	return op(r.Value())
}

// Then35 ...
func Then35[T1, T2, T3, U1, U2, U3, U4, U5 any](r Result3[T1, T2, T3], op func(T1, T2, T3) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if r.err != nil {
		return Err5[U1, U2, U3, U4, U5](r.err)
	}
	return op(r.Value())
}

// Or35 ...
func Or35[T1, T2, T3, U1, U2, U3, U4, U5 any](r Result3[T1, T2, T3], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, op func(T1, T2, T3) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if r.err != nil {
		return Ok5[U1, U2, U3, U4, U5](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5)
	}
	return op(r.Value())
}

// Then36 ...
func Then36[T1, T2, T3, U1, U2, U3, U4, U5, U6 any](r Result3[T1, T2, T3], op func(T1, T2, T3) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if r.err != nil {
		return Err6[U1, U2, U3, U4, U5, U6](r.err)
	}
	return op(r.Value())
}

// Or36 ...
func Or36[T1, T2, T3, U1, U2, U3, U4, U5, U6 any](r Result3[T1, T2, T3], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, defaultV6 U6, op func(T1, T2, T3) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if r.err != nil {
		return Ok6[U1, U2, U3, U4, U5, U6](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5, defaultV6)
	}
	return op(r.Value())
}

// Then41 ...
func Then41[T1, T2, T3, T4, U1 any](r Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) Result1[U1]) Result1[U1] {
	if r.err != nil {
		return Err1[U1](r.err)
	}
	return op(r.Value())
}

// Or41 ...
func Or41[T1, T2, T3, T4, U1 any](r Result4[T1, T2, T3, T4], defaultV1 U1, op func(T1, T2, T3, T4) Result1[U1]) Result1[U1] {
	if r.err != nil {
		return Ok1[U1](defaultV1)
	}
	return op(r.Value())
}

// Then42 ...
func Then42[T1, T2, T3, T4, U1, U2 any](r Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) Result2[U1, U2]) Result2[U1, U2] {
	if r.err != nil {
		return Err2[U1, U2](r.err)
	}
	return op(r.Value())
}

// Or42 ...
func Or42[T1, T2, T3, T4, U1, U2 any](r Result4[T1, T2, T3, T4], defaultV1 U1, defaultV2 U2, op func(T1, T2, T3, T4) Result2[U1, U2]) Result2[U1, U2] {
	if r.err != nil {
		return Ok2[U1, U2](defaultV1, defaultV2)
	}
	return op(r.Value())
}

// Then43 ...
func Then43[T1, T2, T3, T4, U1, U2, U3 any](r Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if r.err != nil {
		return Err3[U1, U2, U3](r.err)
	}
	return op(r.Value())
}

// Or43 ...
func Or43[T1, T2, T3, T4, U1, U2, U3 any](r Result4[T1, T2, T3, T4], defaultV1 U1, defaultV2 U2, defaultV3 U3, op func(T1, T2, T3, T4) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if r.err != nil {
		return Ok3[U1, U2, U3](defaultV1, defaultV2, defaultV3)
	}
	return op(r.Value())
}

// Then44 ...
func Then44[T1, T2, T3, T4, U1, U2, U3, U4 any](r Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if r.err != nil {
		return Err4[U1, U2, U3, U4](r.err)
	}
	return op(r.Value())
}

// Or44 ...
func Or44[T1, T2, T3, T4, U1, U2, U3, U4 any](r Result4[T1, T2, T3, T4], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, op func(T1, T2, T3, T4) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if r.err != nil {
		return Ok4[U1, U2, U3, U4](defaultV1, defaultV2, defaultV3, defaultV4)
	}
	return op(r.Value())
}

// Then45 ...
func Then45[T1, T2, T3, T4, U1, U2, U3, U4, U5 any](r Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if r.err != nil {
		return Err5[U1, U2, U3, U4, U5](r.err)
	}
	return op(r.Value())
}

// Or45 ...
func Or45[T1, T2, T3, T4, U1, U2, U3, U4, U5 any](r Result4[T1, T2, T3, T4], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, op func(T1, T2, T3, T4) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if r.err != nil {
		return Ok5[U1, U2, U3, U4, U5](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5)
	}
	return op(r.Value())
}

// Then46 ...
func Then46[T1, T2, T3, T4, U1, U2, U3, U4, U5, U6 any](r Result4[T1, T2, T3, T4], op func(T1, T2, T3, T4) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if r.err != nil {
		return Err6[U1, U2, U3, U4, U5, U6](r.err)
	}
	return op(r.Value())
}

// Or46 ...
func Or46[T1, T2, T3, T4, U1, U2, U3, U4, U5, U6 any](r Result4[T1, T2, T3, T4], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, defaultV6 U6, op func(T1, T2, T3, T4) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if r.err != nil {
		return Ok6[U1, U2, U3, U4, U5, U6](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5, defaultV6)
	}
	return op(r.Value())
}

// Then51 ...
func Then51[T1, T2, T3, T4, T5, U1 any](r Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) Result1[U1]) Result1[U1] {
	if r.err != nil {
		return Err1[U1](r.err)
	}
	return op(r.Value())
}

// Or51 ...
func Or51[T1, T2, T3, T4, T5, U1 any](r Result5[T1, T2, T3, T4, T5], defaultV1 U1, op func(T1, T2, T3, T4, T5) Result1[U1]) Result1[U1] {
	if r.err != nil {
		return Ok1[U1](defaultV1)
	}
	return op(r.Value())
}

// Then52 ...
func Then52[T1, T2, T3, T4, T5, U1, U2 any](r Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) Result2[U1, U2]) Result2[U1, U2] {
	if r.err != nil {
		return Err2[U1, U2](r.err)
	}
	return op(r.Value())
}

// Or52 ...
func Or52[T1, T2, T3, T4, T5, U1, U2 any](r Result5[T1, T2, T3, T4, T5], defaultV1 U1, defaultV2 U2, op func(T1, T2, T3, T4, T5) Result2[U1, U2]) Result2[U1, U2] {
	if r.err != nil {
		return Ok2[U1, U2](defaultV1, defaultV2)
	}
	return op(r.Value())
}

// Then53 ...
func Then53[T1, T2, T3, T4, T5, U1, U2, U3 any](r Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if r.err != nil {
		return Err3[U1, U2, U3](r.err)
	}
	return op(r.Value())
}

// Or53 ...
func Or53[T1, T2, T3, T4, T5, U1, U2, U3 any](r Result5[T1, T2, T3, T4, T5], defaultV1 U1, defaultV2 U2, defaultV3 U3, op func(T1, T2, T3, T4, T5) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if r.err != nil {
		return Ok3[U1, U2, U3](defaultV1, defaultV2, defaultV3)
	}
	return op(r.Value())
}

// Then54 ...
func Then54[T1, T2, T3, T4, T5, U1, U2, U3, U4 any](r Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if r.err != nil {
		return Err4[U1, U2, U3, U4](r.err)
	}
	return op(r.Value())
}

// Or54 ...
func Or54[T1, T2, T3, T4, T5, U1, U2, U3, U4 any](r Result5[T1, T2, T3, T4, T5], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, op func(T1, T2, T3, T4, T5) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if r.err != nil {
		return Ok4[U1, U2, U3, U4](defaultV1, defaultV2, defaultV3, defaultV4)
	}
	return op(r.Value())
}

// Then55 ...
func Then55[T1, T2, T3, T4, T5, U1, U2, U3, U4, U5 any](r Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if r.err != nil {
		return Err5[U1, U2, U3, U4, U5](r.err)
	}
	return op(r.Value())
}

// Or55 ...
func Or55[T1, T2, T3, T4, T5, U1, U2, U3, U4, U5 any](r Result5[T1, T2, T3, T4, T5], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, op func(T1, T2, T3, T4, T5) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if r.err != nil {
		return Ok5[U1, U2, U3, U4, U5](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5)
	}
	return op(r.Value())
}

// Then56 ...
func Then56[T1, T2, T3, T4, T5, U1, U2, U3, U4, U5, U6 any](r Result5[T1, T2, T3, T4, T5], op func(T1, T2, T3, T4, T5) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if r.err != nil {
		return Err6[U1, U2, U3, U4, U5, U6](r.err)
	}
	return op(r.Value())
}

// Or56 ...
func Or56[T1, T2, T3, T4, T5, U1, U2, U3, U4, U5, U6 any](r Result5[T1, T2, T3, T4, T5], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, defaultV6 U6, op func(T1, T2, T3, T4, T5) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if r.err != nil {
		return Ok6[U1, U2, U3, U4, U5, U6](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5, defaultV6)
	}
	return op(r.Value())
}

// Then61 ...
func Then61[T1, T2, T3, T4, T5, T6, U1 any](r Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) Result1[U1]) Result1[U1] {
	if r.err != nil {
		return Err1[U1](r.err)
	}
	return op(r.Value())
}

// Or61 ...
func Or61[T1, T2, T3, T4, T5, T6, U1 any](r Result6[T1, T2, T3, T4, T5, T6], defaultV1 U1, op func(T1, T2, T3, T4, T5, T6) Result1[U1]) Result1[U1] {
	if r.err != nil {
		return Ok1[U1](defaultV1)
	}
	return op(r.Value())
}

// Then62 ...
func Then62[T1, T2, T3, T4, T5, T6, U1, U2 any](r Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) Result2[U1, U2]) Result2[U1, U2] {
	if r.err != nil {
		return Err2[U1, U2](r.err)
	}
	return op(r.Value())
}

// Or62 ...
func Or62[T1, T2, T3, T4, T5, T6, U1, U2 any](r Result6[T1, T2, T3, T4, T5, T6], defaultV1 U1, defaultV2 U2, op func(T1, T2, T3, T4, T5, T6) Result2[U1, U2]) Result2[U1, U2] {
	if r.err != nil {
		return Ok2[U1, U2](defaultV1, defaultV2)
	}
	return op(r.Value())
}

// Then63 ...
func Then63[T1, T2, T3, T4, T5, T6, U1, U2, U3 any](r Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if r.err != nil {
		return Err3[U1, U2, U3](r.err)
	}
	return op(r.Value())
}

// Or63 ...
func Or63[T1, T2, T3, T4, T5, T6, U1, U2, U3 any](r Result6[T1, T2, T3, T4, T5, T6], defaultV1 U1, defaultV2 U2, defaultV3 U3, op func(T1, T2, T3, T4, T5, T6) Result3[U1, U2, U3]) Result3[U1, U2, U3] {
	if r.err != nil {
		return Ok3[U1, U2, U3](defaultV1, defaultV2, defaultV3)
	}
	return op(r.Value())
}

// Then64 ...
func Then64[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4 any](r Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if r.err != nil {
		return Err4[U1, U2, U3, U4](r.err)
	}
	return op(r.Value())
}

// Or64 ...
func Or64[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4 any](r Result6[T1, T2, T3, T4, T5, T6], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, op func(T1, T2, T3, T4, T5, T6) Result4[U1, U2, U3, U4]) Result4[U1, U2, U3, U4] {
	if r.err != nil {
		return Ok4[U1, U2, U3, U4](defaultV1, defaultV2, defaultV3, defaultV4)
	}
	return op(r.Value())
}

// Then65 ...
func Then65[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4, U5 any](r Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if r.err != nil {
		return Err5[U1, U2, U3, U4, U5](r.err)
	}
	return op(r.Value())
}

// Or65 ...
func Or65[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4, U5 any](r Result6[T1, T2, T3, T4, T5, T6], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, op func(T1, T2, T3, T4, T5, T6) Result5[U1, U2, U3, U4, U5]) Result5[U1, U2, U3, U4, U5] {
	if r.err != nil {
		return Ok5[U1, U2, U3, U4, U5](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5)
	}
	return op(r.Value())
}

// Then66 ...
func Then66[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4, U5, U6 any](r Result6[T1, T2, T3, T4, T5, T6], op func(T1, T2, T3, T4, T5, T6) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if r.err != nil {
		return Err6[U1, U2, U3, U4, U5, U6](r.err)
	}
	return op(r.Value())
}

// Or66 ...
func Or66[T1, T2, T3, T4, T5, T6, U1, U2, U3, U4, U5, U6 any](r Result6[T1, T2, T3, T4, T5, T6], defaultV1 U1, defaultV2 U2, defaultV3 U3, defaultV4 U4, defaultV5 U5, defaultV6 U6, op func(T1, T2, T3, T4, T5, T6) Result6[U1, U2, U3, U4, U5, U6]) Result6[U1, U2, U3, U4, U5, U6] {
	if r.err != nil {
		return Ok6[U1, U2, U3, U4, U5, U6](defaultV1, defaultV2, defaultV3, defaultV4, defaultV5, defaultV6)
	}
	return op(r.Value())
}
