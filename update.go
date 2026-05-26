package fastjson

// Del deletes the entry with the given key from o.
func (o *Object) Del(key string) { _ = "STUB: not implemented"; return }

// Fast path - try searching for the key without object keys unescaping.

// Slow path - unescape object keys before item search.

// Del deletes the entry with the given key from array or object v.
func (v *Value) Del(key string) { _ = "STUB: not implemented"; return }

// Set sets (key, value) entry in the o.
//
// The value must be unchanged during o lifetime.
func (o *Object) Set(key string, value *Value) { _ = "STUB: not implemented"; return }

// Try substituting already existing entry with the given key.

// Add new entry.

// Set sets (key, value) entry in the array or object v.
//
// The value must be unchanged during v lifetime.
func (v *Value) Set(key string, value *Value) { _ = "STUB: not implemented"; return }

// SetArrayItem sets the value in the array v at idx position.
//
// The value must be unchanged during v lifetime.
func (v *Value) SetArrayItem(idx int, value *Value) { _ = "STUB: not implemented"; return }
