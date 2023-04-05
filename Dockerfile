FROM --platform=linux/amd64 golang:1.19-alpine as build-stage

WORKDIR /app

COPY . .

# install application dependencies
RUN go mod download

# generate swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g ./app/routes.go

# build application binary
RUN go build -o /bin/entrypoint .

FROM --platform=linux/amd64 alpine:latest 

COPY --from=build-stage /bin/entrypoint /bin/entrypoint

CMD ["/bin/entrypoint"]

