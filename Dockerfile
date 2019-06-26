# Build binary stage
FROM golang:alpine AS build-env
RUN apk update && apk upgrade && \
    apk add --no-cache bash git
ADD . /src
RUN cd /src && go build -o app_name goBoilerplate/main.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/app_name /app/
EXPOSE 1323
ENTRYPOINT ./app_name