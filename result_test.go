package anyhow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnd(t *testing.T) {
	as := assert.New(t)

	t.Run("and", func(t *testing.T) {
		{
			x := AOk[int](2)
			y := AErr[int](fmt.Errorf("later AError"))
			as.Equal(y, And(x, y))
		}
		{
			x := AErr[string](fmt.Errorf("early AError"))
			y := AOk[string]("foo")
			as.Equal(x, And(x, y))
		}
		{
			x := AErr[int](fmt.Errorf("not a 2"))
			y := AErr[int](fmt.Errorf("later AError"))
			as.Equal(x, And(x, y))
		}
		{
			x := AOk(2)
			y := AOk("different result type")
			as.Equal(y, And(x, y))
		}
	})

	t.Run("and_then", func(t *testing.T) {
		sq := func(x int32) Result[int32] {
			return AOk(x * x)
		}
		fAErr := func(x int32) Result[int32] {
			return AErr[int32](fmt.Errorf("%d", x))
		}
		_ = fAErr

		as.Equal(AOk[int32](16), AndThen(AndThen(AOk[int32](2), sq), sq))
		as.Equal(AErr[int32](fmt.Errorf("4")), AndThen(AndThen(AOk[int32](2), sq), fAErr))
		as.Equal(AErr[int32](fmt.Errorf("2")), AndThen(AndThen(AOk[int32](2), fAErr), sq))
		as.Equal(AErr[int32](fmt.Errorf("3")), AndThen(AndThen(AErr[int32](fmt.Errorf("3")), sq), sq))
	})

	t.Run("expect", func(t *testing.T) {
		as.Equal("str", Expect(AOk("str"), "Testing expect"))

		func() {
			defer func() {
				e := recover()
				as.Equal("Testing expect: emergency failure", fmt.Sprintf("%s", e))
			}()

			Expect(AErr[string](fmt.Errorf("emergency failure")), "Testing expect")
		}()
	})

	t.Run("expect_AErr", func(t *testing.T) {
		as.Equal(fmt.Errorf("AErr"), ExpectErr(AErr[string](fmt.Errorf("AErr")), "Testing expect"))

		func() {
			defer func() {
				e := recover()
				as.Equal("Testing expect_AErr: 10", fmt.Sprintf("%s", e))
			}()

			ExpectErr(AOk(10), "Testing expect_AErr")
		}()
	})

	t.Run("inspect", func(t *testing.T) {
		x := ""
		Inspect(AOk[int32](4), func(t int32) {
			x = fmt.Sprintf("original: %d", t)
		})
		as.Equal("original: 4", x)
	})

	t.Run("inspect_AErr", func(t *testing.T) {
		x := ""
		InspectErr(AErr[int32](fmt.Errorf("AErr")), func(t error) {
			x = fmt.Sprintf("failed: %s", t)
		})
		as.Equal("failed: AErr", x)
	})

	t.Run("into_AErr", func(t *testing.T) {
		as.Equal(fmt.Errorf("AErr"), IntoErr(AErr[int32](fmt.Errorf("AErr"))))
	})

	t.Run("into_AOk", func(t *testing.T) {
		as.Equal(int32(4), IntoOk(AOk[int32](4)))
	})

	t.Run("is_AErr", func(t *testing.T) {
		as.True(IsErr(AErr[int32](fmt.Errorf("AErr"))))
		as.False(IsErr(AOk[int32](4)))
	})

	t.Run("is_AOk", func(t *testing.T) {
		as.True(IsOk(AOk[int32](4)))
		as.False(IsOk(AErr[int32](fmt.Errorf("AErr"))))
	})

	t.Run("map", func(t *testing.T) {
		as.Equal(AOk[int32](4), Map(AOk[int32](2), func(x int32) int32 {
			return x * 2
		}))
		as.Equal(AErr[int32](fmt.Errorf("AErr")), Map(AErr[int32](fmt.Errorf("AErr")), func(x int32) int32 {
			return x * 2
		}))
	})

	t.Run("map_or", func(t *testing.T) {
		as.Equal(int32(4), MapOr(AOk[int32](2), 4, func(x int32) int32 {
			return x * 2
		}))
		as.Equal(int32(3), MapOr(AErr[int32](fmt.Errorf("AErr")), 3, func(x int32) int32 {
			return x * 2
		}))
	})

	t.Run("map_or_else", func(t *testing.T) {
		as.Equal(int32(4), MapOrElse(AOk[int32](2), func(AErr error) int32 {
			return 3
		}, func(x int32) int32 {
			return x * x
		}))
		as.Equal(int32(3), MapOrElse(AErr[int32](fmt.Errorf("err")), func(AErr error) int32 {
			return 3
		}, func(x int32) int32 {
			return x * x
		}))
	})

	t.Run("ok", func(t *testing.T) {
		var s1 = "1"
		as.Equal(&s1, Ok(AOk("1")))

		var s2 *string = nil
		as.Equal(s2, Ok(AErr[string](fmt.Errorf("err"))))
	})

	t.Run("or", func(t *testing.T) {
		{
			x := AOk[int](2)
			y := AErr[int](fmt.Errorf("later AError"))
			as.Equal(x, Or(x, y))
		}
		{
			x := AErr[string](fmt.Errorf("early AError"))
			y := AOk[string]("foo")
			as.Equal(y, Or(x, y))
		}
		{
			x := AErr[int](fmt.Errorf("not a 2"))
			y := AErr[int](fmt.Errorf("later AError"))
			as.Equal(y, Or(x, y))
		}
		{
			x := AOk[int](2)
			y := AOk[int](100)
			as.Equal(x, Or(x, y))
		}
	})

	t.Run("unwrap", func(t *testing.T) {
		as.Equal("str", Unwrap(AOk("str")))

		func() {
			defer func() {
				e := recover()
				as.Equal("emergency failure", fmt.Sprintf("%s", e))
			}()

			Unwrap(AErr[string](fmt.Errorf("emergency failure")))
		}()
	})

	t.Run("unwrap_err", func(t *testing.T) {
		as.Equal(fmt.Errorf("emergency failure"), UnwrapErr(AErr[string](fmt.Errorf("emergency failure"))))

		func() {
			defer func() {
				e := recover()
				as.Equal("string", fmt.Sprintf("%s", e))
			}()

			UnwrapErr(AOk[string]("string"))
		}()
	})

	t.Run("unwrap_err_unchecked", func(t *testing.T) {
		as.Equal(fmt.Errorf("emergency failure"), UnwrapErrUnchecked(AErr[string](fmt.Errorf("emergency failure"))))
		as.Nil(UnwrapErrUnchecked(AOk[string]("string")))
	})

	t.Run("unwrap_or", func(t *testing.T) {
		as.Equal(int32(9), UnwrapOr(AOk[int32](9), 1))
		as.Equal(int32(1), UnwrapOr(AErr[int32](fmt.Errorf("str")), 1))
	})

	t.Run("unwrap_or_else", func(t *testing.T) {
		as.Equal(int32(9), UnwrapOrElse(AOk[int32](9), func(err error) int32 {
			return 1
		}))
		as.Equal(int32(1), UnwrapOrElse(AErr[int32](fmt.Errorf("str")), func(err error) int32 {
			return 1
		}))
	})

	t.Run("unwrap_unchecked", func(t *testing.T) {
		as.Equal(int32(9), UnwrapUnchecked(AOk[int32](9)))
		as.Equal(int32(0), UnwrapUnchecked(AErr[int32](fmt.Errorf("str"))))
	})
}
