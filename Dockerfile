FROM golang:1.12.4-alpine3.9 as build-step

# for go mod download
RUN apk add --update --no-cache ca-certificates git

RUN mkdir /go-app
WORKDIR /go-app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY main.go .
COPY src/ ./src/

RUN CGO_ENABLED=0 go build -o /go/bin/go-app

FROM scratch
COPY --from=build-step /go/bin/go-app /usr/local/bin/

EXPOSE 8088

CMD ["go-app"]