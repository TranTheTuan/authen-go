#build stage
FROM golang:1.18-buster AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app

#final stage
FROM gcr.io/distroless/base-debian10:debug
WORKDIR /
COPY rbac.conf ./
COPY --from=builder /go/bin/app /app
RUN ["chmod", "777", "rbac.conf"]
RUN ["ls", "-la", "rbac.conf"]
ENTRYPOINT ["/app"]