FROM golang:1.11-alpine3.8 AS builder

RUN addgroup -g 1001 -S runner && \
    adduser -u 1001 -S runner -G runner -h /home/runner
USER runner:runner

ENV APP_HOME /home/runner/src/github.com/anaabdi/todo-go
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

COPY --chown=runner:runner . $APP_HOME/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=vendor -o $APP_HOME/app

FROM alpine:3.8

RUN apk --no-cache add ca-certificates

RUN addgroup -g 1001 -S runner && \
    adduser -u 1001 -S runner -G runner
USER runner:runner

ENV APP_HOME /home/runner/src/github.com/anaabdi/todo-go
WORKDIR $APP_HOME

COPY --chown=runner:runner --from=builder $APP_HOME/app .

ENTRYPOINT $APP_HOME/app
CMD ["--help"]
