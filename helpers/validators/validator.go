package validators

import (
	"regexp"
	"strings"
)

// isStringEmpty checks if a string is empty.
func IsStringEmpty(input string) bool {
	return len(strings.TrimSpace(input)) == 0
}

// IsNilOrEmpty checks if any object of type defined by each case is nil or empty.
// Returns bool: true if nil or empty, false otherwise.
func IsNilOrEmpty(input interface{}) bool {
	switch obj := input.(type) {
	case nil:
		return true
	case string:
		return IsStringEmpty(obj)
	default:
		return false
	}
}


// IsValidID validates is the string is an uuid.
func IsValidID(id string) bool {
	regex := regexp.MustCompile(`([0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12})`)
	result := regex.FindStringSubmatch(id)
	if result == nil {
		return false
	}
	return true
}
