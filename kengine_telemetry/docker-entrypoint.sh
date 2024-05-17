#!/bin/sh

if [ "$KHULNASOFT_TELEMETRY_ENABLED" = "true" ]; then
  /go/bin/all-in-one-linux
else
  echo "KHULNASOFT_TELEMETRY_ENABLED is not set to true"
  sleep infinity
fi
