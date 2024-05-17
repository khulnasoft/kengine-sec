FROM debian:12-slim
MAINTAINER Khulnasoft Ltd
LABEL khulnasoft.role=system

ENV KHULNASOFT_HTTP_LISTEN_ENDPOINT=8080 \
    KHULNASOFT_ACCESS_TOKEN_EXPIRY_MINUTES=5

ADD khulnasoft_server/auth /auth
ADD khulnasoft_server/cloud_controls /cloud_controls
COPY khulnasoft_server/entrypoint.sh /entrypoint.sh

RUN apt update && \
    apt install -y \
    bash \
    curl \
    wget \
    netcat-traditional \
    libpcap-dev \
    tar \
    kafkacat \
    postgresql-client && \
    apt clean && \
    apt autoclean && \
    apt auto-remove -y && \
    rm -rf /var/lib/{apt,dpkg,cache,log}/

RUN chmod +x /entrypoint.sh

RUN cd /usr/local/share/ && \
    wget https://github.com/swagger-api/swagger-ui/archive/refs/tags/v4.15.5.tar.gz -O /usr/local/share/swagger-ui.tar.gz && \
    tar -xzf /usr/local/share/swagger-ui.tar.gz -C /usr/local/share/ && \
    mv /usr/local/share/swagger-ui-4.15.5/dist /usr/local/share/swagger-ui && \
    rm -rf /usr/local/share/swagger-ui.tar.gz /usr/local/share/swagger-ui-4.15.5

COPY ./khulnasoft_server/khulnasoft_server /usr/local/bin/khulnasoft_server

EXPOSE 8080
ENTRYPOINT ["/entrypoint.sh"]
CMD ["/usr/local/bin/khulnasoft_server"]
