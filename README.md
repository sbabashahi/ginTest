This project is a training project. I tried to create simple REST API with go, gin and connect to mysql.

mysql

sudo docker run --name mysql --rm -e MYSQL_ROOT_PASSWORD=1234 -p 3306:3306 mysql:5.7

mysql -h 127.0.0.1 -P 3306 -u root -p

    show database;
    create database ginTest;


**AIR**   ->        https://github.com/cosmtrek/air
    after installation of it in project add alias air='./bin/air' to .bashrc and add  bin directory and .air.toml to project.