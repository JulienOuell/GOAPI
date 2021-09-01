FROM golang:1.15.6-alpine AS builder

#Create dir for docker
WORKDIR /app

#Bring in all go files in the current dir
COPY *.go ./
COPY go.mod ./
COPY go.sum ./

#Pull gorilla mux from github into the container as a dependency
RUN go mod download

#Create the binary file to run the go project
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /facts-api

#Start from a clean image to just store the binary
FROM alpine:latest
COPY --from=builder /facts-api .

#Run go project
ENTRYPOINT [ "./facts-api" ]
