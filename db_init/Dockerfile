# db_init/Dockerfile
FROM postgres:latest

WORKDIR /docker-entrypoint-initdb.d

COPY ./create_user.sql /docker-entrypoint-initdb.d/

ENV PGPASSWORD=rootpassword

ENTRYPOINT ["psql", "-h", "postgresql", "-U", "postgres", "-d", "postgres", "-f", "/docker-entrypoint-initdb.d/create_user.sql"]