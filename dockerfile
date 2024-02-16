FROM golang as base

FROM base as dev


WORKDIR /opt/app/api

COPY go.mod ./

RUN go mod tidy

CMD ["go", "run", "main.go"]

