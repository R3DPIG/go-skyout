# go-skyout

苍穹外卖的 Go/Gin 项目骨架。目前有 MySQL 数据模型和 `/ping` 健康检查；管理端和用户端业务接口尚未实现。

## 启动

1. 安装 Go 1.22 或更新版本，并启动 MySQL。
2. 创建数据库：

   ```sql
   CREATE DATABASE `sky-take-out` CHARACTER SET utf8mb4;
   ```
3. 复制 `configs/config.yaml` 为 `configs/config.local.yaml`，在本地文件中填写 MySQL 密码及 JWT 密钥。项目会自动读取这个文件，且它不会被 Git 跟踪。
4. 执行 `go run ./cmd/server`，访问 `http://localhost:8080/ping`。

加载顺序为 `configs/config.yaml`、`configs/config.local.yaml`、环境变量，后者覆盖前者。例如 `APP_PORT=9090` 可以更改监听端口。首次启动会通过 GORM AutoMigrate 创建或更新数据表。

## 常用命令

```sh
make run
make build
make test
go mod tidy
```

Redis 和 JWT 的配置已经预留，但目前尚未接入业务流程。生产环境的数据库变更应使用版本化迁移；当前 AutoMigrate 仅适合这个项目的开发阶段。

金额字段使用 `decimal.Decimal`；序列化为 JSON 时是字符串，例如 `"12.5"`，以避免浮点数精度损失。后续开发接口时应保持这个约定。
