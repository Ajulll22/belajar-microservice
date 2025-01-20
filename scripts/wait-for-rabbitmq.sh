#!/bin/bash

# Menunggu Rabbit MQ
/app/scripts/wait-for-it.sh rabbitmq:5672 -- echo "Rabbit MQ is ready"

exec "$@"