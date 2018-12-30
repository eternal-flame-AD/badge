#!/bin/bash

export GO111MODULE=on 

echo "Building functions"
DIR=`pwd`
shopt -s dotglob
find functions/src/* -prune -type d | while IFS= read -r d; do
    cd $d
    echo "Entering $d"
    if [ -f go.mod ]; then
        echo "Found go.mod, installing modules"
        go get
    fi
    PKG=$(basename $d)
    go build -o "../../dist/$PKG"
    cd $DIR
done
