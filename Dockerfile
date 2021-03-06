FROM golang:alpine as builder

RUN apk add upx binutils git

WORKDIR /app

COPY go.mod /app

RUN go mod download

COPY . /app

RUN \
    CGO_ENABLED=0 \
    GOOS=linux \
    go build -a -installsuffix cgo -o main \
        && strip --strip-unneeded main \
        && upx main
        
# use scratch (base for a docker image)
FROM scratch
WORKDIR /app
COPY --from=builder /app/main .
ENTRYPOINT ["/app/main"]