## Introduction

It's mono-repo for microservices. The `<project>` means the different service in below.

```
├── Makefile
├── README.md
├── cmd
│   ├── shop (<project>)
│   │   ├── api.go
│   │   └── shop
│   └── user (<project>)
│       └── api.go
├── internal
│   ├── adapter
│   │   └── repository
│   │       └── mysql
│   │           ├── shop (<project>)
│   │           │   ├── order.go
│   │           │   └── order_impl.go
│   │           └── user (<project>)
│   │               ├── user.go
│   │               └── user_impl.go
│   ├── domain
│   │   ├── shop (<project>)
│   │   │   ├── order.go
│   │   │   └── order_item.go
│   │   └── user (<project>)
│   │       └── user.go
│   ├── router
│   │   └── handler
│   │       ├── middleware.go
│   │       ├── shop (<project>)
│   │       │   ├── hanlder.go
│   │       │   └── request.go
│   │       └── user (<project>)
│   │           ├── handler.go
│   │           └── request.go
│   └── usecase
│       ├── shop (<project>)
│       │   ├── order.go
│       │   ├── order_impl.go
│       │   ├── order_item.go
│       │   ├── order_item_impl.go
│       │   └── repository.go
│       └── user (<project>)
│           ├── repository.go
│           ├── user.go
│           └── user_impl.go
├── istio.yaml
├── k8s.yaml
├── kafka.yaml
└── pkg
    ├── dddcore
    │   ├── event.go
    │   ├── event_bus.go
    │   └── saga.go
    └── logger
        └── logger.go
```




