package cuttingrope

import (
	"testing"
)

func Test_case1(t *testing.T) {
	if maxProductAfterCutting(1) == 0 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}
}

func Test_case2(t *testing.T) {
	if maxProductAfterCutting(2) == 1 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}
}

func Test_case3(t *testing.T) {
	if maxProductAfterCutting(3) == 2 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}
}

func Test_case4(t *testing.T) {
	if maxProductAfterCutting(4) == 4 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}
}

func Test_case5(t *testing.T) {
	if maxProductAfterCutting(5) == 6 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}
}

func Test_case6(t *testing.T) {
	if maxProductAfterCutting(6) == 9 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}
}

func Test_case7(t *testing.T) {
	if maxProductAfterCutting(7) == 12 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}
}

func Test_case8(t *testing.T) {
	if maxProductAfterCutting(8) == 18 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}
}

func Test_case9(t *testing.T) {
	if maxProductAfterCutting(9) == 27 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}
}

func Test_case10(t *testing.T) {
	if maxProductAfterCutting(10) == 36 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}
}

func Test_case11(t *testing.T) {
	if maxProductAfterCutting(50) == 86093442 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}
}
