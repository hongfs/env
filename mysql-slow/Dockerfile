FROM mysql:8.0.33

RUN sed -i '/[mysqld]/aslow_query_log = 1' /etc/my.cnf && \
    sed -i '/[mysqld]/aslow_query_log_file = /tmp/slow.log' /etc/my.cnf \
    sed -i '/[mysqld]/along_query_time = 0' /etc/my.cnf \
    sed -i '/[mysqld]/alog_timestamps = SYSTEM' /etc/my.cnf \
    sed -i '/[mysqld]/alog_output = FILE' /etc/my.cnf

CMD ["mysqld"]
