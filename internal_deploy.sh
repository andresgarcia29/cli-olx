#!/bin/bash
set -e
GOOS=darwin GOARCH=amd64 go build -o olx .
# Install the required packages

# Move the 'olx' file to /usr/local/bin with superuser privileges
mv olx /usr/local/bin