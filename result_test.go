package anyhow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Result(t *testing.T) {
	as := assert.New(t)

	t.Run("and", func(t *testing.T) {
		{
			x := Ok[int](2)
			y := Err[int](fmt.Errorf("later AError"))
			as.Equal(y, And(x, y))
		}
		{
			x := Err[string](fmt.Errorf("early AError"))
			y := Ok[string]("foo")
			as.Equal(x, And(x, y))
		}
		{
			x := Err[int](fmt.Errorf("not a 2"))
			y := Err[int](fmt.Errorf("later AError"))
			as.Equal(x, And(x, y))
		}
		{
			x := Ok(2)
			y := Ok("different result type")
			as.Equal(y, And(x, y))
		}
	})

	t.Run("and_then", func(t *testing.T) {
		sq := func(x int32) Result[int32] {
			return Ok(x * x)
		}
		fAErr := func(x int32) Result[int32] {
			return Err[int32](fmt.Errorf("%d", x))
		}
		_ = fAErr

		as.Equal(Ok[int32](16), AndThen(AndThen(Ok[int32](2), sq), sq))
		as.Equal(Err[int32](fmt.Errorf("4")), AndThen(AndThen(Ok[int32](2), sq), fAErr))
		as.Equal(Err[int32](fmt.Errorf("2")), AndThen(AndThen(Ok[int32](2), fAErr), sq))
		as.Equal(Err[int32](fmt.Errorf("3")), AndThen(AndThen(Err[int32](fmt.Errorf("3")), sq), sq))
	})

	t.Run("err", func(t *testing.T) {
		as.Equal(fmt.Errorf("err"), Err[int32](fmt.Errorf("err")).Err())
		as.Equal(nil, Ok[int32](1).Err())
	})

	t.Run("expect", func(t *testing.T) {
		as.Equal("str", Ok("str").Expect("Testing expect"))

		func() {
			defer func() {
				e := recover()
				as.Equal("Testing expect: emergency failure", fmt.Sprintf("%s", e))
			}()

			Err[string](fmt.Errorf("emergency failure")).Expect("Testing expect")
		}()
	})

	t.Run("expect_err", func(t *testing.T) {
		as.Equal(fmt.Errorf("Err"), Err[string](fmt.Errorf("Err")).ExpectErr("Testing expect"))

		func() {
			defer func() {
				e := recover()
				as.Equal("Testing expect_AErr: 10", fmt.Sprintf("%s", e))
			}()

			Ok(10).ExpectErr("Testing expect_AErr")
		}()
	})

	t.Run("inspect", func(t *testing.T) {

		{
			x := ""
			Ok[int32](4).Inspect(func(t int32) {
				x = fmt.Sprintf("original: %d", t)
			})
			as.Equal("original: 4", x)
		}
	})

	t.Run("inspect_err", func(t *testing.T) {

		{
			x := ""
			Err[int32](fmt.Errorf("Err")).InspectErr(func(t error) {
				x = fmt.Sprintf("failed: %s", t)
			})
			as.Equal("failed: Err", x)
		}
	})

	t.Run("into_err", func(t *testing.T) {
		as.Equal(fmt.Errorf("Err"), Err[int32](fmt.Errorf("Err")).IntoErr())
	})

	t.Run("into_AOk", func(t *testing.T) {
		as.Equal(int32(4), Ok[int32](4).IntoOk())
	})

	t.Run("is_AErr", func(t *testing.T) {
		as.True(Err[int32](fmt.Errorf("Err")).IsErr())
		as.False(Ok[int32](4).IsErr())
	})

	t.Run("is_AOk", func(t *testing.T) {
		as.True(Ok[int32](4).IsOk())
		as.False(Err[int32](fmt.Errorf("Err")).IsOk())
	})

	t.Run("map", func(t *testing.T) {
		as.Equal(Ok[int32](4), Map(Ok[int32](2), func(x int32) int32 {
			return x * 2
		}))
		as.Equal(Err[int32](fmt.Errorf("Err")), Map(Err[int32](fmt.Errorf("Err")), func(x int32) int32 {
			return x * 2
		}))
	})

	t.Run("map_or", func(t *testing.T) {
		as.Equal(int32(4), MapOr(Ok[int32](2), 4, func(x int32) int32 {
			return x * 2
		}))
		as.Equal(int32(3), MapOr(Err[int32](fmt.Errorf("Err")), 3, func(x int32) int32 {
			return x * 2
		}))
	})

	t.Run("map_or_else", func(t *testing.T) {
		as.Equal(int32(4), MapOrElse(Ok[int32](2), func(AErr error) int32 {
			return 3
		}, func(x int32) int32 {
			return x * x
		}))
		as.Equal(int32(3), MapOrElse(Err[int32](fmt.Errorf("err")), func(AErr error) int32 {
			return 3
		}, func(x int32) int32 {
			return x * x
		}))
	})

	t.Run("ok", func(t *testing.T) {
		var s1 = "1"
		as.Equal(&s1, Ok("1").Ok())

		var s2 *string = nil
		as.Equal(s2, Err[string](fmt.Errorf("err")).Ok())
	})

	t.Run("or", func(t *testing.T) {
		{
			x := Ok[int](2)
			y := Err[int](fmt.Errorf("later AError"))
			as.Equal(x, x.Or(y))
		}
		{
			x := Err[string](fmt.Errorf("early AError"))
			y := Ok[string]("foo")
			as.Equal(y, x.Or(y))
		}
		{
			x := Err[int](fmt.Errorf("not a 2"))
			y := Err[int](fmt.Errorf("later AError"))
			as.Equal(y, x.Or(y))
		}
		{
			x := Ok[int](2)
			y := Ok[int](100)
			as.Equal(x, x.Or(y))
		}
	})

	t.Run("unwrap", func(t *testing.T) {
		as.Equal("str", Ok("str").Unwrap())

		func() {
			defer func() {
				e := recover()
				as.Equal("emergency failure", fmt.Sprintf("%s", e))
			}()

			Err[string](fmt.Errorf("emergency failure")).Unwrap()
		}()
	})

	t.Run("unwrap_err", func(t *testing.T) {
		as.Equal(fmt.Errorf("emergency failure"), Err[string](fmt.Errorf("emergency failure")).UnwrapErr())

		func() {
			defer func() {
				e := recover()
				as.Equal("string", fmt.Sprintf("%s", e))
			}()

			Ok[string]("string").UnwrapErr()
		}()
	})

	t.Run("unwrap_err_unchecked", func(t *testing.T) {
		as.Equal(fmt.Errorf("emergency failure"), Err[string](fmt.Errorf("emergency failure")).UnwrapErrUnchecked())
		as.Nil(Ok[string]("string").UnwrapErrUnchecked())
	})

	t.Run("unwrap_or", func(t *testing.T) {
		as.Equal(int32(9), Ok[int32](9).UnwrapOr(1))
		as.Equal(int32(1), Err[int32](fmt.Errorf("str")).UnwrapOr(1))
	})

	t.Run("unwrap_or_else", func(t *testing.T) {
		as.Equal(int32(9), Ok[int32](9).UnwrapOrElse(func(err error) int32 {
			return 1
		}))
		as.Equal(int32(1), Err[int32](fmt.Errorf("str")).UnwrapOrElse(func(err error) int32 {
			return 1
		}))
	})

	t.Run("unwrap_unchecked", func(t *testing.T) {
		as.Equal(int32(9), Ok[int32](9).UnwrapUnchecked())
		as.Equal(int32(0), Err[int32](fmt.Errorf("str")).UnwrapUnchecked())
	})
}
