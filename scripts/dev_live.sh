#!/bin/bash

if ! command -v nodemon &> /dev/null
then
    npm i -g nodemon
fi

nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run ./cmd/main.go
