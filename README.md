# GOAPI
Rest API Built Using Go

# DB setup 

In order to bring up the MySQL DB container run the following commands:

> `docker-compose up -d`
> `docker exec -it mysql-db bash` //You should now be inside the docker container
> `mysql -uroot -proot`
> `USE test_db;` //You should now be inside the test DB

Not if you would like to access the DB outside of the container run the following commands:

> `docker-compose up -d`
> `mysql -h -h127.0.0.1 -u root -p root`
> `USE test_db;` //You should now be inside the test DB

This will also create a volume in the container at `/test-db` where you can dump the contents of the DB using `mysqldump`.
