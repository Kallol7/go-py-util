package pyutil

import (
	"testing"
)

func TestFunctions(t *testing.T) {
	t.Run("Test formatting functions", func(t *testing.T) {
		// Test Len
		bigNum := int64(999999999999999999)
		l := Len(bigNum)
		if l != 18 {
			t.Errorf("Expected length 18, got %v", l)
		} else if Len(123) != 3 {
			t.Errorf("Expected length 3, got %v", l)
		} else if Len(99.123) != 6 {
			t.Errorf("Expected length 6, got %v", l)
		}

		// Test Len and Str
		x := [5]string{"2", "3", Str(Len(bigNum))}
		if x[2] != "18" {
			t.Errorf("Expected '18', got %v", x[2])
		}
	})
}
