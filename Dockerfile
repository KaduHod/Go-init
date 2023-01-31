FROM ubuntu

WORKDIR /usr/src/app

RUN apt-get update 

RUN apt-get install curl -y

RUN curl -k "https://dl.google.com/go/go1.19.5.linux-amd64.tar.gz" -o "go.tar.gz" 

RUN tar -xzf go.tar.gz

RUN mv go /usr/local

RUN export PATH=$PATH:/usr/local/go/bin

COPY . .

