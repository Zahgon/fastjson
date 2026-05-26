package fastjson

// Parser parses JSON.
//
// Parser may be re-used for subsequent parsing.
//
// Parser cannot be used from concurrent goroutines.
// Use per-goroutine parsers or ParserPool instead.
type Parser struct {
	// b contains working copy of the string to be parsed.
	b []byte

	// c is a cache for json values.
	c cache
}

// Parse parses s containing JSON.
//
// The returned value is valid until the next call to Parse*.
//
// Use Scanner if a stream of JSON values must be parsed.
func (p *Parser) Parse(s string) (*Value, error) { _ = "STUB: not implemented"; return nil, nil }

// ParseBytes parses b containing JSON.
//
// The returned Value is valid until the next call to Parse*.
//
// Use Scanner if a stream of JSON values must be parsed.
func (p *Parser) ParseBytes(b []byte) (*Value, error) { _ = "STUB: not implemented"; return nil, nil }

type cache struct {
	vs []Value
}

func (c *cache) reset() { _ = "STUB: not implemented"; return }

func (c *cache) getValue() *Value { _ = "STUB: not implemented"; return nil }

func skipWS(s string) string { _ = "STUB: not implemented"; return "" }

// Fast path.

func skipWSSlow(s string) string { _ = "STUB: not implemented"; return "" }

type kv struct {
	k string
	v *Value
}

// MaxDepth is the maximum depth for nested JSON.
const MaxDepth = 300

