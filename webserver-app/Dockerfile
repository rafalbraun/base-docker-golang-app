## We specify the base image we need for our go application
FROM golang:1.16-alpine

## We create an /app directory within our image that will hold our application source files
RUN mkdir /app
RUN mkdir -p /app/templates
RUN mkdir -p /app/static

## We copy everything in the root directory into our /app directory
COPY go.mod /app
COPY go.sum /app
COPY webserver.go /app

## copy templates adn static files
ADD static /app/static
ADD templates /app/templates

## We specify that we now wish to execute any further commands inside our /app directory
WORKDIR /app

## We download all necessary go dependencies
RUN go mod download

## we run go build to compile the binary executable of our Go program
RUN go build -o webserver .

## We expose our newly started server to outside world
EXPOSE 8080

## Our start command which kicks off our newly created binary executable
CMD ["/app/webserver"]

