package games

import (
    "testing"
)

var service = NewService()

func TestComputeFlames(t *testing.T) {
    result := service.ComputeFlames("John", "Jane")
    if result != "E" {
        t.Errorf("Expected E, got %s", result)
    }
}
