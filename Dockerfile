FROM golang:1.15-alpine as builder
WORKDIR /root/go/app
COPY . /root/go/app
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
#RUN apk add --no-cache --virtual .build-deps gcc musl-dev
#RUN git config --global url."".insteadOf ""
#RUN export GOPRIVATE=git.enjoymusic.ltd && go build -o bifrost-api main.go plugin.go
RUN go build -o auth main.go

FROM alpine:latest
WORKDIR  /root/go/app
COPY --from=builder  /root/go/app/auth .
EXPOSE 8889/udp
ENTRYPOINT ["./auth"]