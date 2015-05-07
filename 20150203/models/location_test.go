package models

import (
        "github.com/stretchr/testify/assert"
        "testing"
)

func TestGetLocations(t *testing.T) {
        locations := GetLocations()
        assert.NotEmpty(t, locations)
}
