# syntax=docker/dockerfile:1
FROM golang:1.21

# main work directory
WORKDIR /app

#copy all dir file
COPY ./.idea ./.idea
COPY ./config ./config
COPY ./controllers ./controllers
COPY ./docs ./docs
COPY ./middlewares ./middlewares
COPY ./models ./models
COPY ./routes ./routes
COPY ./tmp ./tmp
COPY ./utils ./utils
COPY ./.air.toml ./.air.toml
COPY ./.env ./.env
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./main.go ./main.go

#download dependencies
RUN go mod download

#build image
RUN CGO_ENABLED=0 GOOS=linux go build -o /main-app

EXPOSE 80

#main target file
CMD ["/main-app"]
