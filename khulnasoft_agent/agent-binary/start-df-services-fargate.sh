#!/bin/bash

set -x

# Some defaults
export DF_SERVERLESS=${DF_SERVERLESS:-"true"}
export DF_INSTALL_DIR=${DF_INSTALL_DIR:-"/opt/khulnasoft"}
export DF_BASE_DIR="/khulnasoft"

deep_copy() {
  #eval source=${1}
  mkdir -p "$(dirname "$2")" && cp -r $1 "$2"
}

setup_env_paths() {
  export DF_BIN_DIR="$KE_BASE_DIR/bin"
  export PATH=$PATH:$KE_BIN_DIR
  echo "PATH is $PATH"

  if [[ -z "${SCOPE_HOSTNAME}" ]]; then
    SCOPE_HOSTNAME="$(hostname)"
    echo "Got hostname: "
    echo "$SCOPE_HOSTNAME"
    export SCOPE_HOSTNAME="$SCOPE_HOSTNAME"
  fi
  if [ "$INSTANCE_ID_SUFFIX" == "Y" ]; then
    cloud_instance_id=$(getCloudInstanceId)
    cloud_instance_id="${cloud_instance_id//[$'\t\r\n ']/}"
    export SCOPE_HOSTNAME="$SCOPE_HOSTNAME-$cloud_instance_id"
  fi
  container_info=$(sed -n '/name=systemd/p' /proc/self/cgroup)
  # echo "Container info: $container_info"
  # container_id=$(echo $container_info | cut -d/ -f4 | cut -c1-12,31-)
  # echo "Container id: $container_id"
  # if [[ -z "${container_id}" ]]; then
  #  container_id=$(echo $container_info | cut -d/ -f3 | cut -c1-12)
  #  echo "updated container id: $container_id"
  # fi
  # export SCOPE_HOSTNAME="$SCOPE_HOSTNAME-$container_id"

  if [[ -z "${DF_INSTALL_DIR}" ]]; then
    export DF_INSTALL_DIR="$KE_BASE_DIR/df-agents/$SCOPE_HOSTNAME"
  else
    export DF_INSTALL_DIR="$KE_INSTALL_DIR/df-agents/$SCOPE_HOSTNAME"
  fi

  echo "Khulnasoft agent install dir: $KE_INSTALL_DIR"

  export PATH=$PATH:$KE_INSTALL_DIR/bin:$KE_INSTALL_DIR/usr/local/bin:$KE_INSTALL_DIR/home/khulnasoft:$KE_INSTALL_DIR/usr/bin
  echo "UPDATED PATH is $PATH"

  export FILEBEAT_CERT_PATH="$KE_INSTALL_DIR/etc/filebeat/filebeat.crt"

  export MGMT_CONSOLE_PORT=443
  export DF_TLS_ON="1"
  export SECRET_SCANNER_LD_LIBRARY_PATH=$KE_INSTALL_DIR/home/khulnasoft/lib:$LD_LIBRARY_PATH
}

