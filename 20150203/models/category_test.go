package models

import (
        "github.com/stretchr/testify/assert"
        "testing"
)

func TestGetCategories(t *testing.T) {
        categories := GetCategories()
        assert.NotEmpty(t, categories)
}
