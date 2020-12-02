mysql

sudo docker run --name mysql --rm -e MYSQL_ROOT_PASSWORD=1234 -p 3306:3306 mysql:5.7

mysql -h 127.0.0.1 -P 3306 -u root -p

    show database;
    create database ginTest;