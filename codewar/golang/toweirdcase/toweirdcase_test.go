package toweirdcase

import (
	"testing"
)

func TestToWeirdCase(t *testing.T) {
	if toWeirdCase("abc def") != "AbC DeF" {
		t.Fatal("is not AbC DeF")
	}

	if toWeirdCase("This is a test Looks like you passed") != "ThIs Is A TeSt LoOkS     LiKe YoU PaSsEd" {
		t.Fatal("is not equil: ThIs Is A TeSt LoOkS     LiKe YoU PaSsEd")
	}
}
