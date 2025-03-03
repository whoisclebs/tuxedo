package tuxedo

import (
	"reflect"
	"testing"
)

func assetEqual(t *testing.T, expected, actual interface{}, msg string) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Assert Equal Failed: %s. Expected: %#v, Got: %#v", msg, expected, actual)
	}
}

func assetNotEqual(t *testing.T, expected, actual interface{}, msg string) {
	if reflect.DeepEqual(expected, actual) {
		t.Errorf("Assert Not Equal Failed: %s. Expected: %#v, Got: %#v", msg, expected, actual)
	}
}

func assertNoError(t *testing.T, err error, msg string) {
	if err != nil {
		t.Errorf("Assert NoError Failed: %s. Error: %v", msg, err)
	}
}

func assertTrue(t *testing.T, condition bool, msg string) {
	if !condition {
		t.Errorf("Assert True Failed: %s", msg)
	}
}
