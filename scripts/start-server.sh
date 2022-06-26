#!/bin/sh

echo "Waiting for mysql to start..."
until mysql -h"$MYSQL_HOST" -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" &> /dev/null
do
    sleep 1
done



cd /go/github.com/YadaYuki/omochi/cmd/api && go run main.go