#!/bin/bash

if  [ "$KE_LOG_LEVEL" == "debug" ]; then
    set -x
fi

echo "Starting discovery...."

# DF_BASE_DIR="/khulnasoft"
# export PATH=$PATH:$KE_BASE_DIR/bin/
chmod +x $KE_INSTALL_DIR/usr/local/discovery/khulnasoft-discovery
chmod +x $KE_INSTALL_DIR/home/khulnasoft/run_discovery*

until bash -x $KE_INSTALL_DIR/home/khulnasoft/run_discovery.sh 
do
    echo "Discovery crashed with exit code $?.  Restarting discovery...."
    sleep 5
done

echo "Discovery exited with exit code $?."


