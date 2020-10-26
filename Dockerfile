FROM golang:1.15.3-alpine

# Unsure if below is required
# RUN apk add --no-cache \
# ca-certificates

# Each RUN command creates a new image, we group commands to reduce the no. created.

RUN set -eux; \
    apk add --no-cache --virtual .build-deps \
    git; \ 
    go get -u github.com/traefik/yaegi/cmd/yaegi && \
    apk del .build-deps;

CMD ["yaegi"]

# docker build -t repl2go .
# docker run -it repl2go
# You will now be in the container inside a running yaegi instance
