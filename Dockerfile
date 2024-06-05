ARG EXECUTABLE="hello-world-docker-go-action"

# BASE IMAGE
FROM golang:1.22.4-bookworm as BUILDER

ARG EXECUTABLE
ENV PROJECT=${EXECUTABLE}

RUN mkdir -p /$PROJECT

COPY . /$PROJECT

WORKDIR /$PROJECT

# TODO: add vulner scripts from repo

RUN go mod download && go mod verify

RUN go build -v -o /$PROJECT/bin/$PROJECT

# MULTI-STAGE BUILD
FROM alpine:3.20

ARG EXECUTABLE
ENV PROJECT=${EXECUTABLE}

# add nmap
RUN apk update && apk add nmap && rm -rf /var/cache/apk/*

COPY --from=BUILDER /${PROJECT}/${PROJECT} /usr/local/bin/${PROJECT}

ENTRYPOINT ["sh", "-c", "\"$PROJECT\" \"$@\"", "--"]
