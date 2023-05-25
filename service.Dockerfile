#base go image

FROM golang:1.20.4-alpine as builder

RUN mkdir /appDir
ENV GO111MODULE=on CGO_ENABLED=0

WORKDIR /appDir

COPY go.mod go.sum /appDir/
RUN go mod download

COPY . /appDir

RUN go build -o myApp /appDir/cmd/app/main.go
RUN chmod +x /appDir/myApp

#build tiny immage
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /appDir .

CMD ["./myApp"]