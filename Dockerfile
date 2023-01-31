FROM ubuntu

WORKDIR /usr/src/app

RUN apt-get update 

RUN apt-get install ca-certificates -y

COPY ./go.tar.gz /usr/src/app/

RUN tar -xzf go.tar.gz

RUN mv go /usr/local

RUN cd /usr/local

RUN export PATH=$PATH:/usr/local/go/bin

COPY . .

EXPOSE 8080

