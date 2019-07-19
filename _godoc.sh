#!/usr/bin/env bash

PROJECT_DIR="$(pwd)"
PROJECT_NAME=$(basename $PROJECT_DIR)
UNAME=$1
TMP_PROJECT=/tmp/tmpgopath/src/github.com/$UNAME/$PROJECT_NAME

if [ -f "$(pwd)/go.mod" ]; then
  echo "$(pwd)/go.mod exist, generating docs"

  mkdir -p /tmp/tmpgoroot/doc
  rm -rf $TMP_PROJECT
  mkdir -p $TMP_PROJECT
  tar -c --exclude='.git' --exclude='tmp' . | tar -x -C $TMP_PROJECT
  echo -e "open http://localhost:6060/pkg/github.com/$UNAME/$PROJECT_NAME\n"
  GOROOT=/tmp/tmpgoroot/ GOPATH=/tmp/tmpgopath/ godoc -http=localhost:6060
else
  echo "$(pwd)/go.mod does not exist, bye!"
fi