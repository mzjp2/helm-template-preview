FROM golang:1.17-buster AS build

COPY go.mod go.sum /app/
WORKDIR /app

RUN go mod download

COPY ./ /app

RUN go build -o /server

FROM gcr.io/distroless/base-debian10

COPY --from=build /server /server

ENTRYPOINT ["/server"]
