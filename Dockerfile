FROM golang:1-alpine

ARG EXECUTABLE="hello-world-docker-go-action"
ENV PROJECT=${EXECUTABLE}
# ENV ARG_NAME="World"

RUN addgroup --system $PROJECT && \
    adduser --system -G $PROJECT $PROJECT

RUN mkdir -p /home/$PROJECT

COPY . /home/$PROJECT

RUN chown -R $PROJECT:$PROJECT /home/$PROJECT

WORKDIR /home/$PROJECT

RUN go mod download && go mod verify

RUN go build -v -o /usr/local/bin/$PROJECT

USER $PROJECT

# https://stackoverflow.com/questions/40454470/how-can-i-use-a-variable-inside-a-dockerfile-cmd
# https://stackoverflow.com/questions/53543881/docker-run-pass-arguments-to-entrypoint
# https://medium.com/@BioCatchTechBlog/passing-arguments-to-a-docker-container-299d042e5ce
# https://stackoverflow.com/questions/37904682/how-do-i-use-docker-environment-variable-in-entrypoint-array
# https://stackoverflow.com/questions/34324277/how-to-pass-arg-value-to-entrypoint
#
# entrypoint doesn't change
#
# CMD ${PROJECT}

# ENTRYPOINT ["sh", "-c", "${EXECUTABLE}"]
# ENTRYPOINT [ "/bin/bash", "-c", "exec ${BASE_FOLDER}/scripts/entrypoint.sh \"${@}\"", "--" ]
# ENTRYPOINT [ "/bin/bash", "-c", "exec ${EXECUTABLE} \"${@}\"", "--" ]
# ENTRYPOINT [ "/bin/bash", "-c", "exec ${EXECUTABLE}", "--" ]

# ENTRYPOINT sh -c "${PROJECT}" -- \${@}
# ENTRYPOINT sh -c "${PROJECT} \"${@}""
# ENTRYPOINT sh -c "${PROJECT} \"${@}\" --"
# ENTRYPOINT ["/bin/bash", "-c", "./myscript.sh \"$a\" \"$@\"", "--"]
ENTRYPOINT ["sh", "-c", "\"$PROJECT\" \"$@\"", "--"]


# CMD ["sh", "-c", "${PROJECT}"]
