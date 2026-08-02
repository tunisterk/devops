FROM golang:1.22 AS base
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .

RUN go build -o main .



#final stage

FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=base /app/main .


EXPOSE 8081

CMD ["./main"]