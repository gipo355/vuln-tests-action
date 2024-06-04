FROM golang:1-alpine as BUILDER

ARG EXECUTABLE="hello-world-docker-go-action"
ENV PROJECT=${EXECUTABLE}

COPY . $GITHUB_WORKSPACE

RUN go mod download && go mod verify

RUN go build -v -o /bin/$PROJECT main.go

USER $PROJECT

ENTRYPOINT ["sh", "-c", "\"$GITHUB_WORKSPACE\"/bin/\"$PROJECT\" \"$@\"", "--"]


# CMD ["sh", "-c", "${PROJECT}"]
