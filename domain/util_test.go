package domain_test

import "testing"

// AssertEqual throws an error if the two values are not equal.
func AssertEqual(t *testing.T, actualValue interface{}, expectedValue interface{}) {
	if actualValue != expectedValue {
		t.Errorf("\n got: %v\nwant: %v", actualValue, expectedValue)
	}
}

// AssertNotNil throws an error if the value is nil.
func AssertNotNil(t *testing.T, actualValue interface{}) {
	if actualValue == nil {
		t.Errorf("\n got: %v\ndidn't want: %v", actualValue, nil)
	}
}
