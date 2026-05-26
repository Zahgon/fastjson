package fastjson

// Arena may be used for fast creation and re-use of Values.
//
// Typical Arena lifecycle:
//
//  1. Construct Values via the Arena and Value.Set* calls.
//  2. Marshal the constructed Values with Value.MarshalTo call.
//  3. Reset all the constructed Values at once by Arena.Reset call.
//  4. Go to 1 and re-use the Arena.
//
// It is unsafe calling Arena methods from concurrent goroutines.
// Use per-goroutine Arenas or ArenaPool instead.
type Arena struct {
	b []byte
	c cache
}

// Reset resets all the Values allocated by a.
//
// Values previously allocated by a cannot be used after the Reset call.
func (a *Arena) Reset() { _ = "STUB: not implemented"; return }

// NewObject returns new empty object value.
//
// New entries may be added to the returned object via Set call.
//
// The returned object is valid until Reset is called on a.
func (a *Arena) NewObject() *Value { _ = "STUB: not implemented"; return nil }

// NewArray returns new empty array value.
//
// New entries may be added to the returned array via Set* calls.
//
// The returned array is valid until Reset is called on a.
func (a *Arena) NewArray() *Value { _ = "STUB: not implemented"; return nil }

// NewString returns new string value containing s.
//
// The returned string is valid until Reset is called on a.
func (a *Arena) NewString(s string) *Value { _ = "STUB: not implemented"; return nil }

// NewStringBytes returns new string value containing b.
//
// The returned string is valid until Reset is called on a.
func (a *Arena) NewStringBytes(b []byte) *Value { _ = "STUB: not implemented"; return nil }

// NewNumberFloat64 returns new number value containing f.
//
// The returned number is valid until Reset is called on a.
func (a *Arena) NewNumberFloat64(f float64) *Value { _ = "STUB: not implemented"; return nil }

// NewNumberInt returns new number value containing n.
//
// The returned number is valid until Reset is called on a.
func (a *Arena) NewNumberInt(n int) *Value { _ = "STUB: not implemented"; return nil }

// NewNumberString returns new number value containing s.
//
// The returned number is valid until Reset is called on a.
func (a *Arena) NewNumberString(s string) *Value { _ = "STUB: not implemented"; return nil }

// NewNull returns null value.
func (a *Arena) NewNull() *Value {
	_ = "STUB: not implemented"

	// NewTrue returns true value.
	return nil
}

func (a *Arena) NewTrue() *Value {
	_ = "STUB: not implemented"

	// NewFalse return false value.
	return nil
}

func (a *Arena) NewFalse() *Value { _ = "STUB: not implemented"; return nil }
