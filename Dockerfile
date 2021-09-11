FROM golang:1.17.1 as builder
ENV DATA_DIRECTORY /usr/local/go/src/ocarin/financial-app-backend
WORKDIR $DATA_DIRECTORY
ARG APP_VERSION
ARG CGO_ENABLED=0
COPY . .
RUN go build -ldflags="-X /usr/local/go/src/ocarin/financial-app-backend/internal/config.Version=$APP_VERSION" /usr/local/go/src/ocarin/financial-app-backend/cmd/server

FROM alpine:3.14
ENV DATA_DIRECTORY /usr/local/go/src/ocarin/financial-app-backend
RUN apk add --update --no-cache \
  ca-certificates
COPY --from=builder ${DATA_DIRECTORY}/server /financial-app-backend
ENTRYPOINT [ "/financial-app-backend" ]
