package handlers

import (
        "github.com/stretchr/testify/assert"
        "net/http"
        "net/http/httptest"
        "testing"
)

func TestMethodMatcher(t *testing.T) {
        t.Parallel()
        handler := func(w http.ResponseWriter, req *http.Request) {
                http.StatusText(200)
        }
        handlerFunc := method_matcher("GET", handler)

        server := httptest.NewServer(handlerFunc)
        res, err := http.Post(server.URL, "", nil)
        assert.NoError(t, err)
        assert.Equal(t, 400, res.StatusCode)

        client := &http.Client{}
        req, err := http.NewRequest("PUT", server.URL, nil)
        res, err = client.Do(req)
        assert.NoError(t, err)
        assert.Equal(t, 400, res.StatusCode)

        req, err = http.NewRequest("DELETE", server.URL, nil)
        res, err = client.Do(req)
        assert.NoError(t, err)
        assert.Equal(t, 400, res.StatusCode)

        res, err = http.Get(server.URL)
        assert.NoError(t, err)
        assert.Equal(t, 200, res.StatusCode)
}
