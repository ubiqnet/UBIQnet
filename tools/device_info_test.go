package tools

import (
	"testing"
)

func TestSpeedFun(t *testing.T) {
	down, up := STestSpeed(3)
	println(down)
	println(up)

}
