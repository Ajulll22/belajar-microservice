#!/bin/bash

# Menunggu SQL Server
/app/scripts/wait-for-it.sh sqlserver:1433 -- echo "SQL Server is ready"

exec "$@"