package validators

import (
	"github.com/google/uuid"
	"testing"
)

// Test_IsValidID tests IsValidID
func Test_IsValidID(t *testing.T) {
	t.Run("valid id", func(t *testing.T) {
		if ok := IsValidID(uuid.New().String()); !ok {
			t.Error("expected valid ID, but was invalid")
		}
	})
	t.Run("invalid id", func(t *testing.T) {
		if ok := IsValidID("someId"); ok {
			t.Error("expected invalid ID, but was valid")
		}
	})
}
