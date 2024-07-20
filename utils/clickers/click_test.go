package main

import (
	"fmt"
	"testing"
)

func TestClick(t *testing.T) {
	adb := new(ADB)
	adb.InitDevices()

	fmt.Println(adb.data)
}
