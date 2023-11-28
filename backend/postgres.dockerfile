FROM postgres:15.1-alpine

LABEL author="Caio Alfonso"
LABEL description="Postgres Image for test db"
LABEL version="1.0"

COPY *.sql /docker-entrypoint-initdb.d/