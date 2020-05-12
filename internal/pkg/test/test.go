package test

import (
	"os"
	"testing"
)

func Assert(t *testing.T, title string, actual string, expect string) {
	if actual != expect {
		t.Errorf("\n%v \n-----------------------\nexpects\n%v \n\nreceives\n%v \n\n", title, expect, actual)
		os.Exit(-1)
	}
}
