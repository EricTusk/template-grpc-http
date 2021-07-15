# template-http-grpc

[TOC]

## 1 简介

一个简单的起http以及grpc服务的模板

## 2 操作

1. 编译得到可执行文件, 命令如下:
```bash
make
或者
make build
```

2. 打镜像, 命令如下:
```bash
make image
```

## 3 使用

### 3.1 在本地使用

1. 查看版本, 命令如下:
```bash
./template-http-grpc --version
```

2. 使用默认的配置文件的路径起服务, 命令如下:
```bash
./template-http-grpc
```

3. 打开debug模式起服务, 命令如下:
```bash
./template-http-grpc --verbose
```

### 3.2 在容器中使用

1. 查看版本, 命令如下:
```bash
./template-http-grpc --version
```

2. 配置配置文件的路径起服务, 命令如下:
```bash
./template-http-grpc --config=/config/config.json
```

3. 打开debug模式起服务, 命令如下:
```bash
./template-http-grpc --config=/config/config.json --verbose
```
