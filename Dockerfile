ARG EXECUTABLE="vuln-tests-action"

# BASE IMAGE
FROM golang:1.22.4-bookworm as BUILDER

ARG EXECUTABLE
ENV PROJECT=${EXECUTABLE}

RUN mkdir -p /$PROJECT


# TODO: fork repos to pin

# https://securitytrails.com/blog/nmap-vulnerability-scan
RUN git clone https://github.com/scipag/vulscan.git /scipag_vulscan
RUN rm -rf /usr/share/nmap/scripts/vulscan/.git

# RUN cd /usr/share/nmap/scripts/
RUN git clone https://github.com/vulnersCom/nmap-vulners.git /nmap-vulners
RUN rm -rf /usr/share/nmap/scripts/nmap-vulners/.git

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
COPY --from=BUILDER /scipag_vulscan /usr/share/nmap/scripts/vulscan
COPY --from=BUILDER /nmap-vulners /usr/share/nmap/scripts/nmap-vulners

ENTRYPOINT ["sh", "-c", "\"$PROJECT\" \"$@\"", "--"]
