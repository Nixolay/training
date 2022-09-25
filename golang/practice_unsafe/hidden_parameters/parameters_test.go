package parameters_test

import (
	"bytes"
	"testing"

	. "github.com/Nixolay/training/golang/practice_unsafe/hidden_parameters"
	"github.com/Nixolay/training/golang/practice_unsafe/hidden_parameters/hidden"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/require"
)

func TestGetHidden(t *testing.T) {
	Convey("Test get hidden parameters", t, func() {
		Convey("Func GetHidden", func() {
			data := hidden.CreateHidden()

			expected := GetHidden(data)
			So(expected, ShouldEqual, hidden.IsHidden)
		})

		Convey("Func UnsafeGetBuf", func() {
			data := "test data"
			bb := bytes.NewBufferString(data)

			buf := UnsafeGetBuf(bb)
			require.Equal(t, string(data), data)
			So(string(buf), ShouldEqual, data)
		})
	})
}

func TestUnsafeGetBuf(t *testing.T) {
	data := "test data"
	bb := bytes.NewBufferString(data)
	require.Equal(t, string(UnsafeGetBuf(bb)), data)
}
