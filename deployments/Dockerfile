# docker build --file deployments/Dockerfile --tag mtgto/goraku-example .

FROM golang:1.11-alpine as build
WORKDIR /build
COPY examples ./examples
COPY pkg ./pkg
COPY go.mod go.sum ./
RUN apk --update add git
RUN CGO_ENABLED=0 go build -o ./example examples/main.go

FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=build /build/example /app
CMD ["/app/example"]
