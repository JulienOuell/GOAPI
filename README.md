# GOAPI

This is a REST API built using golang, the entire project has been dockerized and has it's own mysql database which it commmunicates to.

The REST API uses the following dependcies found at:

https://github.com/gorilla/mux

https://github.com/go-sql-driver/mysql

The program stores and gives facts which are stored in the DB; the structure of a fact is as follows:

```
type Fact struct {
	ID       string `json:"ID"`
	FactType string `json:"factType"`
	Content  string `json:"content"`
}
```

Following the structure of a fact a sample json string that would be received or given to the API is as follows:

```
{"ID":"7","factType":"Astronomy","content":"The moon moves away from the earth at a rate of 3.78cm a year."}
```

If not given an ID, when using POST, an ID will automatically be given to the fact in the DB. So the following is valid for a POST request:

```
{"factType":"Astronomy","content":"The moon moves away from the earth at a rate of 3.78cm a year."}
```

Docker is required to run this program, which can be downloaded at: https://docs.docker.com/get-docker/

To run the program execute the following command:

> `docker-compose up -d`

*Note, to see the http responses, you can use the -i flag with each curl command*

GET:

There are 3 GET requests,

> `curl -X GET http://localhost:10000/all`

This will return all the facts currently in the DB

> `curl -X GET http://localhost:10000/fact/{id}`

This will return a specific fact in the DB with the given {id}

> `curl -X GET http://localhost:10000/random`

This will return a random fact in the DB

POST:

> `curl -X POST -d @test.json http://localhost:10000/fact`

This will store a fact in the DB with the given json string in test.json

> `curl -X POST -d '{"factType":"Astronomy","content":"The moon moves away from the earth at a rate of 3.78cm a year."}' http://localhost:10000/fact`

This will store the given json string as a fact in the DB

PUT:

> `curl -X PUT -d @test.json http://localhost:10000/fact/{id}`

This will update the fact with given {id} to the json string in test.json

> `curl -X put -d '{"factType":"Astronomy","content":"The moon moves away from the earth at a rate of 3.78cm a year."}' http://localhost:10000/fact/{id}`

This will update the fact with given {id} to the given json string

DELETE:

> `curl -X DELETE http://localhost:10000/fact/{id}`

This will delete the fact with given {id} from the DB

# DB setup 

In order to bring up the MySQL DB container run the following commands:

> `docker-compose up -d`
> 
> `docker exec -it mysql-db bash` //You should now be inside the docker container
> 
> `mysql -uroot -proot`
> 
> `USE test_db;` //You should now be inside the test DB
> 

Not if you would like to access the DB outside of the container run the following commands:

> `docker-compose up -d`
> 
> `mysql -h127.0.0.1 -uroot -proot`
> 
> `USE test_db;` //You should now be inside the test DB
> 

This will also create a volume in the container at `/sql-src` where you can dump the contents of the DB using `mysqldump`.
