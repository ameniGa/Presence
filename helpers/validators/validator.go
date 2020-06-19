package validators

import "strings"

// isStringEmpty checks if a string is empty.
func isStringEmpty(input string) bool {
	return len(strings.TrimSpace(input)) == 0
}

// IsNilOrEmpty checks if any object of type defined by each case is nil or empty.
// Returns bool: true if nil or empty, false otherwise.
func IsNilOrEmpty(input interface{}) bool {
	switch obj := input.(type) {
	case nil:
		return true
	case string:
		return isStringEmpty(obj)
	default:
		return false
	}
	}
