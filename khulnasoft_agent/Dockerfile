ARG DF_IMG_TAG=latest
ARG IMAGE_REPOSITORY=khulnasoft

FROM $IMAGE_REPOSITORY/khulnasoft_secret_scanner_ce:$KE_IMG_TAG AS secret_build
FROM $IMAGE_REPOSITORY/khulnasoft_package_scanner_ce:$KE_IMG_TAG AS package_build
FROM $IMAGE_REPOSITORY/khulnasoft_malware_scanner_ce:$KE_IMG_TAG AS malware_build
FROM $IMAGE_REPOSITORY/khulnasoft_compliance_scanner_ce:$KE_IMG_TAG AS compliance_build

FROM debian:12-slim

MAINTAINER Khulnasoft Inc
LABEL khulnasoft.role=system

ENV CHECKPOINT_DISABLE=true \
    DOCKERVERSION=24.0.2 \
    DF_TLS_ON="1" \
    MGMT_CONSOLE_PORT=443 \
    DF_KUBERNETES_ON="N" \
    MGMT_CONSOLE_URL_SCHEMA=https \
    KHULNASOFT_KEY="" \
    MGMT_CONSOLE_URL_INTERNAL=127.0.0.1 \
    MGMT_CONSOLE_PORT_INTERNAL=8081

