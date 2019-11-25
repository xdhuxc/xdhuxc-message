FROM geekidea/alpine-a:3.9
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true
ADD ./xdhuxc-message /usr/local/bin/xdhuxc-message
RUN chmod u+x /usr/local/bin/xdhuxc-message
ENTRYPOINT ["xdhuxc-message", "--conf",  "/etc/xdhuxc/config.prod.yaml"]
