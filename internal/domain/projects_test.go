package domain_test

import (
	"os"
	"testing"

	"go.uber.org/zap"

	"github.com/leandrorichard/graphql-gitlab-projects/internal/data"
	"github.com/leandrorichard/graphql-gitlab-projects/internal/domain"
	"github.com/stretchr/testify/assert"
)

var logger *zap.Logger

func TestMain(t *testing.M) {
	// Initialize logger.
	logger, _ = zap.NewProduction()
	defer logger.Sync()

	os.Exit(t.Run())
}

// projectsFinderStub it is a dummy implementation of the projects' finder.
type projectsFinderStub struct{}

// ProjectsList it is a dummy implementation of the projects list finder service.
func (p *projectsFinderStub) ProjectsList(last int) ([]data.ProjectRecord, error) {
	dummyDesc := "Dummy description"

	resp := []data.ProjectRecord{
		{
			Name:        "Heroes of Wesnoth",
			Description: nil,
			ForksCount:  5,
		},
		{
			Name:        "Leiningen",
			Description: &dummyDesc,
			ForksCount:  1,
		},
	}

	return resp, nil
}

func TestListProjects(t *testing.T) {
	names, forks, err := domain.ListProjects(logger, 5, &projectsFinderStub{})

	assert.Equal(t, "Heroes of Wesnoth,Leiningen", names)
	assert.Equal(t, 6, forks)
	assert.Equal(t, nil, err)
}
