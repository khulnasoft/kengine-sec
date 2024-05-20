#!/bin/bash

echo "Building Khulnasoft Go binaries"
make gocode
build_result=$?
if [ $build_result -ne 0 ]
then
    exit 1
fi
