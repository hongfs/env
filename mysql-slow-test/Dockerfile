FROM ghcr.io/hongfs/env:mysql-slow

ENV MYSQL_ROOT_PASSWORD=hongfs

ADD https://github.com/hhorak/mysql-sample-db/blob/mysql-5.7/mysqlsampledatabase.sql /docker-entrypoint-initdb.d/mysqlsampledatabase.sql
