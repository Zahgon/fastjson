package fastjson

// Validate validates JSON s.
func Validate(s string) error { _ = "STUB: not implemented"; return nil }

// ValidateBytes validates JSON b.
func ValidateBytes(b []byte) error { _ = "STUB: not implemented"; return nil }

func validateValue(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Scan the string for control chars.

func validateArray(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func validateObject(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Parse key.

// Scan the key for control chars.

// Parse value

// validateKey is similar to validateString, but is optimized
// for typical object keys, which are quite small and have no escape sequences.
func validateKey(s string) (string, string, error) { _ = "STUB: not implemented"; return "", "", nil }

// Fast path - the key doesn't contain escape sequences.

// Slow path - the key contains escape sequences.

func validateString(s string) (string, string, error) {
	_ = "STUB: not implemented"
	// Try fast path - a string without escape sequences.
	return "", "", nil
}

// Slow path - escape sequences are present.

// Valid escape sequences - see http://json.org/

func validateNumber(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Validate fractional part

// Validate exponent part
