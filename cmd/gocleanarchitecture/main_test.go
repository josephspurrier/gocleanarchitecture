package main

import "testing"

// TestMain ensures Cool returns the correct string.
func TestMain(t *testing.T) {
	if Cool() != "cool" {
		t.Fatal("Cool failed.")
	}
}
