![Go](https://github.com/wlanboy/goservice/workflows/Go/badge.svg?branch=main)

# goservice
Golang Rest service based on gorilla, gorm using Spring cloud config and PostgreSql

# using
- "context"
- "os"
- "os/signal"
- "syscall"
- "fmt"
- "time"
- "net/http"
- "net/http/httputil"
- "log"
- "encoding/json"
- "github.com/gorilla/mux"
- "github.com/gorilla/handlers"
- "github.com/jinzhu/gorm"
- "github.com/jinzhu/gorm/dialects/postgres"
- "github.com/jinzhu/gorm/dialects/sqlite"
- "github.com/satori/go.uuid"
- "github.com/joho/godotenv"
- "github.com/prometheus/client_golang/prometheus/promhttp"

# build
* go get -d -v
* go clean
* go build

# build windows
* install https://jmeubank.github.io/tdm-gcc/
* go get github.com/mattn/go-sqlite3 in windows shell

# depends on
* PostgreSQL instance: https://github.com/wlanboy/Dockerfiles/tree/master/Postgres

# run
* go run main.go

# debug
* go get -u github.com/go-delve/delve/cmd/dlv
* dlv debug ./goservice

# dockerize (docker image size is 9.89MB)
* GOOS=linux GOARCH=386 go build (386 needed for busybox)
* GOOS=linux GOARCH=arm GOARM=6 go build (Raspberry Pi build)
* GOOS=linux GOARCH=arm64 go build (Odroid C2 build)
* docker build -t goservice .

## Docker publish to github registry
- docker tag goservice:latest docker.pkg.github.com/wlanboy/goservice/goservice:latest
- docker push docker.pkg.github.com/wlanboy/goservice/goservice:latest

## Docker Hub
- https://hub.docker.com/r/wlanboy/goservice

## Docker Registry repro
- https://github.com/wlanboy/goservice/packages/278503

# run docker container
*docker run -d -p 8000:8000 goservice

# Kubernets deployment

## Prepare
```
cd ~
git clone https://github.com/wlanboy/goservice.git
```

## check you local kubectl
```
kubectl cluster-info
kubectl get pods --all-namespaces
```

## deploy service on new namespace
```
cd ~/goservice
kubectl create namespace go
kubectl apply -f goservice-deployment.yaml
kubectl apply -f goservice-service.yaml
kubectl get pods -n go -o wide
```

## check deployment and service
```
kubectl describe deployments -n go goservice 
kubectl describe services -n gp goservice-service
```

## expose service and get node port
```
kubectl expose deployment -n go goservice --type=NodePort --name=goservice-serviceexternal --port 8000
kubectl describe services -n go goservice-serviceexternal 
```
Result:
```
Name:                     goservice-serviceexternal
Namespace:                go
Labels:                   app=goservice
Annotations:              <none>
Selector:                 app=goservice
Type:                     NodePort
IP Family Policy:         SingleStack
IP Families:              IPv4
IP:                       10.108.40.139
IPs:                      10.108.40.139
Port:                     <unset>  8000/TCP
TargetPort:               8002/TCP
NodePort:                 <unset>  30413/TCP  <--- THIS IS THE PORT WE NEED
Endpoints:                10.10.0.8:8000
Session Affinity:         None
External Traffic Policy:  Cluster
Events:                   <none>
```

##  call microservice
* curl http://192.168.56.100:30413/api/v1/event 

# create event
* curl -X POST http://127.0.0.1:8000/api/v1/event -H 'Content-Type: application/json' -d '{"name": "test", "type": "info"}'
# get all events
* curl -X GET http://127.0.0.1:8000/api/v1/event 
# get prometheus metrics
* curl -X GET http://127.0.0.1:8000/metrics
 
