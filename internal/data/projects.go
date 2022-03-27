package data

import "github.com/pkg/errors"

// ProjectsFinder it is an interface that defines the projects' finder resource.
type ProjectsFinder interface {
	ProjectsList(int) ([]ProjectRecord, error)
}

// ProjectRecord represents one project data.
type ProjectRecord struct {
	Name        string
	Description *string
	ForksCount  int
}

// ListProjects retrieves the list of N last projects.
// It calls the projects' finder to retrieve the list of projects and returns it.
func ListProjects(last int, finder ProjectsFinder) ([]ProjectRecord, error) {
	resp, err := finder.ProjectsList(last)
	if err != nil {
		return nil, errors.Wrap(err, "finder listing projects")
	}

	return resp, nil
}
