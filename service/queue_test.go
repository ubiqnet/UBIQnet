package service

import (
	"testing"
	"time"
)

func TestName(t *testing.T) {
	ticker := time.NewTicker(2 * time.Second)
	a := 1
	for {
		<-ticker.C
		a++
		println(a)
		if a > 5 {
			ticker.Stop()
		}
	}
}
