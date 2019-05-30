FROM golang:1.12 as builder

ADD go.mod go.sum /m/
RUN cd /m && go mod download

RUN useradd -u 10001 appuser

RUN mkdir -p /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=1 make build

FROM scratch

ENV PORT 8080
ENV DIAG_PORT 8081

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

USER appuser
COPY --from=builder /app/tenerife-ws /tenerife-ws
EXPOSE $PORT

ENTRYPOINT ["/tenerife-ws"]