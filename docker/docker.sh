#!/bin/bash

docker exec EPGStation-Notification bash -c 'cd /opt/src/ && GOOS=linux GOARCH=amd64 go run ../main.go'
