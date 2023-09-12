# ENV

提供我研发过程中常用的环境.

- 镜像会经常性调整，在业务上一定要具体到 SHA256。
- 建议优先使用 podman 去使用容器，如果出现问题，再考虑 docker。

```
docker pull ghcr.io/hongfs/env:php82-fpm
docker pull ghcr.io/hongfs/env:php82-cli
docker pull ghcr.io/hongfs/env:php81-cli
docker pull ghcr.io/hongfs/env:php80-cli
docker pull ghcr.io/hongfs/env:php74-cli
docker pull ghcr.io/hongfs/env:php-language
docker pull ghcr.io/hongfs/env:acme.sh
```

## acme.sh

只是一个镜像备份操作，具体使用可见： https://www.hongfs.cn/2023/06/script/github-actions-generates-wildcard-certificates-for-multiple-domain-names-using-acmesh-and-aliyun-dns/

## aliyun-cli

阿里云 CLI。

```shell
$ podman run --rm ghcr.hongfs.dev/hongfs/env:aliyun-cli aliyun version

$ alias aliyun='podman run --rm ghcr.hongfs.dev/hongfs/env:aliyun-cli aliyun'
$ aliyun version
```

## alpine

提供了默认上海时区。

## alpine-debug

相比 alpine ，默认提供了一些调试的基础依赖包。

## aws-cli

AWS CLI。

## web-1

提供基础的请求接口，方便进行压测。

```shell
$ podman run -d -p 80:80 ghcr.io/hongfs/env:web-1

$ curl -s http://127.0.0.1:80
1

$ curl -s http://127.0.0.1:80/status
{"code":1}

# 会在日志中打印请求头
$ curl -s http://127.0.0.1:80/output
{"code":1}

$ curl -s http://127.0.0.1:80/hostname
c5036d0de87a
```

## 翻译生成

```shell
$ podman run --rm -it -v ~/lang:/data/ ghcr.io/hongfs/env:php-language php index.php
```

你需要将 `data.xlsx` 挂载到 `/data/data.xlsx` 下。
