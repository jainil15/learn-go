package validator

import "testing"

func ValidatiorTest(t *testing.T) {
	// Test Required
	if IsRequired("") != true {
		t.Error("Required function failed")
	}
	if IsRequired("test") != false {
		t.Error("Required function failed")
	}

	// Test MinLength
	if IsMinLength("test", 5) != true {
		t.Error("MinLength function failed")
	}
	if IsMinLength("test", 2) != false {
		t.Error("MinLength function failed")
	}

	// Test MaxLength
	if IsMaxLength("test", 2) != true {
		t.Error("MaxLength function failed")
	}
	if IsMaxLength("test", 5) != false {
		t.Error("MaxLength function failed")
	}

	// Test Email
	if IsEmail("test") != true {
		t.Error("Email function failed")
	}
	if IsEmail("jainil@gmail.com") != false {
		t.Error("Email function failed")
	}
}
