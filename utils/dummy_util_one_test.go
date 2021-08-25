package utils

import (
	"testing"
)

func TestRunOneUtil(t *testing.T) {
	msg := "hello"
	res_msg := RunUtilOne(msg)
	if res_msg != msg+"!" {
		t.Error("Mismatched Message")
	}

}
