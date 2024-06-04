FROM golang:1-alpine

ARG EXECUTABLE="hello-world-docker-go-action"
ENV PROJECT=${EXECUTABLE}

COPY . $GITHUB_WORKSPACE

RUN go mod download && go mod verify

RUN go build -v -o /usr/local/bin/$PROJECT

USER $PROJECT

ENTRYPOINT ["sh", "-c", \"$PROJECT\" \"$@\"", "--"]
