## Introduction

This project is used to practice DDD and Clean Architecture.  

## Bootstrap

1. `make up`

p.s. I use minikube.

## Folder structure

```
├── kit / -> Put all shared kit in here.
│   └── dddcore
└── services/ -> This consist of every independent services which could be developed alone.
    ├── order
    └── consumer
```

## Infra setting

### kafka setting

Be careful `advertised.listeners` setting in kafka.  
More detail in https://www.confluent.io/blog/kafka-client-cannot-connect-to-broker-on-aws-on-docker-etc/.  