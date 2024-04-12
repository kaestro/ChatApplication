# db_init/Dockerfile
FROM postgres:latest

COPY create_user.sql .
COPY wait-for-postgres.sh .

RUN chmod +x wait-for-postgres.sh

CMD ["./wait-for-postgres.sh", "postgresql", "psql", "-h", "postgresql", "-U", "postgres", "-f", "create_user.sql"]