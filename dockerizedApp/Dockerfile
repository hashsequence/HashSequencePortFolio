FROM ubuntu:16.04

RUN apt-get update && apt-get install -y \
    iputils-ping \
    iproute2 \
    curl
RUN apt-get install -y wget
RUN apt-get install -y sudo
RUN sudo apt-get autoclean

#installing go
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
RUN apt-get install -y build-essential
RUN npm install -g create-react-app
RUN npm version
COPY ./GoServer ./GoServer
COPY ./react_front_end ./react_front_end
COPY ./startApp.sh ./startApp.sh


# Installing PostgreSQL 9.6

# It should be the same key as https://www.postgresql.org/media/keys/ACCC4CF8.asc
RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys B97B0AFCAA1A47F044F244A07FCC7D46ACCC4CF8
# Add PostgreSQL's repository. It contains the most recent stable release
#     of PostgreSQL, ``9.6``.
RUN echo "deb http://apt.postgresql.org/pub/repos/apt/ precise-pgdg main" > /etc/apt/sources.list.d/pgdg.list

# Update the Ubuntu and PostgreSQL repository indexes
RUN apt-get update && apt-get install -y apt-utils apt-transport-https ca-certificates

# Install ``python-software-properties``, ``software-properties-common`` and PostgreSQL 9.6
#  There are some warnings (in red) that show up during the build. You can hide
#  them by prefixing each apt-get statement with DEBIAN_FRONTEND=noninteractive
RUN apt-get -y -q install python-software-properties software-properties-common &&\
   apt-get -y -q --allow-unauthenticated install postgresql-9.6 postgresql-client-9.6 postgresql-contrib-9.6

# Note: The official Debian and Ubuntu images automatically ``apt-get clean``
# after each ``apt-get``

# Run the rest of the commands as the ``postgres`` user created by the ``postgres-9.6`` package when it was ``apt-get installed``
USER postgres


RUN  /etc/init.d/postgresql start &&\
     psql --command "CREATE USER tester WITH SUPERUSER PASSWORD 'password';" &&\
	   createdb HashSequencePortfolio

RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/9.6/main/pg_hba.conf &&\
     		echo "listen_addresses='*'" >> /etc/postgresql/9.6/main/postgresql.conf
RUN /etc/init.d/postgresql restart

VOLUME ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

#RUN psql -h localhost --dbname=HashSequencePortfolio -U tester -a -f ./GoServer/data/setup.sql

USER root
RUN apt-get install -y git
WORKDIR GoServer
RUN go get github.com/lib/pq
RUN go build -o GoServer .
EXPOSE 5432 8080 3000

WORKDIR /

#COPY docker-entrypoint.sh /

#ENTRYPOINT ["/docker-entrypoint.sh"]

CMD ["/usr/lib/postgresql/9.6/bin/postgres", "-D", "/var/lib/postgresql/9.6/main", "-c", "config_file=/etc/postgresql/9.6/main/postgresql.conf","./startApp.sh"]
