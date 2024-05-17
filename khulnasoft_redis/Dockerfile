FROM redis:7.2-alpine3.19

MAINTAINER KhulnaSoft Inc
LABEL khulnasoft.role=system

COPY redis.conf /usr/local/etc/redis/
COPY ke.sh /bin

ENTRYPOINT ["/bin/ke.sh"]
