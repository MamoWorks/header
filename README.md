# Header

HTTP 请求记录器。POST 记录请求头和请求体，GET 查看最近 10 条记录。

## Docker 部署

```bash
docker run -d --name header -p 8899:8899 ghcr.io/mamoworks/header:latest
```

推送到 `main` 分支后，GitHub Actions 会自动构建并推送 `latest` 镜像到 GHCR。
