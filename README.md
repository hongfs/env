# ENV

提供我研发过程中常用的环境.

```
docker pull ghcr.io/hongfs/env:php82-fpm
docker pull ghcr.io/hongfs/env:php82-cli
docker pull ghcr.io/hongfs/env:php81-cli
docker pull ghcr.io/hongfs/env:php80-cli
docker pull ghcr.io/hongfs/env:php79-cli
docker pull ghcr.io/hongfs/env:php-language
docker pull ghcr.io/hongfs/env:acme.sh
```

## 翻译生成

```shell
$ docker run --rm -it -v ~/lang:/data/ ghcr.io/hongfs/env:php-language
```

你需要将 `data.xlsx` 挂载到 `/data/data.xlsx` 下。
