package common

import (
	"testing"
)

var (
	input1          int64  = 1
	expectedResult1 string = "rghtzlbmgjroc2fq"
	input2          int64  = 2
	expectedResult2 string = "rghtzlbmgjroc29q"
)

func TestMixerEncode(t *testing.T) {
	result := MixerEncode(input1)

	if result != expectedResult1 {
		t.Errorf("SALAH !!! harusnya %s", result)
	} else {
		t.Log("BENAR !!!")
	}
}

func TestMixerEncode2(t *testing.T) {
	result := MixerEncode(input2)

	if result != expectedResult2 {
		t.Errorf("SALAH !!! harusnya %s", result)
	} else {
		t.Log("BENAR !!!")
	}
}
