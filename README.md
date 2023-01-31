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


