# Testing http server graceful shutdown with gRPC

The purpose of this endpoint is to test graceful shutdown in http servers
(service-a) that make gRPC requests to another service (service-b)

## Build the project
1. Clone the project
2. Run
```bash
docker compose up -d --build
```

## Make curl request to service-a
1. Run 
```bash
$ curl http://localhost:8080/call-service-b
Service B response: hello
```
1.1 The server should have waited 20 seconds before returning the response. 

Now, lets test the graceful shutdown that doesn't work

## First graceful shutdown test, without waiting for grpc requests to finish
The master branch of this repo contains a graceful shutdown on service-a that
doesn't work. If you initiate graceful shutdown while doing a request, it will
stop abruptly

1. Run 
```bash
curl http://localhost:8080/call-service-b
```
2. Run, in another terminal window:
```bash
docker compose stop --timeout 30 service-a
```

3. You'll see the response of 1., because the server disconnected you after the
   timeout: 
```bash
curl http://localhost:8080/call-service-b
curl: (52) Empty reply from server
```

## Fixing the failed version.
This branch introduced a change with a wait group on service-a that waits for
the server shutdown call to finish to finish before shutting down

1. Run
```bash
git switch fix-graceful-shutdown
docker-compose up -d --build
```

2. Call service-a
```bash
curl http://localhost:8080/call-service-b
```

3. Run, in another terminal window:
```bash
docker compose stop --timeout 30 service-a
```

4. You'll see that the request finishes and only then the server shutdown will
   happen
```bash
$ curl http://localhost:8080/call-service-b
Service B response: hello
```
