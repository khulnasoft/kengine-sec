#!/bin/bash

while true
do
	/bin/sh -c "ulimit -l unlimited; $KE_INSTALL_DIR/bin/khulnasoftd >> $KE_INSTALL_DIR/var/log/supervisor/khulnasoftd.log 2>&1"
	sleep 5
done
