package validator

import "testing"

func TestValidator(t *testing.T) {
	t.Run("Testing Required", func(t *testing.T) {
		if IsRequired("") != false {
			t.Error("Required function failed")
		}
		if IsRequired("test") != true {
			t.Error("Required function failed")
		}
	})

	t.Run("Testing Min Length", func(t *testing.T) {
		// Test MinLength
		if IsMinLength("test", 5) != false {
			t.Error("MinLength function failed")
		}
		if IsMinLength("test", 2) != true {
			t.Error("MinLength function failed")
		}
	})

	t.Run("Testing Max Length", func(t *testing.T) {
		// Test MaxLength
		if IsMaxLength("test", 2) != false {
			t.Error("MaxLength function failed")
		}
		if IsMaxLength("test", 5) != true {
			t.Error("MaxLength function failed")
		}
	})

	t.Run("Testing Max Length", func(t *testing.T) {
		if IsEmail("test") != false {
			t.Error("Email function failed")
		}
		if IsEmail("jainil@gmail.com") != true {
			t.Error("Email function failed")
		}
	})
}