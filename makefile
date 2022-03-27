SHELL:=/bin/bash

run:
	docker build -t gitlab-projects . && docker run gitlab-projects

test:
	docker build -f Dockerfile.ci_test -t gitlab-projects-test . && docker run gitlab-projects-test