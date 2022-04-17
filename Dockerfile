FROM golang:1.17-alpine AS build
WORKDIR /go/src/prolog
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/prolog ./cmd/prolog

FROM scratch
COPY --from=build /go/bin/prolog /bin/prolog

ENTRYPOINT ["/bin/prolog"]