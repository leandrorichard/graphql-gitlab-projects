package domain

import (
	"github.com/leandrorichard/graphql-gitlab-projects/internal/data"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// ListProjects receives the N last number of projects and a projects' finder parameters.
// It calls the data layer passing the parameters, join the projects' names with a
// comma delimiter, sum all forks and returns it.
func ListProjects(logger *zap.Logger, last int, finder data.ProjectsFinder) (string, int, error) {
	projects, err := data.ListProjects(last, finder)
	if err != nil {
		return "", 0, errors.Wrap(err, "listing projects from data layer")
	}

	logger.Info("projects retrieved with success from data layer",
		zap.Int("total projects", len(projects)))

	// Join the name of all projects and sum all forks.
	var names string
	var totalForks int

	for _, project := range projects {
		names = names + "," + project.Name
		totalForks = totalForks + project.ForksCount
	}

	// It returns from the index 1 to remove the first comma.
	return names[1:], totalForks, nil
}
