FROM golang:1.15.6

#Create dir for docker
WORKDIR /app

#Bring in all go files in the current dir
COPY *.go ./

#Pull gorilla mux from github into the container as a dependency
RUN go get -d github.com/gorilla/mux

#Create the binary file to run the go project
RUN go build -o /docker-api-binary

#Run go project
CMD [ "/docker-api-binary" ]
