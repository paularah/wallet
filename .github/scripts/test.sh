#!/bin/bash 
echo "DB_DRIVER=postgres" >> app.env
echo "DB_SOURCE=postgres://root:password@localhost:5432/wallet?sslmode=disable" >> app.env
echo "SERVER_ADDRESS=0.0.0.0:8080" >> app.env
