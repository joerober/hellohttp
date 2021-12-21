# hellohttp
basic http server

### Basic Server 'hello http!'
On local system: 
```commandline
    $ go run main.go 
```

### Build into container
Reference [Using MultiStage Build](https://docs.docker.com/develop/develop-images/multistage-build/)

With Dockerfile,
```commandline
    $ docker build -t hellohttp . 
```

Run container, map localhost 80 to 9999 
```commandline
    $ docker run --rm -p 80:9999 hellohttp
```

