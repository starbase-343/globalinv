package env_test

import (
	"github.com/starbase-343/globalinv/internal/config/env"
	"os"
	"testing"
)

func TestOrDefault_VarSet(t *testing.T) {
	key := env.Key("TEST_ENV_VAR")
	expected := "value123"
	err := os.Setenv(string(key), expected)
	if err != nil {
		t.Fatalf("Failed to set env: %v", err)
	}
	defer os.Unsetenv(string(key))

	result := env.OrDefault(key, "defaultVal")
	if result != expected {
		t.Errorf("OrDefault returned %q; want %q when env is set", result, expected)
	}
}

func TestOrDefault_VarNotSet(t *testing.T) {
	key := env.Key("TEST_ENV_VAR_NOT_SET")
	os.Unsetenv(string(key))

	defaultVal := "myDefault"
	result := env.OrDefault(key, defaultVal)
	if result != defaultVal {
		t.Errorf("OrDefault returned %q; want default %q when env is not set", result, defaultVal)
	}
}

func TestMust_VarSet(t *testing.T) {
	key := env.Key("TEST_MUST_VAR")
	expected := "mustValue"
	os.Setenv(string(key), expected)
	defer os.Unsetenv(string(key))

	result := env.Must(key)
	if result != expected {
		t.Errorf("Must returned %q; want %q when env is set", result, expected)
	}
}

func TestMust_VarNotSet_Panics(t *testing.T) {
	key := env.Key("TEST_MUST_VAR_NOT_SET")
	os.Unsetenv(string(key))

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Must did not panic when env var %s is not set", key)
		} else {
			errMsg, ok := r.(error)
			if !ok {
				t.Fatalf("Recovered value is not error: %v", r)
			}
			if !contains(errMsg.Error(), string(key)) {
				t.Errorf("Panic error message %q does not contain key %q", errMsg.Error(), key)
			}
		}
	}()
	env.Must(key)
}

func contains(s, sub string) bool {
	return len(sub) == 0 || (len(s) >= len(sub) && (index(s, sub) >= 0))
}

func index(str, substr string) int {
	for i := 0; i+len(substr) <= len(str); i++ {
		if str[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
