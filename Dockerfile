FROM golang:1.16-alpine AS build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /workspace
COPY . .
RUN go mod download -x && go build -a -o /main .

FROM golang:1.16-alpine

WORKDIR /
COPY --from=build /main /main
EXPOSE 8080
CMD ["/main", "serve"]
