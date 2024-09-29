解释

1. cmd/main.go:
   主入口文件，初始化 Gin 框架，并设置路由。
2. config/config.toml:
   配置文件，存放数据库连接字符串、API 密钥等敏感信息。
3. internal/controllers/:
   控制器层，处理 HTTP 请求，并调用相应的服务层逻辑。
4. internal/models/:
   数据模型层，定义数据结构。
5. internal/services/:
   业务逻辑层，处理与业务逻辑相关的操作。
6. internal/repositories/:
   数据访问层，负责与外部系统的交互，如 GitLab API、FreeNAS API 等。
7. router/router.go:
   路由配置文件，定义所有的 API 路由。
8. utils/:
   工具函数，如认证、加密等。
9. static/:
   存放静态文件，如 CSS、JavaScript 等。
10. templates/:
    存放 HTML 模板文件。
