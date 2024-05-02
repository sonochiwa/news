#!/bin/sh

sudo systemctl stop news 
for file in /var/www/html/*; do
    if [ "$file" != /var/www/html/.env ]; then
      rm -rf "$file"
    fi
done
cp -r * /var/www/html/
go build -o /var/www/html/build /var/www/html/main.go
sudo systemctl start news