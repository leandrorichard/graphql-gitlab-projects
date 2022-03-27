# graphql-gitlab-projects

It is a simple Golang application that queries the N last projects in the Gitlab GraphQL API.

It prints out the names of all returned projects' by a comma delimiter and the sum of all forks.

#### Environment variables
Set the `GITLAB_GRAPHQL_API_URL` in the `Dockerfile`. 
It is the URL of the Gitlab GraphQL API.

#### Run the application
```shell
make run
```

#### Run the tests
```shell
make test
```