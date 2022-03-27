package gitlab

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/leandrorichard/graphql-gitlab-projects/internal/data"
	"github.com/pkg/errors"
)

// Config represents the configuration fields of the Gitlab client.
type Config struct {
	API string
}

// Gitlab represents a Gitlab client, and its configuration fields.
type Gitlab struct {
	config *Config
	cli    http.Client
}

// NewClient retrieves a pointer to a Gitlab client instance.
func NewClient(config *Config) (*Gitlab, error) {
	c := Gitlab{
		config: config,
		cli: http.Client{
			Timeout: time.Second * time.Duration(20),
		},
	}

	return &c, nil
}

// Response represents the response object from the Gitlab GraphQL service.
type Response struct {
	Data Data `json:"data"`
}

// Data represents the data field object from the Gitlab GraphQL service response.
type Data struct {
	Projects Projects `json:"projects"`
}

// Projects represents the projects field object from the Gitlab GraphQL service response.
type Projects struct {
	Nodes []Node `json:"nodes"`
}

// Node represents one node (in this case project) from the Gitlab GraphQL service response.
type Node struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	ForksCount  int     `json:"forksCount"`
}

// ProjectsList is the Gitlab implementation of the ProjectList interface,
// that queries the last N projects from the Gitlab GraphQL API.
func (g *Gitlab) ProjectsList(last int) ([]data.ProjectRecord, error) {
	// Defines the Gitlab GraphQL last projects query.
	lastAsString := strconv.Itoa(last)

	body := map[string]string{
		"query": `query last_projects($n: Int = ` + lastAsString + `) {
			projects(last:$n) {
				nodes {
					name
					description
					forksCount
				}
			}
		}`,
	}

	bodyAsString, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "marshalling body map content")
	}

	buffer := bytes.NewBufferString(string(bodyAsString))

	// Builds the HTTP request and adds the content type application json header.
	req, err := http.NewRequest("POST", g.config.API, buffer)
	if err != nil {
		return nil, errors.Wrap(err, "building request")
	}

	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
	}

	// Performs the HTTP POST request to the Gitlab GraphQL server and
	// parses the response into a Response struct.
	resp, err := g.cli.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "sending the request")
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading the body")
	}

	var dataResponse Response
	if err := json.Unmarshal(responseBody, &dataResponse); err != nil {
		return nil, errors.Wrap(err, "parsing the body into a struct")
	}

	// Converts the Response into a data.ProjectRecord list and returns it.
	var projects []data.ProjectRecord

	for _, node := range dataResponse.Data.Projects.Nodes {
		p := data.ProjectRecord{
			Name:        node.Name,
			Description: node.Description,
			ForksCount:  node.ForksCount,
		}

		projects = append(projects, p)
	}

	return projects, nil
}
