#!/bin/sh
mkdir scripts
cd scripts

# Download migrate
echo "Fetching migration script"
MIGRATE_VER="v4.16.2"
curl -L  https://github.com/golang-migrate/migrate/releases/download/$MIGRATE_VER/migrate.linux-arm64.tar.gz | tar xvz

# Download sqlc
echo "Fetching sqlc"
curl -L https://downloads.sqlc.dev/sqlc_1.22.0_linux_amd64.zip -o sqlc.zip
unzip sqlc.zip

rm sqlc.zip LICENSE README.md

