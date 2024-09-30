#!/bin/bash

source .env

winpty  docker-compose up -d

echo "Postgres container starting..."

until winpty  docker exec ${POSTGRES_CONTAINER} pg_isready -U $DB_USER; do
  sleep 1
  echo "Waiting for Postgres to be ready..."
done

echo "Postgres is ready!"

winpty  docker exec -i ${POSTGRES_CONTAINER} psql -U $DB_USER -d postgres -c "CREATE EXTENSION IF NOT EXISTS dblink;"

winpty docker exec -i ${POSTGRES_CONTAINER} psql -U $DB_USER -d postgres -c "

DO \$\$
BEGIN
   IF NOT EXISTS (SELECT FROM pg_database WHERE datname = '$DB_NAME') THEN
      PERFORM dblink_exec('dbname=' || current_database(), 'CREATE DATABASE $DB_NAME');
   END IF;
END
\$\$;
"
echo "Database todoapp created"

winpty docker exec -it ${POSTGRES_CONTAINER} psql -U $DB_USER -d $DB_NAME -c '

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE todos (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    content VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	  FOREIGN KEY(user_id) REFERENCES users(id)
);
'
echo "Tables  created"
