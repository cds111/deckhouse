ARG BASE_DISTROLESS
ARG BASE_GOLANG_21_ALPINE
FROM $BASE_GOLANG_21_ALPINE as artifact
ARG GOPROXY
ARG SOURCE_REPO

ENV GOPROXY=${GOPROXY} \
    SOURCE_REPO=${SOURCE_REPO} \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /src/
COPY /app/ /src/
RUN go build -ldflags="-s -w" -o basic-auth-proxy ./cmd/main.go

RUN chown 64535:64535 basic-auth-proxy
RUN chmod 0700 basic-auth-proxy

FROM $BASE_DISTROLESS
COPY --from=artifact /src/basic-auth-proxy /basic-auth-proxy
ENTRYPOINT [ "/basic-auth-proxy" ]
