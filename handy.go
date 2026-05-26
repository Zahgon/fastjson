package fastjson

var handyPool ParserPool

// GetString returns string value for the field identified by keys path
// in JSON data.
//
// Array indexes may be represented as decimal numbers in keys.
//
// An empty string is returned on error. Use Parser for proper error handling.
//
// Parser is faster for obtaining multiple fields from JSON.
func GetString(data []byte, keys ...string) string { _ = "STUB: not implemented"; return "" }

// GetBytes returns string value for the field identified by keys path
// in JSON data.
//
// Array indexes may be represented as decimal numbers in keys.
//
// nil is returned on error. Use Parser for proper error handling.
//
// Parser is faster for obtaining multiple fields from JSON.
func GetBytes(data []byte, keys ...string) []byte { _ = "STUB: not implemented"; return nil }

// Make a copy of sb, since sb belongs to p.

// GetInt returns int value for the field identified by keys path
// in JSON data.
//
// Array indexes may be represented as decimal numbers in keys.
//
// 0 is returned on error. Use Parser for proper error handling.
//
// Parser is faster for obtaining multiple fields from JSON.
func GetInt(data []byte, keys ...string) int { _ = "STUB: not implemented"; return 0 }

// GetFloat64 returns float64 value for the field identified by keys path
// in JSON data.
//
// Array indexes may be represented as decimal numbers in keys.
//
// 0 is returned on error. Use Parser for proper error handling.
//
// Parser is faster for obtaining multiple fields from JSON.
func GetFloat64(data []byte, keys ...string) float64 { _ = "STUB: not implemented"; return 0 }

// GetBool returns boolean value for the field identified by keys path
// in JSON data.
//
// Array indexes may be represented as decimal numbers in keys.
//
// False is returned on error. Use Parser for proper error handling.
//
// Parser is faster for obtaining multiple fields from JSON.
func GetBool(data []byte, keys ...string) bool { _ = "STUB: not implemented"; return false }

// Exists returns true if the field identified by keys path exists in JSON data.
//
// Array indexes may be represented as decimal numbers in keys.
//
// False is returned on error. Use Parser for proper error handling.
//
// Parser is faster when multiple fields must be checked in the JSON.
func Exists(data []byte, keys ...string) bool { _ = "STUB: not implemented"; return false }

// Parse parses json string s.
//
// The function is slower than the Parser.Parse for re-used Parser.
func Parse(s string) (*Value, error) { _ = "STUB: not implemented"; return nil, nil }

// MustParse parses json string s.
//
// The function panics if s cannot be parsed.
// The function is slower than the Parser.Parse for re-used Parser.
func MustParse(s string) *Value { _ = "STUB: not implemented"; return nil }

// ParseBytes parses b containing json.
//
// The function is slower than the Parser.ParseBytes for re-used Parser.
func ParseBytes(b []byte) (*Value, error) { _ = "STUB: not implemented"; return nil, nil }

// MustParseBytes parses b containing json.
//
// The function panics if b cannot be parsed.
// The function is slower than the Parser.ParseBytes for re-used Parser.
func MustParseBytes(b []byte) *Value { _ = "STUB: not implemented"; return nil }