RUN export LD_LIBRARY_PATH="/usr/local/lib:$LD_LIBRARY_PATH" \
    && mkdir -p /usr/share/man/man1 /usr/share/man/man2 /usr/share/man/man3 /usr/share/man/man4 /usr/share/man/man5 /usr/share/man/man6 /usr/share/man/man7 /usr/share/man/man8 \
    && echo "Installing some basic stuff"

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    libpcap0.8 \
    gettext \
    ca-certificates \
    supervisor \
    logrotate \
    util-linux \
    dnsutils \
    net-tools \
    cgroup-tools \
    libcap2 \
    libaudit1 \
    conntrack \
    runit \
    auditd \
    apparmor \
    gzip \
    lsof \
    file \
    curl \
    zip \
    at \
    gnupg \
    unzip \
    procps \
    cron \
    sudo \
    bzip2 \
    libssl3 \
    libevent-2.1-7 \
    libevent-openssl-2.1-7 \
    libevent-pthreads-2.1-7 \
    libnet1 \
    gnupg2 \
    libfile-mimeinfo-perl \
    libjansson4 \
    libmagic1 \
    wget \
    bash \
    python3-pip \
    libvectorscan5 \
    skopeo \
    podman && \
    apt-get clean && \
    apt-get -y autoremove && \
    rm -rf /var/lib/apt/lists/*

ARG TARGETARCH

RUN <<EOF
set -eux

if [ "$TARGETARCH" = "arm64" ]; then
    ARCHITECTURE="aarch64"
elif [ "$TARGETARCH" = "amd64" ]; then
    ARCHITECTURE="x86_64"
else
    echo "Unsupported architecture $TARGETARCH" && exit 1;
fi

curl -fsSLO https://download.docker.com/linux/static/stable/${ARCHITECTURE}/docker-${DOCKERVERSION}.tgz
tar xzvf docker-${DOCKERVERSION}.tgz --strip 1 -C /usr/local/bin docker/docker
rm docker-${DOCKERVERSION}.tgz

mkdir -p /etc/license/ /usr/local/bin /usr/local/lib \
    /khulnasoftd /var/tmp/layers /usr/local/lua-waf /var/log/nginx/
chown root:root /khulnasoftd && chmod 0744 /khulnasoftd
mkdir /usr/local/bin/compliance_check && mkdir /usr/local/discovery

EOF

COPY tools/apache/khulnasoft/ke-utils/get_cloud_instance_id/getCloudInstanceId /usr/local/bin/getCloudInstanceId
COPY etc/fenced_logrotate.conf /etc/logrotate.d/fenced_logrotate.conf
COPY etc/certs/* /etc/filebeat/
COPY start_agent.sh /usr/local/bin/start_agent
COPY tools/apache/scope/docker/discovery /usr/local/discovery/khulnasoft-discovery
COPY plugins/compliance/scripts /usr/local/bin/compliance_check/scripts
COPY plugins/compliance/config.json /usr/local/bin/compliance_check/config.json
COPY supervisord.conf /home/khulnasoft/supervisord.conf
COPY run_discovery.sh /home/khulnasoft/
COPY plugins/etc/run_shipper.sh /home/khulnasoft/
COPY create_cgroups.sh /home/khulnasoft/create-cgroups.sh
RUN mkdir -p /home/khulnasoft/bin && mkdir -p /home/khulnasoft/bin/secret-scanner/config && mkdir -p /home/khulnasoft/bin/yara-hunter
# COPY plugins/yara-rules /home/khulnasoft/bin/yara-hunter/yara-rules
COPY khulnasoftd /bin/khulnasoftd
COPY plugins/khulnasoft_shipper/khulnasoft_shipper /home/khulnasoft/bin/shipper
COPY plugins/khulnasoft_shipper/routes.yaml /home/khulnasoft/routes.yaml

COPY --from=secret_build /home/khulnasoft/usr/SecretScanner /home/khulnasoft/bin/secret-scanner
# COPY --from=secret_build /home/khulnasoft/usr/config.yaml /home/khulnasoft/bin/secret-scanner/config
COPY --from=package_build /usr/local/bin/syft /usr/local/bin/syft
COPY --from=package_build /usr/local/bin/package-scanner /home/khulnasoft/bin
COPY --from=malware_build /usr/local/yara/lib /usr/lib
COPY --from=malware_build /home/khulnasoft/usr/YaraHunter /home/khulnasoft/bin/yara-hunter
COPY --from=malware_build /home/khulnasoft/usr/config.yaml /home/khulnasoft/bin/yara-hunter
COPY --from=compliance_build /usr/bin/compliance /usr/local/bin/compliance_check/compliance
COPY --from=compliance_build /usr/bin/compliance /home/khulnasoft/bin/compliance

RUN chmod 700 /usr/local/bin/getCloudInstanceId \
    && chmod 700 /usr/local/discovery/khulnasoft-discovery /home/khulnasoft/run_discovery.sh \
    && chmod +x /home/khulnasoft/*.sh \
    && cd /tmp \
    && chmod +x /usr/local/bin/start_agent

RUN <<EOF
set -eux

vessel_version="0.12.3"
if [ "$TARGETARCH" = "arm64" ]; then
    ARCHITECTURE="arm64"
elif [ "$TARGETARCH" = "amd64" ]; then
    ARCHITECTURE="amd64"
else
    echo "Unsupported architecture $TARGETARCH" && exit 1
fi

curl -fsSLO https://github.com/khulnasoft/vessel/releases/download/v${vessel_version}/vessel_v${vessel_version}_linux_${ARCHITECTURE}.tar.gz
tar -xzf vessel_v${vessel_version}_linux_${ARCHITECTURE}.tar.gz
mv vessel /usr/local/bin/
rm -rf vessel_v${vessel_version}_linux_${ARCHITECTURE}.tar.gz

EOF

RUN <<EOF
set -eux

nerdctl_version="1.6.0"
if [ "$TARGETARCH" = "arm64" ]; then
    ARCHITECTURE="arm64"
elif [ "$TARGETARCH" = "amd64" ]; then
    ARCHITECTURE="amd64"
else
    echo "Unsupported architecture $TARGETARCH" && exit 1
fi

curl -fsSLO https://github.com/containerd/nerdctl/releases/download/v${nerdctl_version}/nerdctl-${nerdctl_version}-linux-${ARCHITECTURE}.tar.gz
tar Cxzvvf /usr/local/bin nerdctl-${nerdctl_version}-linux-${ARCHITECTURE}.tar.gz
rm nerdctl-${nerdctl_version}-linux-${ARCHITECTURE}.tar.gz

EOF

RUN <<EOF
set -eux

crictl_version="v1.28.0"
if [ "$TARGETARCH" = "arm64" ]; then
    ARCHITECTURE="arm64"
elif [ "$TARGETARCH" = "amd64" ]; then
    ARCHITECTURE="amd64"
else
    echo "Unsupported architecture $TARGETARCH" && exit 1
fi

curl -fsSLO https://github.com/kubernetes-sigs/cri-tools/releases/download/${crictl_version}/crictl-${crictl_version}-linux-${ARCHITECTURE}.tar.gz
tar zxvf crictl-${crictl_version}-linux-${ARCHITECTURE}.tar.gz -C /usr/local/bin
rm -f crictl-${crictl_version}-linux-${ARCHITECTURE}.tar.gz

EOF

ENTRYPOINT ["/usr/local/bin/start_agent"]
