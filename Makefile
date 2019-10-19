# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
BINARY_NAME=leaderboard

all: test build run
  build:
				$(GOBUILD) -o $(BINARY_NAME) -v
  test:
				$(GOTEST) -v ./...
  clean:
				$(GOCLEAN)
				rm -f $(BINARY_NAME)
				rm -f $(BINARY_UNIX)
  run:
				$(GOBUILD) -o $(BINARY_NAME) -v ./...
				./$(BINARY_NAME)

