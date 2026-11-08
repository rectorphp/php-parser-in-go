package ast

import "reflect"

// AttributeBag is embedded in every AST node to carry metadata assigned after
// parsing, such as a resolved type. Attributes are keyed by string and allocated
// lazily, so a node that never gets one costs nothing.
type AttributeBag struct {
	attributes map[string]any
}

// Equal lets go-cmp compare nodes that embed a bag without reaching its
// unexported field; two bags are equal when they hold the same attributes.
func (bag AttributeBag) Equal(other AttributeBag) bool {
	return reflect.DeepEqual(bag.attributes, other.attributes)
}

// SetAttribute stores value under key, allocating the backing map on first use.
func (bag *AttributeBag) SetAttribute(key string, value any) {
	if bag.attributes == nil {
		bag.attributes = map[string]any{}
	}
	bag.attributes[key] = value
}

// GetAttribute returns the value stored under key and whether it was present.
func (bag *AttributeBag) GetAttribute(key string) (any, bool) {
	value, ok := bag.attributes[key]
	return value, ok
}
