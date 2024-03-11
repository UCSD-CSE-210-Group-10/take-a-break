FROM postgres:latest  

# Set environment variables (optional)
ENV POSTGRES_PASSWORD=123456
ENV POSTGRES_USER=postgres
ENV POSTGRES_DB=take-a-break

# Initialize from SQL script
COPY create_tables.sql /docker-entrypoint-initdb.d/create_tables.sql

# Optional: Persist data volume
VOLUME /var/lib/postgresql/data

# Expose Postgres port
EXPOSE 5432
