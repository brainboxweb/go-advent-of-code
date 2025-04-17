package dinner_test

import (
	"testing"

	"github.com/brainboxweb/advent/day13/dinner"
)

func TestGetHappiness(t *testing.T) {
	expected := 54
	r := dinner.Relationships{}
	r.AddRelationship("Gary", "David", 54)
	actual := r.GetHappiness("Gary", "David")

	if actual != expected {
		t.Errorf("expected %d, got %d.", actual, expected)
	}
}
