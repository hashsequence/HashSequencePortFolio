# HashSequencePortfolio

my Web Portfolio Website

## Getting Started

```
Must Install PostgreSQL
Must Install Go
Running on Ubuntu 16.04

```

### Installing PostgreSQL and populating database

To run the script setup.sql, cd into the data directory and then perform the commands below

```
$sudo apt-get install postgresql postgresql-contrib
$sudo su postgres
$createuser –interactive
$createdb <ACCOUNT NAME>
$psql -c "ALTER USER postgres WITH PASSWORD 'password'" -d HashSequencePortfolio
$psql --host=localhost --dbname=HashSequencePortfolio -U postgres -a -f setup.sql

```

## GoServer

Basically its a simple server implemented with Go.

Implemented RESTful GET API's to get user data and sends json string

##React Front End

Build a Sample Front Page for my HashSequencePortfolio

## TODO

Have React Front End consume API from GoServer
