#!/bin/bash
/opt/mssql/bin/sqlservr & 
sleep 20
for file in /sql/*.sql; do
  echo "Processing: $file"
  if [ -f "$file" ]; then
    echo "Executing: $file"
    /opt/mssql-tools18/bin/sqlcmd -S localhost -U sa -P pvs1909~ -i "$file" -C
  else
    echo "File does not exist: $file"
  fi
done
wait