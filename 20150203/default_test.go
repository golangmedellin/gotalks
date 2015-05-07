package main

import (
    "encoding/json"
    "github.com/sescobb27/meetup/models"
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
)

// START OMIT
func TestGetCategories(t *testing.T) {
    t.Parallel()
    recorder := httptest.NewRecorder()
    req, err := http.NewRequest("GET", "/categories", nil)
    GetCategories(recorder, req)
    assert.NoError(t, err)
    assert.Equal(t, 200, recorder.Code)
    assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
    categories := []*models.Category{}
    body, err := ioutil.ReadAll(recorder.Body)
    assert.NoError(t, err)
    err = json.Unmarshal(body, &categories)
    assert.NoError(t, err)
    assert.NotEmpty(t, categories)
}

// END OMIT
