package ast

import "strings"

// ToString returns the source text of a name node - a Name, NameFullyQualified,
// or NameRelative joined by "\", a NamePart or Identifier value, or "" for any
// other node. The leading separator of a fully qualified name is not included.
func ToString(node Vertex) string {
	switch n := node.(type) {
	case *Name:
		return partsToString(n.Parts)
	case *NameFullyQualified:
		return partsToString(n.Parts)
	case *NameRelative:
		return partsToString(n.Parts)
	case *NamePart:
		return string(n.Value)
	case *Identifier:
		return string(n.Value)
	default:
		return ""
	}
}

func partsToString(parts []Vertex) string {
	segments := make([]string, 0, len(parts))
	for _, part := range parts {
		if namePart, ok := part.(*NamePart); ok {
			segments = append(segments, string(namePart.Value))
		}
	}
	return strings.Join(segments, "\\")
}

// reservedTypes are the PHP reserved type keywords that are not class names.
var reservedTypes = map[string]bool{
	"int": true, "float": true, "string": true, "bool": true, "void": true,
	"array": true, "iterable": true, "callable": true, "object": true,
	"mixed": true, "never": true, "null": true, "false": true, "true": true,
	"self": true, "static": true, "parent": true,
}

// IsReservedType reports whether name is a PHP reserved type keyword (a scalar,
// array, callable, void, never, mixed, or a self/static/parent relative type),
// matched case-insensitively. Such a name is a builtin type, not a class
// reference, even though the parser produces a Name node for it.
func IsReservedType(name string) bool {
	return reservedTypes[strings.ToLower(name)]
}
