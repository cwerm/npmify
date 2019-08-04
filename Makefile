
# Inspired by https://gist.github.com/sohlich/8432e7c1bd56bc395b101d1ba444e982
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
SCRIPTS=scripts
GODOC=./$(SCRIPTS)/_godoc.sh
BINARY_NAME=npmify

all: test build run

clean-start: clean all

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	echo "Testing not supported yet!"

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	./$(BINARY_NAME)

deps:
	$(GOGET)

docs:
	$(GODOC)

