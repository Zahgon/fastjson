package fastfloat

import (
	"math"
)

// ParseUint64BestEffort parses uint64 number s.
//
// It is equivalent to strconv.ParseUint(s, 10, 64), but is faster.
//
// 0 is returned if the number cannot be parsed.
// See also ParseUint64, which returns parse error if the number cannot be parsed.
func ParseUint64BestEffort(s string) uint64 { _ = "STUB: not implemented"; return 0 }

// The integer part may be out of range for uint64.
// Fall back to slow parsing.

// Unparsed tail left.

// ParseUint64 parses uint64 from s.
//
// It is equivalent to strconv.ParseUint(s, 10, 64), but is faster.
//
// See also ParseUint64BestEffort.
func ParseUint64(s string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// The integer part may be out of range for uint64.
// Fall back to slow parsing.

// Unparsed tail left.

// ParseInt64BestEffort parses int64 number s.
//
// It is equivalent to strconv.ParseInt(s, 10, 64), but is faster.
//
// 0 is returned if the number cannot be parsed.
// See also ParseInt64, which returns parse error if the number cannot be parsed.
func ParseInt64BestEffort(s string) int64 { _ = "STUB: not implemented"; return 0 }

// The integer part may be out of range for int64.
// Fall back to slow parsing.

// Unparsed tail left.

// ParseInt64 parses int64 number s.
//
// It is equivalent to strconv.ParseInt(s, 10, 64), but is faster.
//
// See also ParseInt64BestEffort.
func ParseInt64(s string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// The integer part may be out of range for int64.
// Fall back to slow parsing.

// Unparsed tail left.

// Exact powers of 10.
//
// This works faster than math.Pow10, since it avoids additional multiplication.
var float64pow10 = [...]float64{
	1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9, 1e10, 1e11, 1e12, 1e13, 1e14, 1e15, 1e16,
}

// ParseBestEffort parses floating-point number s.
//
// It is equivalent to strconv.ParseFloat(s, 64), but is faster.
//
// 0 is returned if the number cannot be parsed.
// See also Parse, which returns parse error if the number cannot be parsed.
func ParseBestEffort(s string) float64 { _ = "STUB: not implemented"; return 0 }

// the integer part might be elided to remain compliant
// with https://go.dev/ref/spec#Floating-point_literals

// The integer part may be out of range for uint64.
// Fall back to slow parsing.

// "infinity" is needed for OpenMetrics support.
// See https://github.com/OpenObservability/OpenMetrics/blob/master/OpenMetrics.md

// Fast path - just integer.

// Parse fractional part.

// the fractional part may be elided to remain compliant
// with https://go.dev/ref/spec#Floating-point_literals

// The mantissa is out of range. Fall back to standard parsing.

// Convert the entire mantissa to a float at once to avoid rounding errors.

// Fast path - parsed fractional number.

// Parse exponent part.

// The exponent may be too big for float64.
// Fall back to standard parsing.

// Parse parses floating-point number s.
//
// It is equivalent to strconv.ParseFloat(s, 64), but is faster.
//
// See also ParseBestEffort.
func Parse(s string) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

// the integer part might be elided to remain compliant
// with https://go.dev/ref/spec#Floating-point_literals

// The integer part may be out of range for uint64.
// Fall back to slow parsing.

// "infinity" is needed for OpenMetrics support.
// See https://github.com/OpenObservability/OpenMetrics/blob/master/OpenMetrics.md

// Fast path - just integer.

// Parse fractional part.

// the fractional part might be elided to remain compliant
// with https://go.dev/ref/spec#Floating-point_literals

// The mantissa is out of range. Fall back to standard parsing.

// Convert the entire mantissa to a float at once to avoid rounding errors.

// Fast path - parsed fractional number.

// Parse exponent part.

// The exponent may be too big for float64.
// Fall back to standard parsing.

var inf = math.Inf(1)
var nan = math.NaN()
