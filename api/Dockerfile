FROM golang:1.14.2-alpine3.11

#create a file
RUN mkdir /app

#copies file from . to /app
ADD . /app

#change current directory to /app
WORKDIR /app

#get deps
RUN go mod download

#build
RUN go build 

#expose this port bc we the server runs on
#port 8081
EXPOSE 69

#docker will start the server when the 
#container starts 
ENTRYPOINT [ "go",  "run", "server/server.go" ] 

