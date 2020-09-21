package test_support

import (
	"encoding/base64"
	"strings"
	"testing"
)

func ExpectEqual(t *testing.T, expected string, actual string) {
	trimExpected := strings.TrimSpace(expected)
	trimActual := strings.TrimSpace(actual)

	if strings.Compare(trimExpected, trimActual) != 0 {
		t.Errorf("Expected does not match actual, '%s' != '%s'", trimExpected, trimActual)
	}
}

func ExpectEqualInt(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Expected does not match actual, %d != %d", expected, actual)
	}
}

func ExpectNotEmpty(t *testing.T, value *map[string]string, valueName string) {
	if len(*value) == 0 {
		t.Errorf("%s should not be empty", valueName)
	}
}

func ExpectEqualBase64(t *testing.T, expectedString string, actualBytes string) {

	expected, err := base64.StdEncoding.DecodeString(expectedString)
	if err != nil {
		t.Error(err)
	}

	actual, err := base64.StdEncoding.DecodeString(actualBytes)
	if err != nil {
		t.Error(err)
	}

	trimExpected := strings.TrimSpace(string(expected))
	trimActual := strings.TrimSpace(string(actual))

	if strings.Compare(trimExpected, trimActual) != 0 {
		t.Errorf("Expected does not match actual, '%s' != '%s'", trimExpected, trimActual)
	}
}
