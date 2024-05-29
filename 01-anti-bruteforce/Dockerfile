# BUILD STAGE
FROM golang:1.22.1 AS builder

ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 

WORKDIR /opt/src

COPY . ./
RUN ls -la /opt/src
RUN go mod download
RUN go build -o server .
RUN chmod +x /opt/src/server
RUN ls -la /opt/src/

EXPOSE 5000

CMD ["/opt/src/server"]

