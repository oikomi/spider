package util

import (
    "testing"
)

func Test_ParseSchemeHost(t *testing.T) {
    if u, err := ParseSchemeHost("https://www.example.com/test.html"); u != "https://www.example.com" || err != nil { 
        t.Error("Failed") 
    } else {
    	t.Log("Pass")
    }
}
