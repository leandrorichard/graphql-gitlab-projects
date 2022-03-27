package gitlab_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/leandrorichard/graphql-gitlab-projects/internal/sys/gitlab"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	mockServer := httpMockServer()
	defer mockServer.Close()

	config := gitlab.Config{
		API: mockServer.URL,
	}

	finder, err := gitlab.NewClient(&config)

	assert.NoError(t, err)
	assert.NotNil(t, finder)

	resp, err := finder.ProjectsList(3)

	assert.NoError(t, err)
	assert.True(t, len(resp) == 3)
}

func httpMockServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(projectsMockResponse))
	}))

	return ts
}

// projectsMockResponse it is a mock response used in the HTTP mock server.
const projectsMockResponse = `
{
  "data": {
    "projects": {
      "nodes": [
        {
          "name": "Heroes of Wesnoth",
          "description": null,
          "forksCount": 5
        },
        {
          "name": "Leiningen",
          "description": "",
          "forksCount": 1
        },
        {
          "name": "TearDownWalls",
          "description": null,
          "forksCount": 5
        }
      ]
    }
  }
}`
