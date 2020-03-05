package main

import (
	"testing"
)

func TestGetServicePort(t *testing.T) {
	expects := GetServicePort("Test-GO-PORT", ":12345")
	if expects != ":12345" {
		t.Errorf("FAIL")
	}

}
