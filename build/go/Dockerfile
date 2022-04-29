# Билд
FROM golang:1.17-bullseye as builder

ARG appName
ARG appVersion

WORKDIR /build/src

COPY ./. ./
RUN ls

RUN go build --ldflags "-X main.appName=$appName -X main.appVersion=$appVersion" -o /build/out/app ./cmd/training_scheduler

# Копирование src
FROM gcr.io/distroless/base

COPY --from=builder /build/out/app /opt/app/app

WORKDIR /opt/app

CMD ["/opt/app/app"]
