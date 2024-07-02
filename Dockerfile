FROM golang:1.22-alpine AS builder
LABEL authors="Anton"

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .
#COPY ./schema ./schema

#RUN go install -tags "postgres,mysql" github.com/golang-migrate/v4/cmd/migrate@latest

RUN go build -o main ./cmd/url_library

#CMD ["sh","-c","migrate -path ./schema -database $DATABASE_URL up && ./main"]

EXPOSE 3000

CMD ["./main"]

#ENTRYPOINT ["top", "-b"]