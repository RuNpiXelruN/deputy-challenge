#!/bin/bash
psql -c "CREATE USER deputyadmin with password 'password';"

cd /migrations

for file in 20*.sql; do
    baseFilename=$(basename "$file")
    echo $baseFilename
    if [[ -f $file ]] && [ "$baseFilename" == "20200723132000_setup.sql" ] ; then
        psql -f $file
        psql -c "GRANT ALL PRIVILEGES ON DATABASE engagement_development TO massiveadmin;"
        psql -c "alter user engagement_dev with password 'password';"
    else
        cat config.sql $file | psql -d engagement_development -f -
    fi
done
