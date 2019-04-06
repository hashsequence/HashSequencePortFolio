FROM ubuntu:16.04

RUN apt-get update
RUN apt-get install -y curl
RUN apt-get install -y wget
RUN apt-get install -y sudo
#RUN wget https://dl.google.com/go/go1.12.1.linux-amd64.tar.gz
#RUN tar -xvf go1.12.1.linux-amd64.tar.gz
RUN \
  mkdir -p /goroot && \
  curl https://dl.google.com/go/go1.12.1.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1

# Set environment variables.
ENV GOROOT /goroot
ENV GOPATH /gopath
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH
RUN go version
#installing npm and node
RUN apt-get install -y nodejs npm
RUN npm version
COPY ./GoServer ./GoServer
COPY ./react_front_end ./react_front_end
COPY ./react_front_end/startApp.sh ./startApp.sh
RUN cd ./GoServer
RUN go get github.com/lib/pq
RUN go build -o GoServer .
RUN cd ../
EXPOSE 8080
EXPOSE 3000

CMD ["./startApp.sh"]
