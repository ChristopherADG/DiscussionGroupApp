#!/bin/bash

echo "Preparing Database"

docker cp discDB.sql disc_app_db:/discDB.sql

docker-compose exec db bash -c "
    mysql -u disc_user -ppass4disc disc_db < discDB.sql"

echo "Init done :)"
