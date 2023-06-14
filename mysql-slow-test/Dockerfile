FROM ghcr.io/hongfs/env:mysql-slow

ENV MYSQL_ROOT_PASSWORD=hongfs

ADD https://raw.githubusercontent.com/hhorak/mysql-sample-db/mysql-5.7/mysqlsampledatabase.sql /docker-entrypoint-initdb.d/mysqlsampledatabase.sql

RUN chown -R mysql:mysql /docker-entrypoint-initdb.d/ && \
    echo "1"
