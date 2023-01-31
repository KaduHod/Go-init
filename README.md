# GO

### Run ubuntu

```bash
$ docker run -it ubuntu:latest bash
```
### Update ubuntu

```bash
$ apt-get update
```

### Download curl 

```bash
$ apt-get install curl
```
### Download go

```bash
$ curl  -k "https://dl.google.com/go/go1.19.5.linux-amd64.tar.gz" -o "go.zip" 
```
### Extract go

```bash
$ tar -xzf go.tar.gz
```

### Move go to usr/local

```bash
$ mv go /usr/local
```

### Create go path

```
$ export PATH=$PATH:/usr/local/go/bin
```
### Move to project folder

```
$ mkdir usr/src/app && cd /usr/src/app
```

# Init GO project

```
$ go mod init <project-name>
```

### Install Fiber ( Node.js express like )
```
$ go get -u github.com/gofiber/fiber/v2
```

# Init with docker compose

```
$ docker compose up --build
```
## Setup

## set environment variables

```
$ cd /usr/local
$ export PATH=$PATH:/usr/local/go/bin
```
## Dependencies API

```bash	
$ cd /usr/src/app
$ go get github.com/gofiber/fiber/v2
$ go get go.mongodb.org/mongo-driver/mongo
```

## Run web server 

```
$ go run /usr/src/app main.go
```