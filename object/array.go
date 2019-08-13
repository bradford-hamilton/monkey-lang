package object

import (
	"bytes"
	"strings"
)

// Array type wraps the array's elements in a slice of Objects
type Array struct {
	Elements []Object
}

// Type returns our Array's ObjectType (ArrayObj)
func (a *Array) Type() ObjectType { return ArrayObj }

// Inspect returns a string representation of the Array's elements: [1, 2, 3]
func (a *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
