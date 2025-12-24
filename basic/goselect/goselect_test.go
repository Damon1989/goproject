package goselect

import "testing"

func TestSelectSmt(t *testing.T) {
	for i := 0; i < 10; i++ {
		SelectSmt()
	}

}

func TestSelectMain(t *testing.T) {
	SelectMain()
}
