package validator

import "testing"

func ValidatiorTest(t *testing.T) {
	// Test Required
	if Required("") != true {
		t.Error("Required function failed")
	}
	if Required("test") != false {
		t.Error("Required function failed")
	}

	// Test MinLength
	if MinLength("test", 5) != true {
		t.Error("MinLength function failed")
	}
	if MinLength("test", 2) != false {
		t.Error("MinLength function failed")
	}

	// Test MaxLength
	if MaxLength("test", 2) != true {
		t.Error("MaxLength function failed")
	}
	if MaxLength("test", 5) != false {
		t.Error("MaxLength function failed")
	}

	// Test Email
	if Email("test") != true {
		t.Error("Email function failed")
	}
	if Email("jainil@gmail.com") != false {
		t.Error("Email function failed")
	}
}
