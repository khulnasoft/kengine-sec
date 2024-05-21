#!/bin/bash

IMAGE_REPOSITORY=${IMAGE_REPOSITORY:-khulnasoft}
KE_IMG_TAG=${KE_IMG_TAG:-latest}

building_image(){

    echo "Building GetCloudInstanceId"
    docker run --rm --workdir /go/src/github.com/khulnasoft/khulnasoft_agent -v $(pwd)/../khulnasoft_utils:/go/src/github.com/khulnasoft/khulnasoft_utils -v $(pwd):/go/src/github.com/khulnasoft/khulnasoft_agent:rw -v $(pwd):/go/src/github.com/khulnasoft/khulnasoft_agent:rw --net=host $IMAGE_REPOSITORY/khulnasoft_builder_ce:$KE_IMG_TAG bash -x /home/khulnasoft/gocode-build.sh
    build_result=$?
    if [ $build_result -ne 0 ]
    then
        echo "Khulnasoft code compilation failed, bailing out"
        exit 1
    fi

    echo "Building Shipper Executable"
    docker run --rm --workdir /go/src/github.com/khulnasoft/khulnasoft_agent -v $(pwd)/../golang_sdk:/go/src/github.com/khulnasoft/golang_sdk -v $(pwd)/../khulnasoft_utils:/go/src/github.com/khulnasoft/khulnasoft_utils -v $(pwd):/go/src/github.com/khulnasoft/khulnasoft_agent:rw --net=host -e VERSION=${VERSION} $IMAGE_REPOSITORY/khulnasoft_builder_ce:$KE_IMG_TAG bash -c "cd plugins/khulnasoft_shipper && make"
    build_result=$?
    if [ $build_result -ne 0 ]
    then
        echo "Shipper executable build failed, bailing out"
        exit 1
    fi

    echo "Building Agent Executable"
    docker run --rm --workdir /go/src/github.com/khulnasoft/khulnasoft_agent -v $(pwd)/../golang_sdk:/go/src/github.com/khulnasoft/golang_sdk -v $(pwd)/../khulnasoft_utils:/go/src/github.com/khulnasoft/khulnasoft_utils -v $(pwd):/go/src/github.com/khulnasoft/khulnasoft_agent:rw --net=host -e VERSION=${VERSION} $IMAGE_REPOSITORY/khulnasoft_builder_ce:$KE_IMG_TAG bash -c "make discovery"
    build_result=$?
    if [ $build_result -ne 0 ]
    then
        echo "Agent executable build failed, bailing out"
        exit 1
    fi

    echo "Building Cluster Agent Image"
    docker build --network host --rm=true --tag=$IMAGE_REPOSITORY/khulnasoft_cluster_agent_ce:$KE_IMG_TAG -f Dockerfile.cluster-agent .
    build_result=$?
    if [ $build_result -ne 0 ]
    then
        echo "Khulnasoft cluster agent building failed, bailing out"
        exit 1
    fi

    echo "Building Agent Image"
    docker build --network host --rm=true --build-arg KE_IMG_TAG="${KE_IMG_TAG}" --build-arg IMAGE_REPOSITORY="${IMAGE_REPOSITORY}" --tag=$IMAGE_REPOSITORY/khulnasoft_agent_ce:$KE_IMG_TAG -f Dockerfile .
    build_result=$?
    if [ $build_result -ne 0 ]
    then
        echo "Khulnasoft agent building failed, bailing out"
        exit 1
    fi
}


main () {
    building_image "$@"
}
main "$@"
