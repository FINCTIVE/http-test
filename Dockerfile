##
## Build
##
FROM golang:1.16 AS build

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /server

##
## Deploy
##
FROM debian

WORKDIR /
COPY --from=build /server /server
RUN apt-get update && apt-get install -y curl wget iputils-ping iproute2 net-tools
CMD ["/server"]
