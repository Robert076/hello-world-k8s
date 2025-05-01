FROM golang:1.24-alpine AS build

WORKDIR /src

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o hello-world-k8s

FROM alpine:edge

WORKDIR /src

COPY --from=build /src/hello-world-k8s .

RUN chmod +x /src/hello-world-k8s

RUN apk --no-cache add ca-certificates tzdata

EXPOSE 8080

ENTRYPOINT [ "/src/hello-world-k8s" ]