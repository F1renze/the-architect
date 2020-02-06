package utils

import "testing"

func TestValidateMobile(t *testing.T) {
	tc := []struct {
		Num string
		OK  bool
	}{
		{
			"12334545761",
			false,
		},
		{
			"13926999139",
			true,
		},
	}

	for i := range tc {
		got := ValidateMobile(tc[i].Num)
		if got != tc[i].OK {
			t.Fatal("test verify mobile failed")
		}
	}
}

func TestValidateEmailFormat(t *testing.T) {
	tc := []struct {
		Addr string
		OK   bool
	}{
		{
			"ç$€§/az@gmail.com",
			false,
		},
		{
			"Eugene@gmail.com",
			true,
		},
	}
	for i := range tc {
		got := ValidateEmailFormat(tc[i].Addr)
		if got != tc[i].OK {
			t.Fatal("test verify mobile failed")
		}
	}
}
