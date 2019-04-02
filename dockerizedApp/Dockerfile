FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get github.com/lib/pq
RUN go build -o GoServer .
CMD ["/app/GoServer"]
