#!/bin/bash

if [[ -e .env ]];then
  export $(egrep -v '^#' .env | xargs)
fi

go run cmd/service/main.go