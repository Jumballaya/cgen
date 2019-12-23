# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
BINARY_NAME = cgen

all: test build
build:
	$(GOBUILD) -o dist/$(BINARY_NAME) -v
test:
	$(GOTEST) ./...
cover:
	$(GOTEST) ./... -cover
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o dist/$(BINARY_NAME) -v
	./dist/$(BINARY_NAME)
deps:
	echo "No dependencies..."