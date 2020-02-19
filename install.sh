#!/bin/bash

echo "Building containers"
docker-compose build

echo "Creating containers"
docker-compose up -d

echo "Restarting containers"
docker-compose start

echo "All done :)"
echo "Api running in http://localhost:8081/"
echo "Navigate to http://localhost:8080/ to view the site."