package main

import (
	"os"

	"github.com/leandrorichard/graphql-gitlab-projects/internal/domain"
	"github.com/leandrorichard/graphql-gitlab-projects/internal/sys/gitlab"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger.
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Initialize the Gitlab client to be called in the data layer to retrieve the list of N last projects.
	config := gitlab.Config{
		API: envMust(logger, "GITLAB_GRAPHQL_API_URL"),
	}

	finder, err := gitlab.NewClient(&config)
	if err != nil {
		logger.Fatal("instantiating gitlab client",
			zap.Any("config", &config),
			zap.Any("error", err))
	}

	// Calls the list projects to retrieve the list of projects and prints out the result,
	// which are the names joined by a comma and the sum of all forks.
	last := 5

	names, sumOfAllForks, err := domain.ListProjects(logger, last, finder)
	if err != nil {
		logger.Fatal("calling domain list projects",
			zap.Any("error", err))
	}

	logger.Info("projects retrieved with success",
		zap.String("join names", names),
		zap.Int("sum of all forks", sumOfAllForks))
}

// envMust fetches and returns the given env variable.
// If the variable is an empty string, it will log and Fatal.
func envMust(logger *zap.Logger, name string) string {
	value := os.Getenv(name)
	if value == "" {
		logger.Fatal("missing environment variable",
			zap.String("var name", name))
	}

	return value
}
