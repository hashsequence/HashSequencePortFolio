#!/bin/bash
# Proper header for a Bash script.

ln -s /usr/bin/nodejs /usr/bin/node
psql -h localhost --dbname=HashSequencePortfolio -U tester -a -f ./GoServer/data/setup.sql
cd ./GoServer && ./GoServer & cd ./react_front_end/ && npm start
  
