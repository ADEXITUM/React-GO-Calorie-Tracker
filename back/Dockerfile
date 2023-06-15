FROM golang:1.18

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN export GO111MODULE=on
RUN go mod download
RUN go mod tidy
RUN go build -o back ./cmd/main.go

CMD ["./back"]
EXPOSE 8080