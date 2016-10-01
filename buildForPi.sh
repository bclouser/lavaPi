#!/usr/bin/env bash

env GOOS=linux GOARCH=arm GOARM=5 go build -v github.com/bclouser/lavaPi
