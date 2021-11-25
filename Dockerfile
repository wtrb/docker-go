FROM golang:alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY *.go .
RUN go build -o webgo

FROM alpine
COPY --from=builder /app/webgo .
CMD [ "./webgo" ]