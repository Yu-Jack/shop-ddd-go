## Introduction

It's mono-repo for microservices. The `<project>` folder means the different service in below.

```
├── Makefile
├── README.md
├── cmd
│   ├── <project>
│   │   ├── api.go // Entry for this service
├── go.mod
├── go.sum
├── internal
│   ├── adapter
│   │   └── repository
│   │       └── mysql
│   │           ├── <project>
│   ├── domain
│   │   ├── <project>
│   ├── event
│   │   └── <project>
│   ├── router
│   │   └── handler
│   │       ├── middleware.go
│   │       ├── <project>
│   │       │   ├── hanlder.go
│   │       │   └── request.go
│   └── usecase
│       ├── <project>
├── istio.yaml
├── k8s.yaml
├── kafka.yaml
└── pkg
    ├── public pkg 1
    └── public pkg 2
```

## Dependency Graph

![](/images/dependency_graph.png)


