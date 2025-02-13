#!/bin/bash

IMAGE_REPOSITORY=${IMAGE_REPOSITORY:-khulnasoft}

docker build --build-arg AGENT_BINARY_BUILD_RELATIVE="$AGENT_BINARY_BUILD_RELATIVE" --network host --rm=true --tag=$IMAGE_REPOSITORY/khulnasoft_agent_ce:fargate-${KE_IMG_TAG:-latest} -f agent-binary/Dockerfile.fargate ..
