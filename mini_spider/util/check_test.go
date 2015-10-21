package util

import (
	"testing"
)

func Test_CheckBaseurl(t *testing.T) {
	if u, err := CheckBaseurl("www.example.com"); u != "http://www.example.com/" || err != nil {
		t.Error("Failed")
	} else {
		t.Log("Pass")
	}
}
