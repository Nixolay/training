package repair_test

import (
	"testing"

	. "github.com/Nixolay/training/codewar/golang/first_non_repeating"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFirstNonRepair(t *testing.T) {
	Convey("Basic Tests", t, func() {
		Convey("should handle simple tests: a, t, e", func() {
			So(FirstNonRepeating("a"), ShouldEqual, "a")
			So(FirstNonRepeating("stress"), ShouldEqual, "t")
			So(FirstNonRepeating("moonmen"), ShouldEqual, "e")
		})
		Convey("should handle empty strings", func() {
			So(FirstNonRepeating(""), ShouldEqual, "")
		})
		Convey("should handle all repeating strings", func() {
			So(FirstNonRepeating("abba"), ShouldEqual, "")
			So(FirstNonRepeating("aa"), ShouldEqual, "")
		})
		Convey("should handle odd characters: #, w", func() {
			So(FirstNonRepeating("~><#~><"), ShouldEqual, "#")
			So(FirstNonRepeating("hello world, eh?"), ShouldEqual, "w")
		})
		Convey("should handle letter cases: T, ','", func() {
			So(FirstNonRepeating("sTreSS"), ShouldEqual, "T")
			So(FirstNonRepeating("Go hang a salami, I'm a lasagna hog!"), ShouldEqual, ",")
		})
	})

	rangeData := map[string]map[string]string{
		"should handle simple tests: a, t, e": {"a": "a", "stress": "t", "moonmen": "e", "": ""},
		"should handle all repeating strings": {"abba": "", "aa": ""},
		"should handle odd characters: #, w":  {"hello world, eh?": "w", "~><#~><": "#"},
		"should handle letter cases: T, ','":  {"sTreSS": "T", "Go hang a salami, I'm a lasagna hog!": ","},
	}

	for word, data := range rangeData {
		println(word)
		Equal(t, data)
	}
}

func Equal(t *testing.T, data map[string]string) {
	t.Helper()

	for word, expected := range map[string]string{"a": "a", "stress": "t", "moonmen": "e", "": ""} {
		if actual := FirstNonRepeating(word); actual != expected {
			t.Fatalf("Ожидалось что в слове '%s' найденный символ '%s' будет равен '%s'", word, actual, expected)
		}

		print("| ✚ ")
	}

	println()
}
