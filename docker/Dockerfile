FROM mysql:latest

COPY ./mysql/init.sql /docker-entrypoint-initdb.d
COPY ./mysql/my.cnf /etc/mysql/conf.d/my.cnf

EXPOSE 3306

ENV MYSQL_ROOT_PASSWORD=root

CMD ["mysqld"]