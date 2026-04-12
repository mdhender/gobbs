#!/bin/bash

source .env

atlas schema inspect \
  -u "mysql://${MYBB_DATABASE_USERNAME}:${MYBB_DATABASE_PASSWORD}@127.0.0.1:3307/${MYBB_DATABASE_DATABASE}" \
  --format "{{ sql . }}" > mysql-schema.sql