trim() {
    local var="$*"
    # remove leading whitespace characters
    var="${var#"${var%%[![:space:]]*}"}"
    # remove trailing whitespace characters
    var="${var%"${var##*[![:space:]]}"}"
    echo -n "$var"
}

echoerr() { echo "$@" 1>&2; }

configure_cron() {
    #Setup cron jobs for misc tasks, it needs to be killed and restarted
    #doesnt work smoothly inside docker!
    service cron start
    chmod 600 $KE_INSTALL_DIR/etc/logrotate.d/fenced_logrotate.conf
    envsubst '${DF_INSTALL_DIR}' <$KE_INSTALL_DIR/etc/logrotate.d/fenced_logrotate.conf >$KE_INSTALL_DIR/etc/logrotate.d/fenced_logrotate_new.conf
    mv $KE_INSTALL_DIR/etc/logrotate.d/fenced_logrotate_new.conf $KE_INSTALL_DIR/etc/logrotate.d/fenced_logrotate.conf

    logrotate_path=$(which logrotate)
    (echo "*/5 * * * * $logrotate_path $KE_INSTALL_DIR/etc/logrotate.d/fenced_logrotate.conf") | crontab -
}

launch_khulnasoftd() {
    # In k8s, if agent pod restarts these files are not cleared
    rm -rf $KE_INSTALL_DIR/var/log/fenced/* 2>/dev/null
    mkdir -p $KE_INSTALL_DIR/tmp $KE_INSTALL_DIR/var/log/fenced/malware-scan $KE_INSTALL_DIR/var/log/fenced/malware-scan-log $KE_INSTALL_DIR/var/log/fenced/secret-scan $KE_INSTALL_DIR/var/log/fenced/secret-scan-log $KE_INSTALL_DIR/var/log/fenced/compliance $KE_INSTALL_DIR/var/log/fenced/compliance-scan-logs 2>/dev/null
    configure_cron
    if [[ -z "${SCOPE_HOSTNAME}" ]]; then
        SCOPE_HOSTNAME="$(hostname)"
        export SCOPE_HOSTNAME="$SCOPE_HOSTNAME"
    fi
    if [ "$INSTANCE_ID_SUFFIX" == "Y" ]; then
        cloud_instance_id=$($KE_INSTALL_DIR/usr/local/bin/getCloudInstanceId)
        cloud_instance_id="${cloud_instance_id//[$'\t\r\n ']/}"
        export SCOPE_HOSTNAME="$SCOPE_HOSTNAME-$cloud_instance_id"
    fi
    if [ "$KE_PROXY_MODE" == "1" ]; then
        # echo "App security : Active Mode, Listening on port $KE_LISTEN_PORT "
        DOCKER_API_VERSION=$DOCKER_API_VERSION run_dind.sh -a $MGMT_CONSOLE_PORT -s 0
    fi

    export PROBE_LOG_LEVEL=${DF_LOG_LEVEL:-info}

    envsubst '${KHULNASOFT_KEY}:${MGMT_CONSOLE_URL}:${MGMT_CONSOLE_PORT}:${MGMT_CONSOLE_URL_SCHEMA}:${SCOPE_HOSTNAME}:${DF_INSTALL_DIR}' <$KE_INSTALL_DIR/etc/td-agent-bit/fluentbit-agent.conf >$KE_INSTALL_DIR/etc/td-agent-bit/fluentbit-agent-new.conf
    mv $KE_INSTALL_DIR/etc/td-agent-bit/fluentbit-agent-new.conf $KE_INSTALL_DIR/etc/td-agent-bit/fluentbit-agent.conf

    envsubst '${KHULNASOFT_KEY}:${MGMT_CONSOLE_URL}:${MGMT_CONSOLE_PORT}:${MGMT_CONSOLE_URL_SCHEMA}:${SCOPE_HOSTNAME}:${DF_INSTALL_DIR}' <$KE_INSTALL_DIR/etc/td-agent-bit/fluentbit-cluster-agent.conf >$KE_INSTALL_DIR/etc/td-agent-bit/fluentbit-cluster-agent-new.conf
    mv $KE_INSTALL_DIR/etc/td-agent-bit/fluentbit-cluster-agent-new.conf $KE_INSTALL_DIR/etc/td-agent-bit/fluentbit-cluster-agent.conf

    chmod 600 $KE_INSTALL_DIR/etc/td-agent-bit/*

    echo "Starting agent..."
    echo "Khulnasoft agent install dir: $KE_INSTALL_DIR"
    echo "Khulnasoft agent base dir: $KE_BASE_DIR"
    echo "Khulnasoft agent hostname: $SCOPE_HOSTNAME"
    echo "Khulnasoft management console url: $MGMT_CONSOLE_URL"
    echo "Khulnasoft management console port: $MGMT_CONSOLE_PORT"
    # $KE_INSTALL_DIR/bin/khulnasoftd&
    # start in background using nohup
    mkdir -p "$KE_INSTALL_DIR/var/log/supervisor"
    touch "$KE_INSTALL_DIR/var/log/supervisor/khulnasoftd.log"

    echo "Starting agent ..."
    /bin/sh -c "ulimit -l unlimited; $KE_INSTALL_DIR/bin/khulnasoftd >> $KE_INSTALL_DIR/var/log/supervisor/khulnasoftd.log 2>&1"

}

create_cgroups() {
    #    echo "creating cgroups to perform resource-control"
    /bin/sh $KE_INSTALL_DIR/home/khulnasoft/create-cgroups.sh >/dev/null 2>&1
}

setup_env_paths

# Setup and install DF agent
if [[ "$KE_BASE_DIR" != "$KE_INSTALL_DIR" ]]; then
  mkdir -p $KE_INSTALL_DIR
  echo "Copying agent to DF installation dir"
  deep_copy "$KE_BASE_DIR/bin/*" "$KE_INSTALL_DIR/bin/."
  deep_copy "$KE_BASE_DIR/etc/*" "$KE_INSTALL_DIR/etc/."
  deep_copy "$KE_BASE_DIR/home/*" "$KE_INSTALL_DIR/home/."
  deep_copy "$KE_BASE_DIR/usr/*" "$KE_INSTALL_DIR/usr/."
  deep_copy "$KE_BASE_DIR/opt/*" "$KE_INSTALL_DIR/opt/."
fi

# Make rules.tar.gz
cd $KE_INSTALL_DIR/home/khulnasoft && $KE_INSTALL_DIR/bin/tar czvf rules.tar.gz rules
cd -

# To create fenced log and tmp directory
mkdir -p $KE_INSTALL_DIR/var/log/fenced/ $KE_INSTALL_DIR/var/log/supervisor/ $KE_INSTALL_DIR/var/log/fenced/coredump
chmod +x $KE_INSTALL_DIR/home/khulnasoft/*.sh

# for compliance change config
cat $KE_INSTALL_DIR/usr/local/bin/compliance_check/config.json | jq --arg dfinstalldir "$KE_INSTALL_DIR" '.[] |= . + { "files": [ $dfinstalldir + .files[] ] }' > modified_config.json
rm $KE_INSTALL_DIR/usr/local/bin/compliance_check/config.json
mv modified_config.json $KE_INSTALL_DIR/usr/local/bin/compliance_check/config.json

if [ "$KE_USE_DUMMY_SCOPE" == "" ]; then
    pidVal=$(/bin/pidof $KE_INSTALL_DIR/bin/khulnasoftd)
    if [ -n "$pidVal" ]; then
        echo "Warning: Another bootstrap is running."
    fi
    # create_cgroups
fi

echo "Start Khulnasoft services... Console is $MGMT_CONSOLE_URL"
launch_khulnasoftd

sleep 2

tail --follow=name $KE_INSTALL_DIR/var/log/supervisor/khulnasoftd.log &