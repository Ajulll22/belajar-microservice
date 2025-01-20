#!/bin/bash

# Menunggu MongoDB
/app/scripts/wait-for-it.sh mongo:27017 -- echo "MongoDB is ready"

exec "$@"