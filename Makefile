
# Inspired by https://gist.github.com/sohlich/8432e7c1bd56bc395b101d1ba444e982
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=npmify

all:
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	echo "Testing not supported yet!"

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	make clean
	make build
	./$(BINARY_NAME)

get:
	$(GOGET)

