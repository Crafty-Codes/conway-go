package logic

import (
	"testing"

	"github.com/Crafty-Codes/conway-go/internal/object"
)

func TestGetSurvivalCount(t *testing.T) {
	space := object.NewSpace(3, 3)

	// Test what happens if row and column are at upper left corner
	if getSurvivalCount(space, 0, 0) != 0 {
		t.Error("Upper left corner hat survival neighbours")
	}

	// Test what happens if row and column are at lower right corner
	if getSurvivalCount(space, 2, 2) != 0 {
		t.Error("Lower rightcorner hat survival neighbours")
	}

	// Reviving block in center
	space[1][1].Vitality = object.ALIVE

	if getSurvivalCount(space, 0, 0) != 1 {
		t.Error("Neightbour didn't apear")
	}

	// Test that own block doesn't counts as Neighbour
	if getSurvivalCount(space, 1, 1) != 0 {
		t.Error("Block is it's own neighbour")
	}
}