func (c *cache) parseValue(s string, depth int) (*Value, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// Try parsing NaN

func (c *cache) parseArray(s string, depth int) (*Value, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (c *cache) parseObject(s string, depth int) (*Value, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// Parse key.

// Parse value

func escapeString(dst []byte, s string) []byte { _ = "STUB: not implemented"; return nil }

// Fast path - nothing to escape.

// Slow path.

func hasSpecialChars(s string) bool { _ = "STUB: not implemented"; return false }

func unescapeStringBestEffort(s string) string { _ = "STUB: not implemented"; return "" }

// Fast path - nothing to unescape.

// Slow path - unescape string.
// It is safe to do, since s points to a byte slice in Parser.b.

// Too short escape sequence. Just store it unchanged.

// Invalid escape sequence. Just store it unchanged.

// Surrogate.
// See https://en.wikipedia.org/wiki/Universal_Character_Set_characters#Surrogates

// Unknown escape sequence. Just store it unchanged.

// parseRawKey is similar to parseRawString, but is optimized
// for small-sized keys without escape sequences.
func parseRawKey(s string) (string, string, error) { _ = "STUB: not implemented"; return "", "", nil }

// Fast path.

// Slow path.

func parseRawString(s string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// Fast path. No escaped ".

// Slow path - possible escaped " found.

func parseRawNumber(s string) (string, string, error) {
	_ = "STUB: not implemented"
	// The caller must ensure len(s) > 0
	return "", "", nil
}

// Find the end of the number.

// Object represents JSON object.
//
// Object cannot be used from concurrent goroutines.
// Use per-goroutine parsers or ParserPool instead.
type Object struct {
	kvs           []kv
	keysUnescaped bool
}

func (o *Object) reset() {
	_ = "STUB: not implemented"
	// o.kvs entries can point to external byte slices. Clear these references, so GC could free memory.
	return
}

// MarshalTo appends marshaled o to dst and returns the result.
func (o *Object) MarshalTo(dst []byte) []byte { _ = "STUB: not implemented"; return nil }

// String returns string representation for the o.
//
// This function is for debugging purposes only. It isn't optimized for speed.
// See MarshalTo instead.
func (o *Object) String() string { _ = "STUB: not implemented"; return "" }

// It is safe converting b to string without allocation, since b is no longer
// reachable after this line.

func (o *Object) getKV() *kv { _ = "STUB: not implemented"; return nil }

func (o *Object) unescapeKeys() { _ = "STUB: not implemented"; return }

// Len returns the number of items in the o.
func (o *Object) Len() int {
	_ = "STUB: not implemented"

	// Get returns the value for the given key in the o.
	//
	// Returns nil if the value for the given key isn't found.
	//
	// The returned value is valid until Parse is called on the Parser returned o.
	return 0
}

func (o *Object) Get(key string) *Value { _ = "STUB: not implemented"; return nil }

// Fast path - try searching for the key without object keys unescaping.

// Slow path - unescape object keys.

// Visit calls f for each item in the o in the original order
// of the parsed JSON.
//
// f cannot hold key and/or v after returning.
func (o *Object) Visit(f func(key []byte, v *Value)) { _ = "STUB: not implemented"; return }

// Value represents any JSON value.
//
// Call Type in order to determine the actual type of the JSON value.
//
// Value cannot be used from concurrent goroutines.
// Use per-goroutine parsers or ParserPool instead.
type Value struct {
	o Object
	a []*Value
	s string
	t Type
}

func (v *Value) reset() { _ = "STUB: not implemented"; return }

// MarshalTo appends marshaled v to dst and returns the result.
func (v *Value) MarshalTo(dst []byte) []byte { _ = "STUB: not implemented"; return nil }

// String returns string representation of the v.
//
// The function is for debugging purposes only. It isn't optimized for speed.
// See MarshalTo instead.
//
// Don't confuse this function with StringBytes, which must be called
// for obtaining the underlying JSON string for the v.
func (v *Value) String() string { _ = "STUB: not implemented"; return "" }

// It is safe converting b to string without allocation, since b is no longer
// reachable after this line.

// Type represents JSON type.
type Type int

const (
	// TypeNull is JSON null.
	TypeNull Type = 0

	// TypeObject is JSON object type.
	TypeObject Type = 1

	// TypeArray is JSON array type.
	TypeArray Type = 2

	// TypeString is JSON string type.
	TypeString Type = 3

	// TypeNumber is JSON number type.
	TypeNumber Type = 4

	// TypeTrue is JSON true.
	TypeTrue Type = 5

	// TypeFalse is JSON false.
	TypeFalse Type = 6

	typeRawString Type = 7
)

// String returns string representation of t.
func (t Type) String() string { _ = "STUB: not implemented"; return "" }

// typeRawString is skipped intentionally,
// since it shouldn't be visible to user.

// Type returns the type of the v.
func (v *Value) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

// Exists returns true if the field exists for the given keys path.
//
// Array indexes may be represented as decimal numbers in keys.
func (v *Value) Exists(keys ...string) bool { _ = "STUB: not implemented"; return false }

// Get returns value by the given keys path.
//
// Array indexes may be represented as decimal numbers in keys.
//
// nil is returned for non-existing keys path.
//
// The returned value is valid until Parse is called on the Parser returned v.
func (v *Value) Get(keys ...string) *Value { _ = "STUB: not implemented"; return nil }

// GetObject returns object value by the given keys path.
//
// Array indexes may be represented as decimal numbers in keys.
//
// nil is returned for non-existing keys path or for invalid value type.
//
// The returned object is valid until Parse is called on the Parser returned v.
func (v *Value) GetObject(keys ...string) *Object { _ = "STUB: not implemented"; return nil }

// GetArray returns array value by the given keys path.
//
// Array indexes may be represented as decimal numbers in keys.
//
// nil is returned for non-existing keys path or for invalid value type.
//
// The returned array is valid until Parse is called on the Parser returned v.
func (v *Value) GetArray(keys ...string) []*Value { _ = "STUB: not implemented"; return nil }

// GetFloat64 returns float64 value by the given keys path.
//
// Array indexes may be represented as decimal numbers in keys.
//
// 0 is returned for non-existing keys path or for invalid value type.
func (v *Value) GetFloat64(keys ...string) float64 { _ = "STUB: not implemented"; return 0 }

// GetInt returns int value by the given keys path.
//
// Array indexes may be represented as decimal numbers in keys.
//
// 0 is returned for non-existing keys path or for invalid value type.
func (v *Value) GetInt(keys ...string) int { _ = "STUB: not implemented"; return 0 }

// GetUint returns uint value by the given keys path.
//
// Array indexes may be represented as decimal numbers in keys.
//
// 0 is returned for non-existing keys path or for invalid value type.
func (v *Value) GetUint(keys ...string) uint { _ = "STUB: not implemented"; return 0 }

// GetInt64 returns int64 value by the given keys path.
//
// Array indexes may be represented as decimal numbers in keys.
//
// 0 is returned for non-existing keys path or for invalid value type.
func (v *Value) GetInt64(keys ...string) int64 { _ = "STUB: not implemented"; return 0 }

// GetUint64 returns uint64 value by the given keys path.
//
// Array indexes may be represented as decimal numbers in keys.
//
// 0 is returned for non-existing keys path or for invalid value type.
func (v *Value) GetUint64(keys ...string) uint64 { _ = "STUB: not implemented"; return 0 }

// GetStringBytes returns string value by the given keys path.
//
// Array indexes may be represented as decimal numbers in keys.
//
// nil is returned for non-existing keys path or for invalid value type.
//
// The returned string is valid until Parse is called on the Parser returned v.
func (v *Value) GetStringBytes(keys ...string) []byte { _ = "STUB: not implemented"; return nil }

// GetBool returns bool value by the given keys path.
//
// Array indexes may be represented as decimal numbers in keys.
//
// false is returned for non-existing keys path or for invalid value type.
func (v *Value) GetBool(keys ...string) bool { _ = "STUB: not implemented"; return false }

// Object returns the underlying JSON object for the v.
//
// The returned object is valid until Parse is called on the Parser returned v.
//
// Use GetObject if you don't need error handling.
func (v *Value) Object() (*Object, error) { _ = "STUB: not implemented"; return nil, nil }

// Array returns the underlying JSON array for the v.
//
// The returned array is valid until Parse is called on the Parser returned v.
//
// Use GetArray if you don't need error handling.
func (v *Value) Array() ([]*Value, error) { _ = "STUB: not implemented"; return nil, nil }

// StringBytes returns the underlying JSON string for the v.
//
// The returned string is valid until Parse is called on the Parser returned v.
//
// Use GetStringBytes if you don't need error handling.
func (v *Value) StringBytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Float64 returns the underlying JSON number for the v.
//
// Use GetFloat64 if you don't need error handling.
func (v *Value) Float64() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

// Int returns the underlying JSON int for the v.
//
// Use GetInt if you don't need error handling.
func (v *Value) Int() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Uint returns the underlying JSON uint for the v.
//
// Use GetInt if you don't need error handling.
func (v *Value) Uint() (uint, error) { _ = "STUB: not implemented"; return 0, nil }

// Int64 returns the underlying JSON int64 for the v.
//
// Use GetInt64 if you don't need error handling.
func (v *Value) Int64() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Uint64 returns the underlying JSON uint64 for the v.
//
// Use GetInt64 if you don't need error handling.
func (v *Value) Uint64() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// Bool returns the underlying JSON bool for the v.
//
// Use GetBool if you don't need error handling.
func (v *Value) Bool() (bool, error) { _ = "STUB: not implemented"; return false, nil }

var (
	valueTrue  = &Value{t: TypeTrue}
	valueFalse = &Value{t: TypeFalse}
	valueNull  = &Value{t: TypeNull}
)
