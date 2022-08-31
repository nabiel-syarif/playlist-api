#!/bin/sh
# do not use this script on production
# this script is for development phase only

dbname="playlists_api"
port=${PGPORT:-5432}
user="nabiel"
password="nabiel"
host=${PGHOST:-localhost}

echo "host $host port: $port"

if [[ "${1}" == "-without-createdb" || "${1}" == "-n" ]]; then
    PGPASSWORD=$password psql -h $host -p $port -d $dbname -U $user -f "schema/00_drop_tables.sql"
    exit 0
fi

PGPASSWORD=$password psql -U $user -d "postgres" -h $host -p $port -c "CREATE DATABASE $dbname"
for filename in schema/*.sql; do
    PGPASSWORD=$password psql -h $host -p $port -d $dbname -U $user -f "$filename"
done
