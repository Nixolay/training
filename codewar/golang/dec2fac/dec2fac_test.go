package dec2fac

import "testing"

func TestFac2string(t *testing.T) {
	println(Dec2FactString(463))
	println(FactString2Dec(Dec2FactString(463)))
}
