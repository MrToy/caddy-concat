# caddy-concat

This is a plugin for caddy, just like https://github.com/alibaba/nginx-http-concat

## Download

https://github.com/MrToy/caddy-concat/releases


## Usage

Caddyfile:
```
http://localhost {
    cors
}
```

## Requirement

* golang

## Build 

### Step1

```bash
go get
```

### Step2

在 $GOPATH/src/github.com/mholt/caddy/caddyhttp/httpserver/plugin.go 文件里的directives(约635行的位置)添加"concat" 

### Step3

```bash
#GOOS 可能为 linux、windows、darwin
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o caddy ./cmd
```

得到caddy可执行文件

### Step4

将caddy文件添加到环境变量


## Note

如果需要添加插件，在cmd/main.go里添加