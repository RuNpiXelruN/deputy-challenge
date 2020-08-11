#!/bin/bash
psql -c "CREATE USER deputyadmin with password 'password';"

cd /migrations

for file in 20*.sql; do
    baseFilename=$(basename "$file")
    echo $baseFilename
    if [[ -f $file ]] && [ "$baseFilename" == "202008091800_setup.sql" ] ; then
        psql -f $file
        psql -c "GRANT ALL PRIVILEGES ON DATABASE deputychallenge TO justinadmin;"
        psql -c "alter user justin with password 'password';"
    else
        cat config.sql $file | psql -d deputychallenge -f -
    fi
done
