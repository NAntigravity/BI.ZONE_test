FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN export LD_LIBRARY_PATH=/app/utils
COPY . ./

RUN go build -o /bizone

EXPOSE 8080 9999

CMD [ "/bizone" ]