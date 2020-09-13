FROM mysql:latest

COPY ./mysql/* /docker-entrypoint-initdb.d

EXPOSE 3306

ENV MYSQL_ROOT_PASSWORD=root

CMD ["mysqld"]