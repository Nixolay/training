/*
Современные римские цифры пишутся путем выражения
каждой десятичной цифры числа, которое должно быть
закодировано отдельно, начиная с крайней левой цифры
и пропуская любые 0. Таким образом, 1990 отображается
как:
«MCMXC» (1000 = M, 900 = CM, 90 = XC), а
2008 - «MMVIII» (2000 = MM, 8 = VIII).
Римская цифра 1666 года «MDCLXVI»
использует каждую букву в порядке убывания.

Пример:
*/
package romandecoder //nolint:testpackage

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDecode(t *testing.T) {
	Convey("test roman to decimal converter", t, func() {
		So(Decode("XXI"), ShouldEqual, 21)
		So(Decode("I"), ShouldEqual, 1)
		So(Decode("IV"), ShouldEqual, 4)
		So(Decode("MMVIII"), ShouldEqual, 2008)
		So(Decode("MDCLXVI"), ShouldEqual, 1666)
		So(Decode("CDXX"), ShouldEqual, 420)
		So(Decode("MMMIX"), ShouldEqual, 3009)
		So(Decode("CDLXXX"), ShouldEqual, 480)
		So(Decode("CXLIII"), ShouldEqual, 143)
		So(Decode("CDLXXVII"), ShouldEqual, 477)
		So(Decode("MMXC"), ShouldEqual, 2090)
		So(Decode("MMCDXX"), ShouldEqual, 2420)
	})
}
