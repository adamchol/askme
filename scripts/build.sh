#!/bin/bash

go build -o ask

rm -f $GOPATH/bin/ask

mv ask $GOPATH/bin/
