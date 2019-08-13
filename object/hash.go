package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"
)

// Hashable is one method called HashKey. Any object that that can be used as a HashKey
// must implement this interface (*object.String, *object.boolean, *object.integer)
type Hashable interface {
	HashKey() HashKey
}

// HashKey type wraps the key's type and holds its value
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// HashKey returns a HashKey with a Value of 1 or 0 (true or false) and a Type of BooleanObj
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{
		Type:  b.Type(),
		Value: value,
	}
}

// HashKey returns a HashKey with a Value of the Integer and a Type of IntegerObj
func (i *Integer) HashKey() HashKey {
	return HashKey{
		Type:  i.Type(),
		Value: uint64(i.Value),
	}
}

// HashKey returns a HashKey with a Value of a 64-bit FNV-1a hash of the String and a Type of StringObj
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{
		Type:  s.Type(),
		Value: h.Sum64(),
	}
}

// HashPair holds key value pairs
type HashPair struct {
	Key   Object
	Value Object
}

// Hash hold Pairs which are a map of HashKey -> HashPair. We map the keys to HashPairs instead
// of Objects for a better ability to Inspect() and see the key and value
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Type returns Hash's ObjectType (HashObj)
func (h *Hash) Type() ObjectType { return HashObj }

// Inspect returns a string representation of the Hash
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(
			pairs,
			fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()),
		)
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
