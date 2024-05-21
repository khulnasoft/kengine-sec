#!/bin/bash

IMAGE_REPOSITORY=${IMAGE_REPOSITORY:-khulnasoft}

rm -rf $AGENT_BINARY_BUILD/*

wget https://khulnasoft-public.s3.amazonaws.com/Kengine/agent-sensor/v2.1.0/cc11435d-bf5f-4a16-8c92-0a5a27e06b92/khulnasoft-agent-2.tar.gz

tar -zxvf khulnasoft-agent-2.tar.gz -C $AGENT_BINARY_BUILD/
rm -rf khulnasoft-agent-2.tar.gz

docker build --build-arg AGENT_BINARY_BUILD_RELATIVE="$AGENT_BINARY_BUILD_RELATIVE" --network host --rm=true --tag=$IMAGE_REPOSITORY/khulnasoft_agent_ce:fargate-${KE_IMG_TAG:-latest} -f agent-binary/Dockerfile.fargate ..
