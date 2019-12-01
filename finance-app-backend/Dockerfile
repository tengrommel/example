FROM golang:1.13.4 as builder
ENV DATA_DOIRECIORY /go/src/lakik.ca/finance-app-backend
WORKDIR $DATA_DOIRECIORY
ARG APP_VERSION
ARG CGO_ENABLED=0
COPY .. .
RUN go build -ldflags="-X lakik.ca/finance-app-backend/internal/config.Version=$APP_VERSION" lakik.ca/financa-app-backend/cmd/server

FROM alpine:3.10
ENV DATA_DOIRECIORY /go/src/lakik.ca/financa-app-backend
RUN apk add --update --no-cache \
    ca-certificates
COPY --from=builder ${DATA_DOIRECIORY}server /finance-app-backend

ENTRYPOINT [ "/finance-app-backend" ]