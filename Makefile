VERSION=1.0.1

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run

init:
	$(GOCMD) mod download

test:
	$(GOTEST)  -v ./...

test-race:
	$(GOTEST) --race -v ./...

coverage:
	$(GOTEST) ./src/... -coverprofile cover.out
	go tool cover -html=./cover.out

coverage-badge:
	gopherbadger -md="README.md"

