# template-http-grpc

## 1 简介

一个简单的起http以及grpc服务的模板.

## 2 操作

### 2.1 编译

如果想在本地运行, 可以通过如下命令编译得到可执行文件:
```bash
make
或者
make build
```

可执行文件在本地的使用见[下文](#jump-local).

### 2.2 打镜像

如果想在容器中运行, 则需要打镜像, 需要用到的镜像见[下文](#jump-image).
首先将镜像导入本地, 命令如下:
```bash
docker load -i ./images/centos.7-amd64.tar
docker load -i ./images/debian.stretch.tar
docker load -i ./images/golang.1.15.11-gomod-cuda10-gcc49.tar
docker load -i ./images/proto-tools.3.6.tar

rm -rf ./images
```

然后使用如下命令打镜像:
```bash
make image
```

可执行文件在容器中的使用见[下文](#jump-container).

## 3 使用

### 3.1 在本地使用
<span id = "jump-local"></span>

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
<span id = "jump-container"></span>

首先进入容器中, 命令如下:
```bash
docker run --name template-http-grpc -it template-http-grpc:v0.1.0-xxxxxxx-amd64
```

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

## 4 镜像说明
<span id = "jump-image"></span>

* centos.7-amd64.tar: centos镜像
* debian.stretch.tar: debian镜像
* golang.1.15.11-gomod-cuda10-gcc49.tar: golang镜像
* proto-tools.3.6.tar: proto镜像

该工程需要用到上面几个镜像, 可以先把这几个镜像load到本地, 然后删掉, 方便后面使用.
