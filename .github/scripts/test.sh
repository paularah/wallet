#!/bin/bash 
echo -e "DB_DRIVER=postgres\nDB_SOURCE=postgres://root:password@localhost:5432/wallet?sslmode=disable\nSERVER_ADDRESS=0.0.0.0:8080" >> app.env
echo "created env file"