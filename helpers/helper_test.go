package helpers

import "testing"

func TestGetLogger(t *testing.T) {
	logger := GetLogger()
	if logger == nil {
		t.Errorf("expected a logger got nil")
	}
}
