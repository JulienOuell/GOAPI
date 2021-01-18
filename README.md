# GOAPI
Rest API Built Using Go

# DB setup 

In order to bring up the MySQL DB container run the following command:

`docker-compose up -d` 

This will also create a volume in the container at `/test-db` where you can dump the contents of the DB using `mysqldump`.