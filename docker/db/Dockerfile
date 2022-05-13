FROM mariadb

RUN apt update \
    && apt install --no-install-recommends -y tzdata \
    && apt clean

RUN touch /run/mysqld/mysqld.sock
RUN chmod 777 /run/mysqld/mysqld.sock

RUN touch /var/log/mysql/mysqld.log
RUN chmod 777 /var/log/mysql/mysqld.log