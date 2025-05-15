package profile_test

import (
	"github.com/starbase-343/globalinv/internal/config/profile"
	"testing"
)

func TestParse_Dev(t *testing.T) {
	p, err := profile.Parse("dev")
	if err != nil {
		t.Fatalf("Parse(\"dev\") returned error: %v", err)
	}
	if p != profile.Dev {
		t.Errorf("Parse(\"dev\") = %v; want %v", p, profile.Dev)
	}
}

func TestParse_Prod(t *testing.T) {
	p, err := profile.Parse("prod")
	if err != nil {
		t.Fatalf("Parse(\"prod\") returned error: %v", err)
	}
	if p != profile.Prod {
		t.Errorf("Parse(\"prod\") = %v; want %v", p, profile.Prod)
	}
}

func TestParse_Invalid(t *testing.T) {
	invalidInputs := []string{"", "DEV", "production", "staging", "development"}
	for _, input := range invalidInputs {
		p, err := profile.Parse(input)
		if err == nil {
			t.Errorf("Parse(%q) expected error, got nil, profile=%v", input, p)
			continue
		}
		errMsg := err.Error()
		if !contains(errMsg, input) {
			t.Errorf("Parse(%q) error = %q; want message to contain %q", input, errMsg, input)
		}
		if p != "unknown" {
			t.Errorf("Parse(%q) = %v; want %q on error", input, p, "unknown")
		}
	}
}

func contains(str, substr string) bool {
	return len(substr) == 0 || (len(str) >= len(substr) && (stringIndex(str, substr) >= 0))
}

func stringIndex(s, sep string) int {
	for i := 0; i+len(sep) <= len(s); i++ {
		if s[i:i+len(sep)] == sep {
			return i
		}
	}
	return -1
}
