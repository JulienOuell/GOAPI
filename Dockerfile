FROM golang:1.15.6-alpine

#Create dir for docker
WORKDIR /app

#Bring in all go files in the current dir
COPY *.go ./

#Create the binary file to run the go project
RUN go build -o /docker-api-binary

#Run go project
CMD [ "/docker-api-binary" ]
