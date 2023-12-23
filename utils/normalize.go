package utils

import (
	"regexp"
	"strings"
)

// NormalizeType provides methods for normalizing type names.
type NormalizeType struct{}

// NormalizationInfo represents the result of a type normalization.
// It includes the normalized type name and a flag indicating whether
// the type name was actually normalized.
type NormalizationInfo struct {
	TypeName   string
	Normalized bool
}

// NewNormalizeType creates and returns a new NormalizeType instance.
func NewNormalizeType() *NormalizeType {
	return &NormalizeType{}
}

// Normalize attempts to normalize the given type name.
// It returns the NormalizationInfo which contains the normalized type name and a flag indicating
// if the provided type name was actually normalized.
func (n *NormalizeType) Normalize(typeName string) NormalizationInfo {
	normalizedTypeName, isNormalized := n.normalizeTypeNameWithStatus(typeName)
	return NormalizationInfo{
		TypeName:   normalizedTypeName,
		Normalized: isNormalized,
	}
}

// isBuiltInType checks if the provided type name is one of the recognized built-in types.
func (n *NormalizeType) isBuiltInType(typeName string) bool {
	cases := []string{
		"uint", "int", "bool", "bytes", "string", "address", "addresspayable", "tuple", "enum", "error",
	}

	for _, bcase := range cases {
		if strings.Contains(typeName, bcase) {
			return true
		}
	}

	return false
}

// normalizeTypeNameWithStatus attempts to normalize the provided type name.
// It returns the normalized type name and a flag indicating if it was normalized.
func (n *NormalizeType) normalizeTypeNameWithStatus(typeName string) (string, bool) {
	isArray, _ := regexp.MatchString(`\[\d+\]`, typeName)
	isSlice := strings.HasPrefix(typeName, "[]")
	isSliceRight := strings.HasSuffix(typeName, "[]")

	if isArray || isSlice || isSliceRight {
		if !n.isBuiltInType(typeName) {
			return typeName, false
		}

		switch {
		case isArray:
			numberPart := typeName[strings.Index(typeName, "[")+1 : strings.Index(typeName, "]")]
			typePart := typeName[strings.Index(typeName, "]")+1:]
			normalizedTypeName, found := n.normalizeTypeName(typePart)
			return "[" + numberPart + "]" + normalizedTypeName, found

		case isSlice:
			typePart := typeName[2:]
			normalizedTypeName, found := n.normalizeTypeName(typePart)
			return "[]" + normalizedTypeName, found

		case isSliceRight:
			typePart := typeName[:len(typeName)-2]
			normalizedTypeName, found := n.normalizeTypeName(typePart)
			return normalizedTypeName + "[]", found
		}
	}

	switch {
	case strings.HasPrefix(typeName, "uint"):
		if typeName == "uint" {
			return "uint256", true
		}
		return typeName, true

	case strings.HasPrefix(typeName, "int"):
		if typeName == "int" {
			return "int256", true
		}
		return typeName, true

	case strings.HasPrefix(typeName, "enum"):
		return "uint8", true

	case strings.HasPrefix(typeName, "bool"):
		return typeName, true

	case strings.HasPrefix(typeName, "bytes"):
		return typeName, true

	case typeName == "string":
		return "string", true

	case typeName == "address":
		return "address", true

	case typeName == "addresspayable":
		return "address", true

	case typeName == "tuple":
		return "tuple", true

	default:
		return typeName, false // Custom struct types or unrecognized types are not considered normalized
	}
}

// normalizeTypeName provides the normalized version of the provided type name.
func (n *NormalizeType) normalizeTypeName(typeName string) (string, bool) {
	normalizedTypeName, found := n.normalizeTypeNameWithStatus(typeName)
	return normalizedTypeName, found
}
